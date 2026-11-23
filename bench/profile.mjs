// Capture Go pprof profiles from a running IFC language server while a burst
// of didChange + hover requests is in flight.
//
// Usage:
//   IFC_PPROF=127.0.0.1:6060 node bench/profile.mjs <server-exe> <ifc-file> [edits] [gap-ms]
//
// Writes profiles into bench/profiles/ and prints pprof top summaries.
import { spawn, execFileSync } from "node:child_process";
import * as fs from "node:fs";
import * as path from "node:path";

const addr = process.env.IFC_PPROF ?? "127.0.0.1:6060";
const serverExe = process.argv[2] ?? "./ifc-lsp-prof.exe";
const file = process.argv[3] ?? "t.ifc";
const edits = Number(process.argv[4] ?? 20);
const gap = Number(process.argv[5] ?? 30);
const uri = "file:///" + path.resolve(file).replace(/\\/g, "/");
const suffix = "/" + path.basename(file);
const outDir = path.resolve("bench/profiles");
fs.mkdirSync(outDir, { recursive: true });

const text = fs.readFileSync(file, "utf-8");
const kw = text.indexOf("IFCAPPLICATION(");
const quote = text.indexOf("'", kw);
const lines = text.slice(0, quote + 1).split(/\r\n|\r|\n/);
const pos = { line: lines.length - 1, character: lines[lines.length - 1].length };

const child = spawn(serverExe, [], {
  stdio: ["pipe", "pipe", "pipe"],
  env: { ...process.env, IFC_PPROF: addr },
});
child.stderr.on("data", (d) => process.stderr.write("[srv] " + d.toString()));
let buf = Buffer.alloc(0);
let started = false;
const send = (o) => {
  const json = JSON.stringify(o);
  child.stdin.write(`Content-Length: ${Buffer.byteLength(json)}\r\n\r\n${json}`);
};
const capture = (name, query = "") => {
  const url = `http://${addr}/debug/pprof/${name}${query}`;
  const out = path.join(outDir, `${name}${query ? "-stacks" : ""}.txt`);
  try {
    const data = execFileSync("curl", ["-sS", "-L", url], { maxBuffer: 256 * 1024 * 1024 });
    fs.writeFileSync(out, data);
    return out;
  } catch (e) {
    console.error(`capture ${name} failed: ${e.message}`);
    return null;
  }
};

child.stdout.on("data", (chunk) => {
  buf = Buffer.concat([buf, chunk]);
  for (;;) {
    const end = buf.indexOf("\r\n\r\n");
    if (end < 0) return;
    const head = buf.slice(0, end).toString("utf8");
    const m = /Content-Length: (\d+)/i.exec(head);
    if (!m) { buf = buf.slice(end + 4); continue; }
    const len = Number(m[1]);
    const start = end + 4;
    if (buf.length < start + len) return;
    const body = buf.slice(start, start + len).toString("utf8");
    buf = buf.slice(start + len);
    let msg;
    try { msg = JSON.parse(body); } catch { continue; }
    if (msg.method === "textDocument/publishDiagnostics" && !started) {
      started = true;
      console.log("open diagnostics received, starting burst + profiling");
      setTimeout(() => {
        for (let i = 0; i < edits; i++) {
          setTimeout(() => {
            send({ jsonrpc: "2.0", method: "textDocument/didChange", params: {
              textDocument: { uri, version: i + 2 },
              contentChanges: [{ range: { start: pos, end: { line: pos.line, character: pos.character + 1 } }, text: i % 2 ? "A" : "Z" }],
            } });
            send({ jsonrpc: "2.0", id: 200 + i, method: "textDocument/hover", params: { textDocument: { uri }, position: pos } });
          }, i * gap);
        }
        const captured = [];
        captured.push(capture("goroutine", "?debug=2"));
        captured.push(capture("block", "?debug=2"));
        captured.push(capture("mutex", "?debug=2"));
        console.log("captured:", captured.filter(Boolean).join(", "));
        console.log("stacks: bench/profiles/goroutine-stacks.txt (search for sync.(*Cond).Wait / chan receive / Lock)");
        child.kill();
        process.exit(0);
      }, 300);
    }
  }
});

send({ jsonrpc: "2.0", id: 1, method: "initialize", params: {
  processId: process.pid,
  rootUri: uri.slice(0, uri.lastIndexOf("/") + 1),
  capabilities: { textDocument: { publishDiagnostics: {}, hover: { contentFormat: ["markdown"] } } },
} });
setTimeout(() => {
  send({ jsonrpc: "2.0", method: "initialized", params: {} });
  send({ jsonrpc: "2.0", method: "textDocument/didOpen", params: { textDocument: { uri, languageId: "ifc", version: 1, text } } });
}, 200);
setTimeout(() => { console.error("TIMEOUT"); child.kill(); process.exit(1); }, 30000);
