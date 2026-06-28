package ifc

import (
	"context"
	"fmt"

	core "typefox.dev/fastbelt"
)

func countParameters(params []PARAMETER) int {
	count := 0
	for _, param := range params {
		count += parameterSlotCount(param)
	}
	return count
}

func parameterSlotCount(param PARAMETER) int {
	if param == nil {
		return 0
	}
	switch p := param.(type) {
	case LIST:
		// A LIST is one STEP parameter.
		return 1
	case TYPED_PARAMETER:
		return 1
	default:
		_ = p
		return 1
	}
}

func validateParameterCount(def EntityDef, paramCount int) (ok bool, message string) {
	attrs := def.Attrs
	maxCount := len(attrs)
	if paramCount > maxCount {
		return false, fmt.Sprintf("Entity %q accepts at most %d parameter(s), got %d", def.ExpressName, maxCount, paramCount)
	}
	for i := 0; i < maxCount; i++ {
		if !attrs[i].Optional && i >= paramCount {
			return false, fmt.Sprintf("Entity %q is missing required parameter at position %d (got %d parameter(s))", def.ExpressName, i+1, paramCount)
		}
	}
	return true, ""
}

func validateEntityInstance(owner core.AstNode, recordKeyword string, keywordToken *core.Token, paramCount int, accept core.ValidationAcceptor) {
	def, ok := LookupEntity(recordKeyword)
	if !ok {
		accept(core.NewDiagnostic(
			core.SeverityError,
			fmt.Sprintf("Unknown entity %q (not defined in schema %s)", recordKeyword, SchemaName),
			owner,
			core.WithToken(keywordToken),
			core.WithCode("ifc-unknown-entity"),
		))
		return
	}

	if def.Abstract {
		accept(core.NewDiagnostic(
			core.SeverityWarning,
			fmt.Sprintf("Entity %q is abstract in schema %s", def.ExpressName, SchemaName),
			owner,
			core.WithToken(keywordToken),
			core.WithCode("ifc-abstract-entity"),
		))
	}

	if valid, message := validateParameterCount(def, paramCount); !valid {
		accept(core.NewDiagnostic(
			core.SeverityError,
			message,
			owner,
			core.WithToken(keywordToken),
			core.WithCode("ifc-param-count"),
		))
	}
}

func (i *SIMPLE_ENTITY_INSTANCEImpl) Validate(_ context.Context, _ string, accept core.ValidationAcceptor) {
	record := i.Record()
	if record == nil {
		return
	}
	params := record.Parameters()
	paramCount := 0
	if params != nil {
		paramCount = countParameters(params.Parameters())
	}
	validateEntityInstance(i, record.Keyword(), record.KeywordToken(), paramCount, accept)
}

func (i *COMPLEX_ENTITY_INSTANCEImpl) Validate(_ context.Context, _ string, accept core.ValidationAcceptor) {
	record := i.Record()
	if record == nil {
		return
	}
	for _, simpleRecord := range record.Records() {
		if simpleRecord == nil {
			continue
		}
		params := simpleRecord.Parameters()
		paramCount := 0
		if params != nil {
			paramCount = countParameters(params.Parameters())
		}
		validateEntityInstance(i, simpleRecord.Keyword(), simpleRecord.KeywordToken(), paramCount, accept)
	}
}
