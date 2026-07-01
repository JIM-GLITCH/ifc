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
| **Diagnostics** | Syntax errors plus **EXPRESS schema validation** inline and in the Problems panel |
| **Go to Definition** | Jump to the entity instance referenced by a `#` id |
| **Find All References** | See every place an entity instance is used in the file |
| **Completion** | Context-aware suggestions based on the ISO 10303-21 grammar |
| **Document sync** | Keeps analysis up to date while you type |

Server activity is logged in the **Output** panel under **Ifc Language Server**.

### EXPRESS schema validation (IFC4 / IFC4X3)

This extension validates `.ifc` files against the **IFC4X3** EXPRESS schema (compiled from the official schema; files declaring `IFC4`, `IFC4X3`, or `IFC4X3_DEV_*` in `FILE_SCHEMA` are accepted).

Validation goes beyond STEP syntax and attribute typing — it evaluates **EXPRESS constraints** from the schema:

| Layer | What is checked |
|-------|-----------------|
| **STEP encoding** | Parameter count; `$` (unset optional) and `*` (derived); literals, lists, typed parameters, entity references |
| **Attribute types** | Entity / select / enum / aggregate typing; nested list elements |
| **Type WHERE** | WHERE rules on simple and defined types (e.g. `IfcPositiveInteger`, `IfcCompoundPlaneAngleMeasure`) |
| **Entity WHERE** | All entity WHERE rules in the schema (752 rules for IFC4X3) |
| **Global RULE** | Schema-wide rules that span multiple entities |
| **UNIQUE** | Entity UNIQUE clauses and `OF UNIQUE` aggregate elements |
| **Aggregate bounds** | Declared sizes such as `LIST [1:3]`, `SET [0:?]`, `ARRAY [1:2]` |
| **Derived attributes** | `*` omits exchanged values; DERIVE expressions supply effective values for WHERE checks (e.g. `NVL(Scale, 1.0)`) |

Optional `$` values are treated as absent (no type-level WHERE on missing data). Derived `*` values are checked using computed DERIVE results where evaluators are available.

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
| `ifc-type-where` | Violates a WHERE rule on a defined type |
| `ifc-entity-where` | Violates an entity WHERE rule |
| `ifc-global-rule` | Violates a global RULE |
| `ifc-unique` | Violates an entity UNIQUE constraint |
| `ifc-unique-aggregate` | Duplicate element in a `UNIQUE` aggregate |
| `ifc-aggregate-bounds` | Aggregate element count outside declared `[low:high]` bounds |
| `ifc-eval-unknown` | Expression or constraint check could not be evaluated (e.g. unresolved `#` reference, DERIVE failure) |

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
- Check the **Problems** panel (`Ctrl+Shift+M` / `Cmd+Shift+M`) for parse and EXPRESS validation issues.
- Filter Problems by diagnostic code (e.g. `ifc-entity-where`, `ifc-type-where`) to focus on a class of errors.
- For troubleshooting, open **View → Output** and select **Ifc Language Server** from the dropdown.

## Known limitations

- Support targets the ISO 10303-21 physical file format (`.ifc` text files). Other IFC representations (e.g. IFC-XML, IFC-ZIP) are not covered.
- Validation uses a compiled **IFC4X3** schema; entities or types from other schema releases may not be recognized.
- **DERIVE** codegen covers all schema DERIVE slots; unresolved references or failed DERIVE/type-WHERE evaluation now report `ifc-eval-unknown` instead of silently skipping. Optional `$` / EXPRESS absent values still skip checks.
- This is an **EXPRESS / schema** validator, not a full BIM certification tool: no geometry kernel, IDS/MVD filtering, or clash detection.
- The extension focuses on single-file editing; cross-file workspace analysis depends on how your files are organized.
