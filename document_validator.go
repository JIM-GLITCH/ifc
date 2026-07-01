package ifc

import (
	"context"
	"fmt"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

type IfcDocumentValidator struct {
	inner workspace.DocumentValidator
}

func NewIfcDocumentValidator(sc *service.Container) workspace.DocumentValidator {
	return &IfcDocumentValidator{
		inner: workspace.NewDefaultDocumentValidator(sc),
	}
}

func (v *IfcDocumentValidator) Validate(ctx context.Context, doc *core.Document, level string) []*core.Diagnostic {
	pop := BuildPopulation(ctx, doc)
	ctx = WithPopulation(ctx, pop)
	diags := v.inner.Validate(ctx, doc, level)
	if doc != nil && doc.Root != nil {
		var accept core.ValidationAcceptor = func(d *core.Diagnostic) {
			diags = append(diags, d)
		}
		validateGlobalRules(ctx, doc.Root, accept)
		validateUniqueRules(ctx, accept)
	}
	return diags
}

func validateEntityWhereRules(
	ctx context.Context,
	owner core.AstNode,
	stepName string,
	self SIMPLE_ENTITY_INSTANCE,
	params []PARAMETER,
	accept core.ValidationAcceptor,
) {
	pop := PopulationFromContext(ctx)
	ok, detail, evalUnknowns := CheckEntityWhereRules(ctx, stepName, self, params, pop)
	token := self.NameToken()
	acceptEvalUnknowns(accept, owner, token, evalUnknowns)
	if ok {
		return
	}
	if detail == "" {
		return
	}
	msg := fmt.Sprintf("Entity %q WHERE constraint violated: %s", stepName, detail)
	opts := []core.DiagnosticOption{core.WithCode("ifc-entity-where")}
	if token != nil {
		opts = append(opts, core.WithToken(token))
	}
	accept(core.NewDiagnostic(core.SeverityError, msg, owner, opts...))
}
