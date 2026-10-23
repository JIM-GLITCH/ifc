import { addRegisteredLanguage } from "lsbench";
import type { BenchContext } from "lsbench";
import * as fs from "node:fs";
import * as path from "node:path";

// The ifc language server registers the custom language id "ifc" for the
// ".ifc" extension, so tell lsbench about it before opening documents.
addRegisteredLanguage("ifc", ".ifc");

// lsbench's pathToUri() builds "file://C:\..." on Windows (backslashes, no
// leading slash). The fastbelt server parses such URIs leniently and
// normalizes them to "file:///C%3A%5C...%5Ct.ifc", then publishes diagnostics
// under the normalized URI. waitForDiagnostics() compares URI strings, so it
// would never match and time out. This helper reconstructs the normalized form
// the server actually uses, so the waits below resolve on the real
// publishDiagnostics notification.
function serverNormalizedUri(ctx: BenchContext, relativePath: string): string {
  const abs = path.join(ctx.workspaceRoot, relativePath);
  return (
    "file:///" +
    abs.replace(/^([A-Za-z]):/, (_match, drive: string) => `${drive}%3A`).replace(/\\/g, "%5C")
  );
}

// Align the URI lsbench waits on with the normalized URI the server publishes
// diagnostics for (see serverNormalizedUri above).
function alignUri(ctx: BenchContext, relativePath: string): void {
  const openDocs = (
    ctx as unknown as {
      openDocuments: Map<string, { uri: string; version: number; text: string }>;
    }
  ).openDocuments;
  const entry = openDocs.get(relativePath);
  if (entry) {
    entry.uri = serverNormalizedUri(ctx, relativePath);
  }
}

// End-of-file position of a document, used for the trivial edit below.
function endPosition(relativePath: string): { line: number; character: number } {
  const abs = path.join(process.cwd(), relativePath);
  const text = fs.readFileSync(abs, "utf-8");
  const lines = text.split(/\r\n|\r|\n/);
  return { line: lines.length - 1, character: lines[lines.length - 1].length };
}

// Position of the first character inside the first quoted string of the first
// IFCAPPLICATION instance. Editing one character there changes the text of
// exactly one entity instance, which is what the incremental validation path is
// for. Any character is syntactically valid inside a STEP string literal, so the
// edit cannot introduce a parse error and cannot invalidate the schema header.
function insideInstancePosition(relativePath: string): { line: number; character: number } | null {
  const abs = path.join(process.cwd(), relativePath);
  const text = fs.readFileSync(abs, "utf-8");
  const keyword = text.indexOf("IFCAPPLICATION(");
  if (keyword < 0) {
    return null;
  }
  const quote = text.indexOf("'", keyword);
  if (quote < 0) {
    return null;
  }
  const before = text.slice(0, quote + 1);
  const lines = before.split(/\r\n|\r|\n/);
  return { line: lines.length - 1, character: lines[lines.length - 1].length };
}

// End-to-end benchmark for the ifc language server.
//
// Metrics (one line pair per iteration):
//   [bench] open->diagnostics        didOpen -> publishDiagnostics on a freshly
//                                    started server: the full cold build
//                                    pipeline (lex -> parse -> symbols -> link
//                                    -> validate).
//   [bench] noop-edit->diagnostics   didChange -> publishDiagnostics after
//                                    appending a space at end of file. No entity
//                                    instance text changes, so this measures the
//                                    best case: every cached entity diagnostic is
//                                    reused as-is. Kept for comparability with
//                                    releases up to 0.0.13.
//   [bench] edit->diagnostics        didChange -> publishDiagnostics after
//                                    changing one character inside one entity
//                                    instance. This is the representative
//                                    interactive case: a full re-parse plus an
//                                    incremental re-validation.
// The per-method table additionally reports documentSymbol/definition/
// references/hover latencies on the fully built document.
//
// IMPORTANT: run with --restart. fastbelt reuses an unchanged document across
// didOpen calls (republishing cached diagnostics) and a didClose empty publish
// can arrive skewed into the next iteration, which makes keep-alive
// iterations measure the wrong notification. With --restart every iteration
// starts from a fresh server, so each phase produces exactly one publish.
//
// Usage (from the repository root, so that "go run ./cmd/ifc-lsp" resolves):
//   npx lsbench "go run ./cmd/ifc-lsp" -w . -s ./bench/bench-actions.ts ^
//     -n 5 --warmup 1 --restart ^
//     --init-options-file ./bench/init-options-full.json ^
//     -o report.json
//
// To isolate parse + link from validation, swap in
// ./bench/init-options-parse-only.json (all validation tiers off); the
// difference between the two runs is the net validation cost.
//
// For lower noise, prebuild the server once ("go build -o bin/ifc-lsp
// ./cmd/ifc-lsp") and use "bin/ifc-lsp" as the server command.
//
// Positions below refer to t.ifc (LSP positions are 0-based):
//   line 28 = "#5= IFCAPPLICATION(#1,'2020',...)", where "#1" starts at char 21.
export default async function (ctx: BenchContext) {
  const rel = "t.ifc";

  const openStart = process.hrtime.bigint();
  await ctx.openDocument(rel);
  alignUri(ctx, rel);
  const openDiags = await ctx.waitForDiagnostics(rel, 60_000);
  console.log(`[bench] open wait resolved with ${openDiags.length} diags`);
  const openMs = Number(process.hrtime.bigint() - openStart) / 1e6;
  console.log(`[bench] open->diagnostics: ${openMs.toFixed(1)}ms`);

  // Best case: force a full rebuild with a trivial edit that changes no entity
  // instance text, and measure the warm path.
  const end = endPosition(rel);
  const noopStart = process.hrtime.bigint();
  await ctx.edit(rel, {
    range: { start: end, end },
    text: " ",
  });
  const noopDiags = await ctx.waitForDiagnostics(rel, 60_000);
  console.log(`[bench] noop-edit wait resolved with ${noopDiags.length} diags`);
  const noopMs = Number(process.hrtime.bigint() - noopStart) / 1e6;
  console.log(`[bench] noop-edit->diagnostics: ${noopMs.toFixed(1)}ms`);

  // Representative case: change one character inside one entity instance.
  const inside = insideInstancePosition(rel);
  if (!inside) {
    throw new Error("could not locate an entity instance to edit in t.ifc");
  }
  const editStart = process.hrtime.bigint();
  await ctx.edit(rel, {
    range: { start: inside, end: { line: inside.line, character: inside.character + 1 } },
    text: "Z",
  });
  const editDiags = await ctx.waitForDiagnostics(rel, 60_000);
  console.log(`[bench] edit wait resolved with ${editDiags.length} diags`);
  const editMs = Number(process.hrtime.bigint() - editStart) / 1e6;
  console.log(`[bench] edit->diagnostics: ${editMs.toFixed(1)}ms`);

  await ctx.documentSymbol(rel);
  await ctx.definition(rel, 28, 21);
  await ctx.references(rel, 28, 21);
  await ctx.hover(rel, 28, 21);

  await ctx.closeDocument(rel);
}
