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
	sc *service.Container
}

func NewIfcDocumentValidator(sc *service.Container) workspace.DocumentValidator {
	return &IfcDocumentValidator{
		sc: sc,
	}
}

func (v *IfcDocumentValidator) Validate(ctx context.Context, doc *core.Document, level string) []*core.Diagnostic {
	pop := BuildPopulation(ctx, doc)
	ctx = WithPopulation(ctx, pop)
	ctx = WithEntityEvalCache(ctx)
	settings := ValidationSettingsFromService(v.sc)
	ctx = WithValidationSettings(ctx, settings)

	// Lexer/parser/linker diagnostics (equivalent to the fastbelt default
	// validator's preamble), followed by the parallel AST validation walk.
	source := validationDiagnosticSource(v.sc, doc)
	diags := make([]*core.Diagnostic, 0, 64)
	diags = append(diags, workspace.CreateLexerDiagnostics(doc, source)...)
	diags = append(diags, workspace.CreateParserDiagnostics(doc, source)...)
	diags = append(diags, workspace.CreateLinkerDiagnostics(doc, source)...)
	diags = append(diags, parallelValidate(ctx, doc, level, source)...)

	if doc != nil && doc.Root != nil && !documentHasParseErrors(doc) {
		var accept core.ValidationAcceptor = func(d *core.Diagnostic) {
			if d.Source == "" {
				d.Source = source
			}
			diags = append(diags, d)
		}
		if settings.Rules {
			validateGlobalRules(ctx, doc.Root, accept)
			validateUniqueRules(ctx, accept)
		}
	}
	return diags
}

// validationDiagnosticSource mirrors the fastbelt default validator's source
// resolution: the document's language id, falling back to the workspace
// language id.
func validationDiagnosticSource(sc *service.Container, doc *core.Document) string {
	if doc != nil && doc.TextDoc != nil {
		if source := doc.TextDoc.LanguageID(); source != "" {
			return source
		}
	}
	if languageID, err := service.Get[workspace.LanguageID](sc); err == nil {
		return string(languageID)
	}
	return ""
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
