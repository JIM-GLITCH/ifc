package ifc

import (
	"context"

	core "typefox.dev/fastbelt"
)

func validateEntityInstance(ctx context.Context, owner core.AstNode, recordKeyword string, keywordToken *core.Token, params []PARAMETER, accept core.ValidationAcceptor) {
	def, ok := LookupEntity(recordKeyword)
	if !ok {
		accept(core.NewDiagnostic(
			core.SeverityError,
			"Unknown entity "+recordKeyword+" (not defined in schema "+SchemaName+")",
			owner,
			core.WithToken(keywordToken),
			core.WithCode("ifc-unknown-entity"),
		))
		return
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

	paramCount := len(params)
	if valid, message := validateParameterCount(def, paramCount); !valid {
		accept(core.NewDiagnostic(
			core.SeverityError,
			message,
			owner,
			core.WithToken(keywordToken),
			core.WithCode("ifc-param-count"),
		))
		return
	}

	validateParameterTypes(ctx, owner, def, params, accept)
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
