// Decides whether a language server uses the text the client sends over LSP
// or reads the file from disk.
//
// The same URI is opened twice: once with the real file content, once with
// content that differs from the file on disk. If the diagnostics change, the
// server is using the LSP text; if they stay identical, it is reading the file.
//
// Usage: node bench/probe-source.mjs <server-exe> <ifc-file>
import { spawn } from "node:child_process";
import * as fs from "node:fs";
import * as path from "node:path";

const server = process.argv[2];
const file = process.argv[3] ?? "t.ifc";
const uri = "file:///" + path.resolve(file).replace(/\\/g, "/");
const suffix = "/" + path.basename(file);

const real = fs.readFileSync(file, "utf-8");
// Same URI, but the document is truncated to a quarter and given a bogus
// header. A disk-reading server would never see this content.
const doctored = "NOT-A-STEP-FILE;\n" + real.slice(0, Math.floor(real.length / 4));

function run(label, text) {
  return new Promise((resolve, reject) => {
    const child = spawn(server, [], { stdio: ["pipe", "pipe", "pipe"] });
    let buf = Buffer.alloc(0);

    child.on("error", reject);
    child.stderr.on("data", () => {});
    const send = (obj) => {
      const json = JSON.stringify(obj);
      child.stdin.write(`Content-Length: ${Buffer.byteLength(json)}\r\n\r\n${json}`);
    };

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
        if (msg.method === "textDocument/publishDiagnostics" && String(msg.params.uri).endsWith(suffix)) {
          child.kill();
          resolve({
            published: msg.params.diagnostics.length,
            messages: msg.params.diagnostics.slice(0, 3).map((d) => d.message.slice(0, 70)),
          });
        }
      }
    });

    send({
      jsonrpc: "2.0",
      id: 1,
      method: "initialize",
      params: {
        processId: process.pid,
        rootUri: uri.slice(0, uri.lastIndexOf("/") + 1),
        capabilities: { textDocument: { publishDiagnostics: {} } },
      },
    });
    setTimeout(() => {
      send({ jsonrpc: "2.0", method: "initialized", params: {} });
      send({
        jsonrpc: "2.0",
        method: "textDocument/didOpen",
        params: { textDocument: { uri, languageId: "ifc", version: 1, text } },
      });
    }, 200);
    setTimeout(() => {
      child.kill();
      reject(new Error(`${label}: no diagnostics published`));
    }, 120000);
  });
}

console.log(`  file on disk : ${path.resolve(file)} (${real.length} bytes)`);
const a = await run("real text   ", real);
console.log(`  ${a.label}: ${a.published} diagnostics`);
for (const m of a.messages) console.log(`      - ${m}`);

const b = await run("doctored    ", doctored);
console.log(`  ${b.label}: ${b.published} diagnostics`);
for (const m of b.messages) console.log(`      - ${m}`);

console.log(
  `\n  verdict: ${
    a.published !== b.published
      ? "USES THE LSP TEXT (diagnostics differ)"
      : "READS FROM DISK (diagnostics identical)"
  }`
);
