# Changelog

All notable changes to the Ifc language server and VS Code extension are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [0.0.7] - 2026-06-26

### Added

- **EXPRESS WHERE validation**
  - Entity WHERE: all 752 entity WHERE rules for IFC4X3 (`ifc-entity-where`).
  - Type WHERE: WHERE rules on defined types (e.g. `IfcPositiveInteger`, `IfcCompoundPlaneAngleMeasure`) (`ifc-type-where`).
  - Global RULE: schema-wide cross-entity rules (`ifc-global-rule`).
  - EXPRESS schema functions used by WHERE/RULE expressions (codegen + runtime evaluators).
- **UNIQUE constraints**
  - Entity UNIQUE clauses (e.g. `IfcRoot.GlobalId`) (`ifc-unique`).
  - Aggregate `OF UNIQUE` element uniqueness (`ifc-unique-aggregate`).
- **Aggregate bounds** — validate declared sizes such as `LIST [1:3]`, `SET [0:?]`, `ARRAY [1:2]` (`ifc-aggregate-bounds`).
- **Derived attribute checks** — `*` on DERIVE attributes: Type WHERE and aggregate bounds run on computed DERIVE values where evaluators exist.
- **`ifc-eval-unknown`** — report evaluation failures (parameter conversion, DERIVE errors, UNIQUE key failures) instead of silently skipping; documented in README diagnostic table.
- Tests: `unique_rule_test.go`, `aggregate_bounds_test.go`, `validator_where_type_test.go`, `eval_unknown_test.go`, schema function regression tests.

### Changed

- **`$` vs `*`** — optional `$` skips Type WHERE (absent value); derived `*` is validated via DERIVE evaluation, not skipped like `$`.
- **EXPRESS `?` and `$`** — compile to `wvAbsentVal()` (was `unknown`); aligns with absent-parameter semantics in WHERE/DERIVE.
- **WHERE three-valued logic** — `unknown`/`absent` are falsy in `wvTruthy`; `wvAnd`/`wvOr`/`wvXor`/`wvNot` use Kleene logic so indeterminate conditions no longer pass by default.
- **Schema function codegen** — emit all 48 IFC schema functions (including unreferenced ones such as `IfcVectorSum`); 192 DERIVE slots with 50 shared evaluators.
- **Unresolved `#` references** — reported by the linker only (`Could not resolve reference to '#…'`); validator no longer emits duplicate `ifc-eval-unknown` for the same unresolved reference.
- Entity WHERE / Global RULE collect real evaluation failures during rule checks; EXPRESS type-narrowing paths (`QualEntity` mismatch, empty inverse lists) remain silent `unknown` for Kleene logic.
- **README** — documents full EXPRESS validation coverage, diagnostic codes, and updated limitations (eval-unknown vs linker, not a geometry/IDS cert tool).

### Fixed

- Schema function codegen: `TYPEOF` / `IN TYPEOF` no longer miscompiled to entity subtype shortcuts; `IfcScalarTimesVector`, `IfcUniquePropertySetNames`, and similar functions evaluate correctly.
- `wvEntityInTypeOf` fallback for record-typed WHERE values.
- **QualEntity** — no longer falls back to `SELF` before checking the parameter entity instance; fixes false `ifc-entity-where` on e.g. `#150` / `UniquePropertySetNames`.
- **UNIQUE `GlobalId`** — case-sensitive key comparison; fixes false `ifc-unique` when only letter case differs (e.g. `#64609` / `#65761`).

## [0.0.6] - 2026-06-29

### Added

- VS Code extension publish scripts (`publish:vsce`, `publish:ovx`).

### Changed

- Extension version: `0.0.5` → `0.0.6`.
- Updated `README.md` with initial schema validation section and diagnostic codes.

## [0.0.5] - 2026-06-28 / 2026-06-29

### Added

- **EXPRESS → Go schema generation pipeline**
  - Extract entities, inheritance, attributes, and types from `EXPRESS/t.exp`; emit `schema_gen.go` (876 entities, 436 types, IFC4X3).
  - Orchestrated via `npm run generate:schema` and `ifc/gen.go`.
- **IFC entity instance schema validation** (LSP build Phase 3)
  - Unknown / abstract entities (`ifc-unknown-entity`, `ifc-abstract-entity`).
  - Parameter count: required missing, trailing OPTIONAL may be omitted (`ifc-param-count`).
  - Parameter types and entity reference constraints (`ifc-param-type`, `ifc-entity-type`).
  - Enum value validation (`ifc-enum-value`).
  - TYPED_PARAMETER keyword vs attribute type matching (`ifc-typed-keyword`).
  - LIST/SET aggregate element validation; nested LIST (e.g. `IFCCARTESIANPOINTLIST2D`).
  - DERIVE placeholder `*` vs optional omission `$` (`ifc-param-derived`, `ifc-param-required`).
- Schema metadata: `Derived`, `EnumValues`, aggregate type aliases (e.g. `IfcCompoundPlaneAngleMeasure`).
- Validation implementation: `validator.go`, `validator_types.go`, `validator_deep.go`, `validator_test.go`.

### Fixed

- `IFCDIRECTION`: DERIVE attributes are excluded from STEP parameter count.
- `IFCSIUNIT`: first parameter `*` denotes derived `Dimensions`; no longer reported as a missing required value.
- `IFCRATIOMEASURE` / `IfcValue`: nested SELECT member matching (e.g. `IFCMEASUREWITHUNIT`).
- `IFCSITE` lat/long: `IfcCompoundPlaneAngleMeasure` recognized as `LIST [3:4] OF INTEGER` aggregate type.
- Aggregate list elements: EXPRESS built-in types such as `INTEGER` are recognized correctly.
- `IFCCARTESIANPOINTLIST2D`: nested LIST element type resolved via the first ` OF ` segment.

### Known limitations

- EXPRESS WHERE rules and GLOBAL functions are not evaluated.
- Validation applies to `#` references within a single file; cross-file analysis depends on workspace layout.

## [0.0.4] - 2026-06-26

### Added

- `LICENSE`.

### Changed

- Improved `package.json` description and keywords (BIM, STEP, LSP, etc.).

## [0.0.3] - 2026-06-26

### Changed

- Updated extension icon and `package.json` version.

## [0.0.2] - 2026-06-26

### Added

- Extension icon assets.

## [0.0.1] - 2026-06-26

### Added

- Initial IFC language service scaffold based on [fastbelt](https://github.com/TypeFox/fastbelt).
- `ifc.fb`: ISO 10303-21 STEP physical file grammar.
- `go generate` / fastbelt-generated lexer, parser, AST, and completion.
- `cmd/ifc-lsp`: stdio JSON-RPC language server.
- Extension: syntax highlighting, bracket matching, LSP client.
- Core LSP features: parse diagnostics, Go to Definition, Find References, Completion.
