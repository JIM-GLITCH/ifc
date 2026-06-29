package ifc

import (
	"context"

	core "typefox.dev/fastbelt"
)

func validateEntityInstance(ctx context.Context, owner core.AstNode, recordKeyword string, keywordToken *core.Token, params []PARAMETER, accept core.ValidationAcceptor) {
	def, ok := validateEntityName(owner, recordKeyword, keywordToken, accept)
	if !ok {
		return
	}
	if !validateEntityParameters(ctx, owner, def, params, keywordToken, accept) {
		return
	}
	validateEntityWhereConstraints(ctx, owner, recordKeyword, params, accept)
}

func validateEntityName(owner core.AstNode, recordKeyword string, keywordToken *core.Token, accept core.ValidationAcceptor) (EntityDef, bool) {
	def, ok := LookupEntity(recordKeyword)
	if !ok {
		accept(core.NewDiagnostic(
			core.SeverityError,
			"Unknown entity "+recordKeyword+" (not defined in schema "+SchemaName+")",
			owner,
			core.WithToken(keywordToken),
			core.WithCode("ifc-unknown-entity"),
		))
		return EntityDef{}, false
	}
	if def.Abstract {
		accept(core.NewDiagnostic(
			core.SeverityWarning,
			"Entity "+def.ExpressName+" is abstract in schema "+SchemaName,
			owner,
			core.WithToken(keywordToken),
			core.WithCode("ifc-abstract-entity"),
		))
	}
	return def, true
}

func validateEntityParameters(ctx context.Context, owner core.AstNode, def EntityDef, params []PARAMETER, keywordToken *core.Token, accept core.ValidationAcceptor) bool {
	if valid, message := validateParameterCount(def, len(params)); !valid {
		accept(core.NewDiagnostic(
			core.SeverityError,
			message,
			owner,
			core.WithToken(keywordToken),
			core.WithCode("ifc-param-count"),
		))
		return false
	}
	validateParameterTypes(ctx, owner, def, params, accept)
	return true
}

func validateEntityWhereConstraints(ctx context.Context, owner core.AstNode, recordKeyword string, params []PARAMETER, accept core.ValidationAcceptor) {
	simple, ok := owner.(SIMPLE_ENTITY_INSTANCE)
	if !ok {
		return
	}
	validateEntityWhereRules(ctx, owner, recordKeyword, simple, params, accept)
}

func (i *SIMPLE_ENTITY_INSTANCEImpl) Validate(ctx context.Context, _ string, accept core.ValidationAcceptor) {
	record := i.Record()
	if record == nil {
		return
	}
	var params []PARAMETER
	if paramList := record.Parameters(); paramList != nil {
		params = paramList.Parameters()
	}
	validateEntityInstance(ctx, i, record.Keyword(), record.KeywordToken(), params, accept)
}

func (i *COMPLEX_ENTITY_INSTANCEImpl) Validate(ctx context.Context, _ string, accept core.ValidationAcceptor) {
	record := i.Record()
	if record == nil {
		return
	}
	for _, simpleRecord := range record.Records() {
		if simpleRecord == nil {
			continue
		}
		var params []PARAMETER
		if paramList := simpleRecord.Parameters(); paramList != nil {
			params = paramList.Parameters()
		}
		validateEntityInstance(ctx, i, simpleRecord.Keyword(), simpleRecord.KeywordToken(), params, accept)
	}
}
