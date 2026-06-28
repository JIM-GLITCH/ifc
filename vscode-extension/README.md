# Ifc for VS Code

VS Code extension for editing **IFC physical files** (`.ifc`) — the text-based exchange format defined by [ISO 10303-21](https://www.iso.org/standard/63141.html), commonly used in building information modeling (BIM) and other STEP-based workflows.

Open any `.ifc` file and the extension activates automatically. No extra configuration is required.

## Features

### Language recognition

- Registers the **Ifc** language for files with the `.ifc` extension
- Applies the correct language mode, grammar, and editor behavior as soon as you open a file

### Syntax highlighting

Colorizes the structure of ISO 10303-21 exchange files so you can scan large models more easily, including:

- File header and data section markers (`ISO-10303-21;`, `HEADER;`, `DATA;`, `ENDSEC;`, etc.)
- Entity instance references (`#123`)
- Keywords, strings, numbers, enumerations, and binary values
- Line (`//`) and block (`/* */`) comments

### Smart editing

- **Bracket and quote matching** for `( )`, `[ ]`, `{ }`, `'`, and `"`
- **Auto-closing pairs** when you type opening brackets or quotes
- Familiar comment toggling via `Ctrl+/` / `Cmd+/` (line comments) and block comment support

### Language Server (LSP)

A built-in language server analyzes your file as you edit and integrates with standard VS Code features:

| Feature | What it does |
|---------|----------------|
| **Diagnostics** | Syntax errors plus schema-aware validation inline and in the Problems panel |
| **Go to Definition** | Jump to the entity instance referenced by a `#` id |
| **Find All References** | See every place an entity instance is used in the file |
| **Completion** | Context-aware suggestions based on the ISO 10303-21 grammar |
| **Document sync** | Keeps analysis up to date while you type |

Server activity is logged in the **Output** panel under **Ifc Language Server**.

### Schema validation (IFC4 / IFC4X3)

Beyond parse errors, the language server validates entity instances against the **IFC4X3** EXPRESS schema (files declaring `IFC4`, `IFC4X3`, or `IFC4X3_DEV_*` in `FILE_SCHEMA`).

Checks include:

- **Entity name** — unknown or abstract entity types
- **Parameter count** — required vs optional attributes; trailing optional parameters may be omitted
- **Parameter types** — literals, lists, entity references, and typed parameters (`IFCLABEL(...)`, `IFCRATIOMEASURE(...)`, etc.)
- **Entity references** — `#` references must match the expected entity type (including subtypes)
- **Enumerations** — values such as `.MODIFIED.` must belong to the declared enum type
- **Aggregates** — `LIST` / `SET` elements, including nested lists (e.g. `IFCCARTESIANPOINTLIST2D` coordinate arrays)
- **Derived attributes** — `*` is accepted where the schema marks an attribute as derived; `$` is used for omitted optional values

Diagnostic codes (shown in the Problems panel):

| Code | Meaning |
|------|---------|
| `ifc-unknown-entity` | Entity keyword is not in the schema |
| `ifc-abstract-entity` | Abstract entity cannot be instantiated |
| `ifc-param-count` | Too many parameters, or a required parameter is missing |
| `ifc-param-required` | Required attribute is `$` (unset) |
| `ifc-param-derived` | `*` used on a non-derived attribute |
| `ifc-param-type` | Value shape or type does not match the attribute |
| `ifc-entity-type` | `#` reference points to the wrong entity type |
| `ifc-enum-value` | Invalid enumeration literal |
| `ifc-typed-keyword` | Typed parameter keyword does not match the attribute type |

## Requirements

- [Visual Studio Code](https://code.visualstudio.com/) `1.104.0` or later

## Installation

1. Download or obtain the extension package (`ifc-vscode.vsix`).
2. In VS Code, open the **Extensions** view (`Ctrl+Shift+X` on Windows/Linux, `Cmd+Shift+X` on macOS).
3. Click the `...` menu at the top of the Extensions sidebar.
4. Choose **Install from VSIX...** and select the `.vsix` file.
5. Reload VS Code if prompted.

After installation, open any `.ifc` file to start using syntax highlighting and language-server features.

## Usage tips

- Use **Go to Definition** (`F12`) on a `#` entity reference to navigate to its declaration.
- Use **Find All References** (`Shift+F12`) to trace how an entity is referenced elsewhere in the file.
- Check the **Problems** panel (`Ctrl+Shift+M` / `Cmd+Shift+M`) for parse and schema validation issues.
- Filter Problems by diagnostic code (e.g. `ifc-param-type`) to focus on a class of errors.
- For troubleshooting, open **View → Output** and select **Ifc Language Server** from the dropdown.

## Known limitations

- Support targets the ISO 10303-21 physical file format (`.ifc` text files). Other IFC representations (e.g. IFC-XML, IFC-ZIP) are not covered.
- Validation uses a compiled IFC4X3 schema; entities or types from other schema releases may not be recognized.
- WHERE rules and full EXPRESS semantic constraints are not evaluated — validation covers STEP encoding shape and declared attribute types.
- The extension focuses on single-file editing; cross-file workspace analysis depends on how your files are organized.
