package ifc

import (
	"context"
	"fmt"
	"strings"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

type IfcDocumentValidator struct {
	inner workspace.DocumentValidator
	sc    *service.Container
}

func NewIfcDocumentValidator(sc *service.Container) workspace.DocumentValidator {
	return &IfcDocumentValidator{
		inner: workspace.NewDefaultDocumentValidator(sc),
		sc:    sc,
	}
}

func (v *IfcDocumentValidator) Validate(ctx context.Context, doc *core.Document, level string) []*core.Diagnostic {
	pop := BuildPopulation(ctx, doc)
	ctx = WithPopulation(ctx, pop)
	ctx = WithEntityEvalCache(ctx)
	settings := ValidationSettingsFromService(v.sc)
	ctx = WithValidationSettings(ctx, settings)
	diags := v.inner.Validate(ctx, doc, level)
	if doc != nil && doc.Root != nil && !documentHasParseErrors(doc) {
		var accept core.ValidationAcceptor = func(d *core.Diagnostic) {
			diags = append(diags, d)
		}
		if settings.Rules {
			validateGlobalRules(ctx, doc.Root, accept)
			validateUniqueRules(ctx, accept)
		}
	}
	return diags
}

func documentHasParseErrors(doc *core.Document) bool {
	return len(doc.LexerErrors) > 0 || len(doc.ParserErrors) > 0
}

func validationAnchorNode(root core.AstNode) (core.AstNode, *core.Token) {
	file, ok := root.(EXCHANGE_FILE)
	if !ok || file == nil {
		return root, nil
	}
	header := file.HeaderSection()
	if header == nil {
		return root, nil
	}
	for _, entity := range header.HeaderEntityList() {
		if entity == nil {
			continue
		}
		if strings.EqualFold(entity.Keyword(), "FILE_SCHEMA") {
			return entity, entity.KeywordToken()
		}
	}
	return root, nil
}

func validateEntityWhereRules(
	ctx context.Context,
	owner core.AstNode,
	stepName string,
	self SIMPLE_ENTITY_INSTANCE,
	params []PARAMETER,
	accept core.ValidationAcceptor,
) {
	if !validationEnabled(ctx, LayerEntityWhere) {
		return
	}
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
