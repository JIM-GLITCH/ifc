# Ifc

IFC (`.ifc`) language server and VS Code extension built on [fastbelt](https://github.com/TypeFox/fastbelt). Validates ISO 10303-21 STEP physical files against the **IFC4X3** EXPRESS schema: entity/type WHERE rules, global RULE, UNIQUE constraints, aggregate bounds, and DERIVE attribute checks.

**Current extension release:** `0.0.9` — validation tier toggles (Structure / Types / Rules) in the status bar, per-document eval cache, and configuration sync with the language server. See [`vscode-extension/CHANGELOG.md`](vscode-extension/CHANGELOG.md) and [`vscode-extension/README.md`](vscode-extension/README.md) for details, diagnostic codes, and packaging.

This repository includes:

- **`ifc.fb`** — ISO 10303-21 STEP physical file grammar.
- **`gen.go`** — `//go:generate` running `go tool typefox.dev/fastbelt/cmd/fastbelt generate` to emit `*_gen.go`.
- **`schema_gen.go`**, **`entity_where_gen.go`**, **`derived_gen.go`**, etc. — EXPRESS schema and constraint codegen (via `npm run generate`).
- **`services.go`** — LSP service wiring (document, workspace, linking, generated lexer/parser).
- **`cmd/ifc-lsp`** — language server entrypoint (stdio JSON-RPC).
- **`vscode-extension/`** — VS Code language client, TextMate grammar, and packaged LSP binary.
- **Root `package.json`** — npm workspace for building and packaging the extension from the repo root.

## Requirements

- Go (see `go.mod`)
- Node.js 20+ and npm 10+ for the extension

## Regenerate Go sources

From the module root:

```sh
go generate ./...
```

## Language server

```sh
go run ./cmd/ifc-lsp
```

## VS Code extension

From the repository root:

```sh
npm install
npm run build
npm run package
```

- **`npm run build`** runs TypeScript checks, bundles the client with esbuild, and builds the Go language server into `vscode-extension/dist/server`.
- **`npm run package`** produces `vscode-extension/ifc-vscode.vsix` (a `.vsix` you can install with “Install from VSIX…”).
Upstream docs and sources: [https://github.com/TypeFox/fastbelt](https://github.com/TypeFox/fastbelt)
