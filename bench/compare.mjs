// Cross-server comparison harness.
//
// Runs the identical LSP conversation against any IFC language server:
//   initialize -> initialized -> didOpen -> wait for the first
//   publishDiagnostics -> didChange (one character inside one entity
//   instance) -> wait for the next publishDiagnostics.
//
// Both phases are timed from the moment the request is written to the moment
// the server's publishDiagnostics arrives, which is what an editor experiences.
//
// Usage:
//   node bench/compare.mjs <server-exe> <ifc-file> [iterations]
//
// The publishDiagnostics URI is matched by file name suffix, because servers
// normalise Windows URIs differently (the ifc server percent-encodes the drive
// letter and backslashes, the NepomukWolf server does not).
import { spawn } from "node:child_process";
import * as fs from "node:fs";
import * as path from "node:path";

const server = process.argv[2];
const file = process.argv[3] ?? "t.ifc";
const iterations = Number(process.argv[4] ?? 5);
const uri = "file:///" + path.resolve(file).replace(/\\/g, "/");
const suffix = "/" + path.basename(file);

function insideInstancePosition(text) {
  const keyword = text.indexOf("IFCAPPLICATION(");
  if (keyword < 0) throw new Error("IFCAPPLICATION( not found");
  const quote = text.indexOf("'", keyword);
  const before = text.slice(0, quote + 1);
  const lines = before.split(/\r\n|\r|\n/);
  return { line: lines.length - 1, character: lines[lines.length - 1].length };
}

const text = fs.readFileSync(file, "utf-8");
const inside = insideInstancePosition(text);

function runOnce(i) {
  return new Promise((resolve, reject) => {
    const child = spawn(server, [], { stdio: ["pipe", "pipe", "pipe"] });
    let buf = Buffer.alloc(0);
    let openMs = null;
    let editMs = null;
    let openSentAt = 0n;
    let editSentAt = 0n;
    let opened = false;

    child.on("error", reject);
    if (process.env.COMPARE_VERBOSE) {
      child.stderr.on("data", (d) => process.stderr.write("[srv] " + d.toString()));
    } else {
      child.stderr.on("data", () => {});
    }

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
        if (msg.method === "textDocument/publishDiagnostics") {
          if (!String(msg.params.uri).endsWith(suffix)) return;
          const count = msg.params.diagnostics.length;
          if (!opened) {
            opened = true;
            openMs = Number(process.hrtime.bigint() - openSentAt) / 1e6;
            editSentAt = process.hrtime.bigint();
            send({
              jsonrpc: "2.0",
              method: "textDocument/didChange",
              params: {
                textDocument: { uri, version: 2 },
                contentChanges: [
                  {
                    range: {
                      start: inside,
                      end: { line: inside.line, character: inside.character + 1 },
                    },
                    text: "Z",
                  },
                ],
              },
            });
            setTimeout(() => {
              if (editMs === null) {
                child.kill();
                reject(new Error("no diagnostics after didChange"));
              }
            }, 300000);
          } else if (editMs === null) {
            editMs = Number(process.hrtime.bigint() - editSentAt) / 1e6;
            child.kill();
            resolve({ openMs, openCount: count, editMs, editCount: count });
          }
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
        capabilities: {
          textDocument: {
            publishDiagnostics: { relatedInformation: true },
            hover: { contentFormat: ["markdown", "plaintext"] },
            definition: { linkSupport: true },
            documentSymbol: { hierarchicalDocumentSymbolSupport: true },
          },
        },
      },
    });
    setTimeout(() => {
      send({ jsonrpc: "2.0", method: "initialized", params: {} });
      openSentAt = process.hrtime.bigint();
      send({
        jsonrpc: "2.0",
        method: "textDocument/didOpen",
        params: { textDocument: { uri, languageId: "ifc", version: 1, text } },
      });
    }, 200);
    setTimeout(() => {
      if (openMs === null) {
        child.kill();
        reject(new Error("no diagnostics after didOpen"));
      }
    }, 300000);
  });
}

const results = [];
for (let i = 0; i < iterations; i++) {
  const r = await runOnce(i);
  results.push(r);
  console.log(
    `  run ${i + 1}: open=${r.openMs.toFixed(1)}ms (${r.openCount} diags)  edit=${r.editMs.toFixed(1)}ms (${r.editCount} diags)`
  );
}

const median = (xs) => {
  const s = [...xs].sort((a, b) => a - b);
  return s[Math.floor(s.length / 2)];
};
const pct = (xs, p) => {
  const s = [...xs].sort((a, b) => a - b);
  return s[Math.min(s.length - 1, Math.floor((p / 100) * s.length))];
};
// Report the spread, not just the median: a Go server's GC pauses show up in
// the tail, and the tail is what a user actually feels while typing.
const spread = (xs) =>
  `min ${Math.min(...xs).toFixed(0).padStart(5)}  p50 ${median(xs).toFixed(0).padStart(5)}` +
  `  p95 ${pct(xs, 95).toFixed(0).padStart(5)}  max ${Math.max(...xs).toFixed(0).padStart(5)} ms`;

console.log(`\n  open->diagnostics : ${spread(results.map((r) => r.openMs))}`);
console.log(`  edit->diagnostics : ${spread(results.map((r) => r.editMs))}`);
console.log(`  diagnostics on open : ${results[0].openCount}`);
console.log(`  file                : ${path.resolve(file)} (${text.length} bytes)`);
