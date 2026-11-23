package ifc

import (
	"context"
	"strings"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/server"
	"typefox.dev/fastbelt/textdoc"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
	"typefox.dev/lsp"
)

type IfcHoverProvider struct {
	sc *service.Container
}

func NewIfcHoverProvider(sc *service.Container) server.HoverProvider {
	return &IfcHoverProvider{sc: sc}
}

// RequiredState matches the default hover provider: entity hover resolves
// references, so it needs the document linked. Scoping to the single document
// keeps hover from waiting on the whole workspace build.
func (p *IfcHoverProvider) RequiredState() (core.DocumentState, bool) {
	return core.DocStateLinked, false
}

func (p *IfcHoverProvider) HandleHoverRequest(ctx context.Context, params *lsp.HoverParams) (*lsp.Hover, error) {
	documentManager := p.sc.MustGet[workspace.DocumentManager]()
	doc := documentManager.Get(core.ParseURI(string(params.TextDocument.URI)))
	if doc == nil || doc.TextDoc == nil {
		return nil, nil
	}

	offset := doc.TextDoc.OffsetAt(params.Position)
	first, second := doc.Tokens.SearchOffset2(offset)
	if first == nil {
		return nil, nil
	}

	if hover := entityHover(first, second, doc.TextDoc); hover != nil {
		return hover, nil
	}
	if hover := enumHover(first, second, doc.TextDoc); hover != nil {
		return hover, nil
	}
	if hover := derivedHover(ctx, doc, first); hover != nil {
		return hover, nil
	}
	if hover := derivedHover(ctx, doc, second); hover != nil {
		return hover, nil
	}

	nameFinder := p.sc.MustGet[server.NameFinder]()
	foundName := nameFinder.Find(ctx, first, second)
	if foundName.Target == nil || foundName.Source == nil || foundName.Source == foundName.Target {
		return nil, nil
	}

	target := foundName.Target.Owner()
	if target == nil {
		return nil, nil
	}
	if _, ok := target.(ENTITY_INSTANCE); !ok {
		return nil, nil
	}

	preview := strings.TrimSpace(target.Text())
	if preview == "" {
		return nil, nil
	}

	return &lsp.Hover{
		Contents: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: "```ifc\n" + preview + "\n```",
		},
		Range: foundName.Source.TextRange().LspRange(doc.TextDoc),
	}, nil
}

func entityHover(first, second *core.Token, handle textdoc.Handle) *lsp.Hover {
	for _, token := range []*core.Token{first, second} {
		if token == nil || token.Type != Token_KEYWORD {
			continue
		}
		record, ok := token.Element.(SIMPLE_RECORD)
		if !ok {
			continue
		}
		stepName := strings.ToUpper(record.Keyword())
		params, ok := entitySignatureParameters(stepName)
		if !ok {
			declaration, ok := entityExpressDocs[stepName]
			if !ok {
				continue
			}
			return &lsp.Hover{
				Contents: lsp.MarkupContent{
					Kind:  lsp.Markdown,
					Value: "```express\n" + declaration + "\n```",
				},
				Range: token.TextRange().LspRange(handle),
			}
		}

		var b strings.Builder
		b.WriteString("`" + entitySignatureLabel(stepName, params, false) + "`")
		if types := entityParameterTypeMarkdown(params); types != "" {
			b.WriteString("\n\n" + types)
		}
		return &lsp.Hover{
			Contents: lsp.MarkupContent{
				Kind:  lsp.Markdown,
				Value: b.String(),
			},
			Range: token.TextRange().LspRange(handle),
		}
	}
	return nil
}
