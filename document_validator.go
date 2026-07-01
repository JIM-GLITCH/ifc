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
	ok, detail := CheckEntityWhereRules(ctx, stepName, self, params, pop)
	if ok {
		return
	}
	msg := detail
	if msg == "" {
		msg = fmt.Sprintf("Entity %q violates a WHERE constraint", stepName)
	} else {
		msg = fmt.Sprintf("Entity %q WHERE constraint violated: %s", stepName, detail)
	}
	token := self.NameToken()
	opts := []core.DiagnosticOption{core.WithCode("ifc-entity-where")}
	if token != nil {
		opts = append(opts, core.WithToken(token))
	}
	accept(core.NewDiagnostic(core.SeverityError, msg, owner, opts...))
}
