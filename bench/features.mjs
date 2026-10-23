// Per-feature latency probe.
//
// Times the individual LSP features (documentSymbol, definition, references,
// hover, completion) on an already-open document, so the numbers exclude the
// build pipeline and isolate the query path.
//
// Usage:
//   node bench/features.mjs <server-exe> <ifc-file> [rounds]
import { spawn } from "node:child_process";
import * as fs from "node:fs";
import * as path from "node:path";

const server = process.argv[2];
const file = process.argv[3] ?? "t.ifc";
const rounds = Number(process.argv[4] ?? 20);
const uri = "file:///" + path.resolve(file).replace(/\\/g, "/");
const suffix = "/" + path.basename(file);

// "#5= IFCAPPLICATION(#1,'2020',...)" — "#1" starts at line 28, char 21.
const position = { line: 28, character: 21 };

const text = fs.readFileSync(file, "utf-8");

const child = spawn(server, [], { stdio: ["pipe", "pipe", "pipe"] });
let buf = Buffer.alloc(0);
const pending = new Map();
let nextId = 100;
let opened = false;

child.stderr.on("data", () => {});

const send = (obj) => {
  const json = JSON.stringify(obj);
  child.stdin.write(`Content-Length: ${Buffer.byteLength(json)}\r\n\r\n${json}`);
};

const request = (method, params) =>
  new Promise((resolve, reject) => {
    const id = nextId++;
    pending.set(id, { resolve, reject });
    send({ jsonrpc: "2.0", id, method, params });
  });

child.stdout.on("data", (chunk) => {
  buf = Buffer.concat([buf, chunk]);
  for (;;) {
    const headerEnd = buf.indexOf("\r\n\r\n");
    if (headerEnd < 0) return;
    const header = buf.slice(0, headerEnd).toString("utf8");
    const match = /Content-Length: (\d+)/i.exec(header);
    if (!match) {
      buf = buf.slice(headerEnd + 4);
      continue;
    }
    const length = Number(match[1]);
    const start = headerEnd + 4;
    if (buf.length < start + length) return;
    const body = buf.slice(start, start + length).toString("utf8");
    buf = buf.slice(start + length);
    let msg;
    try {
      msg = JSON.parse(body);
    } catch {
      continue;
    }
    if (msg.id !== undefined && pending.has(msg.id)) {
      const { resolve } = pending.get(msg.id);
      pending.delete(msg.id);
      resolve(msg.result);
    }
  }
});

await new Promise((resolve) => {
  send({
    jsonrpc: "2.0",
    id: 1,
    method: "initialize",
    params: {
      processId: process.pid,
      rootUri: uri.slice(0, uri.lastIndexOf("/") + 1),
      capabilities: {
        textDocument: {
          publishDiagnostics: {},
          hover: { contentFormat: ["markdown", "plaintext"] },
          definition: { linkSupport: true },
          references: {},
          documentSymbol: { hierarchicalDocumentSymbolSupport: true },
          completion: { completionItem: { snippetSupport: true } },
        },
      },
    },
  });
  const wait = setInterval(() => {
    if (pending.size === 0 && nextId > 100) {
      clearInterval(wait);
      resolve();
    }
  }, 50);
  setTimeout(() => {
    clearInterval(wait);
    resolve();
  }, 10000);
});

send({ jsonrpc: "2.0", method: "initialized", params: {} });
send({
  jsonrpc: "2.0",
  method: "textDocument/didOpen",
  params: { textDocument: { uri, languageId: "ifc", version: 1, text } },
});

// Wait for the document to be built before timing queries.
await new Promise((resolve) => setTimeout(resolve, 4000));

const probes = [
  ["documentSymbol", () => request("textDocument/documentSymbol", { textDocument: { uri } })],
  [
    "definition",
    () =>
      request("textDocument/definition", {
        textDocument: { uri },
        position,
      }),
  ],
  [
    "references",
    () =>
      request("textDocument/references", {
        textDocument: { uri },
        position,
        context: { includeDeclaration: true },
      }),
  ],
  ["hover", () => request("textDocument/hover", { textDocument: { uri }, position })],
  [
    "completion",
    () =>
      request("textDocument/completion", {
        textDocument: { uri },
        position: { line: 28, character: 10 },
      }),
  ],
];

const median = (xs) => {
  const s = [...xs].sort((a, b) => a - b);
  return s[Math.floor(s.length / 2)];
};

for (const [name, fn] of probes) {
  const times = [];
  let size = 0;
  for (let i = 0; i < rounds; i++) {
    const t0 = process.hrtime.bigint();
    const result = await fn();
    times.push(Number(process.hrtime.bigint() - t0) / 1e6);
    size = JSON.stringify(result ?? null).length;
  }
  console.log(`  ${name.padEnd(15)} median ${median(times).toFixed(2).padStart(8)} ms   (payload ${size} B)`);
}

child.kill();
process.exit(0);
