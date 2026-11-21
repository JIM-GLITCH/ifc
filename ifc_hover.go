package ifc

import (
	"context"
	"strings"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/server"
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

func (p *IfcHoverProvider) HandleHoverRequest(ctx context.Context, params *lsp.HoverParams) (*lsp.Hover, error) {
	documentManager := service.MustGet[workspace.DocumentManager](p.sc)
	doc := documentManager.Get(core.ParseURI(string(params.TextDocument.URI)))
	if doc == nil {
		return nil, nil
	}

	offset := doc.TextDoc.OffsetAt(params.Position)
	first, second := doc.Tokens.SearchOffset2(offset)
	if first == nil {
		return nil, nil
	}

	nameFinder := service.MustGet[server.NameFinder](p.sc)
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
		Range: foundName.Source.Segment().Range.LspRange(),
	}, nil
}
