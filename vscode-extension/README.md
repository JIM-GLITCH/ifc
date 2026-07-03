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

### Validation layer status bar

When an `.ifc` file is the active editor, the status bar shows **`IFC n/3`** — how many of the three validation tiers are enabled. Click it (or run **IFC: Toggle Validation Layers**) to open a quick pick:

| Tier | Setting | What it controls |
|------|---------|------------------|
| **Structure** | `ifc.validation.structure` | STEP encoding (`ifc-param-*`), attribute types (`ifc-param-type`, `ifc-entity-type`, `ifc-enum-*`), aggregate bounds, unknown/abstract entity |
| **Types** | `ifc.validation.types` | Type WHERE (`ifc-type-where`), DERIVE `*` evaluation (`ifc-eval-unknown` on derived attributes) |
| **Rules** | `ifc.validation.rules` | Entity WHERE (`ifc-entity-where`), Global RULE (`ifc-global-rule`), UNIQUE (`ifc-unique`, `ifc-unique-aggregate`) |

Lexer/parser errors and unresolved `#` linker diagnostics are **always** reported regardless of these toggles. Changing a setting re-validates open `.ifc` files automatically.

You can also edit the same keys under **Settings → Ifc**.

#### Performance

During each validation pass the language server:

- **Indexes inverse relations once** when building the entity population (forward `#` references → inverse lists such as `IsTypedBy`). Entity WHERE rules query this index; they do not rescan the file.
- **Caches attribute values** (`eval_cache`) keyed by entity `#id` and attribute name so Types-layer DERIVE checks and Rules-layer Entity WHERE do not repeat the same DERIVE or parameter conversion work.

Turn off **Types** or **Rules** in the status bar if you want faster feedback while editing and only need structural checks.

### EXPRESS schema validation (IFC4 / IFC4X3)

This extension validates `.ifc` files against the **IFC4X3** EXPRESS schema (compiled from the official schema; files declaring `IFC4`, `IFC4X3`, or `IFC4X3_DEV_*` in `FILE_SCHEMA` are accepted).

Validation goes beyond STEP syntax and attribute typing — it evaluates **EXPRESS constraints** from the schema. The status-bar tiers map to these checks:

| Tier (status bar) | What is checked |
|-------------------|-----------------|
| **Structure** | Parameter count; `$` / `*` encoding; attribute types; enums; `#` entity types; aggregate bounds; unknown/abstract entity |
| **Types** | Type WHERE on defined types; DERIVE evaluation for omitted `*` attributes |
| **Rules** | Entity WHERE (752 rules for IFC4X3); Global RULE; UNIQUE and `OF UNIQUE` aggregates |

Always reported (not toggled): STEP **parse errors** and unresolved `#` **linker** diagnostics.

Optional `$` values are treated as absent (no type-level WHERE on missing data). Derived `*` values are checked using computed DERIVE results where evaluators are available.

WHERE and RULE evaluation uses **lazy short-circuiting**: once a branch cannot be evaluated, later OR/AND branches in the same rule are skipped, and further WHERE rules on the same entity may be skipped when an eval-unknown is already recorded.

Diagnostics are **anchored** to the most specific location available: derived `*` tokens, other parameter tokens, or the declaring entity `#id`. Related entities in a reference chain (e.g. a parent `IFCGEOMETRICREPRESENTATIONCONTEXT` and its `IFCGEOMETRICREPRESENTATIONSUBCONTEXT` children) may each report issues so you can see the full impact radius.

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
| `ifc-eval-unknown` | Expression or constraint check could not be evaluated (e.g. DERIVE failure, parameter conversion failure). Entity WHERE / Global RULE messages include the rule label when applicable. |

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
- Click **`IFC n/3`** in the status bar (when an `.ifc` file is active) to enable or disable Structure / Types / Rules validation tiers.
- Filter Problems by diagnostic code (e.g. `ifc-entity-where`, `ifc-type-where`) to focus on a class of errors.
- For troubleshooting, open **View → Output** and select **Ifc Language Server** from the dropdown.

## Known limitations

- Support targets the ISO 10303-21 physical file format (`.ifc` text files). Other IFC representations (e.g. IFC-XML, IFC-ZIP) are not covered.
- Validation uses a compiled **IFC4X3** schema; entities or types from other schema releases may not be recognized.
- **DERIVE** codegen covers all schema DERIVE slots; failed DERIVE/type-WHERE evaluation reports `ifc-eval-unknown` instead of silently skipping. Unresolved `#` references are reported by the linker only (not duplicated as eval-unknown). Optional `$` / EXPRESS absent values still skip checks. Disable the **Types** tier to skip DERIVE and Type WHERE when validation feels slow on large files.
- Files with **parse errors** skip Global RULE and UNIQUE validation; remaining EXPRESS checks may still report eval-unknown on affected entities and parameters.
- This is an **EXPRESS / schema** validator, not a full BIM certification tool: no geometry kernel, IDS/MVD filtering, or clash detection.
- The extension focuses on single-file editing; cross-file workspace analysis depends on how your files are organized.
