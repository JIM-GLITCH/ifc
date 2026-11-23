package ifc

import (
	"context"
	"strings"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
	"typefox.dev/lsp"
)

// IfcInlayHintProvider shows IFC attribute names inline before each
// positional argument of an entity instance, e.g.
// `#1=IFCAXIS2PLACEMENT3D(Location: #2, Axis: $, RefDirection: $)`.
type IfcInlayHintProvider struct {
	sc *service.Container
}

func NewIfcInlayHintProvider(sc *service.Container) *IfcInlayHintProvider {
	return &IfcInlayHintProvider{sc: sc}
}

func (p *IfcInlayHintProvider) HandleInlayHintRequest(ctx context.Context, params *lsp.InlayHintParams) ([]lsp.InlayHint, error) {
	documentManager := p.sc.MustGet[workspace.DocumentManager]()
	doc := documentManager.Get(core.ParseURI(string(params.TextDocument.URI)))
	if doc == nil || doc.TextDoc == nil {
		return nil, nil
	}

	tokens := doc.Tokens
	if len(tokens) == 0 {
		return nil, nil
	}

	rangeStart := doc.TextDoc.OffsetAt(params.Range.Start)
	rangeEnd := doc.TextDoc.OffsetAt(params.Range.End)

	var hints []lsp.InlayHint
	for i := 0; i+1 < len(tokens); i++ {
		token := &tokens[i]
		if token.Type != Token_KEYWORD || tokens[i+1].Image != "(" {
			continue
		}
		keyword := strings.ToUpper(strings.TrimPrefix(token.Image, "!"))
		sigParams, ok := entitySignatureParameters(keyword)
		if !ok || len(sigParams) == 0 {
			continue
		}

		// Scan the argument list starting after the record's '('. Track
		// parenthesis depth so that commas inside nested lists, typed
		// parameters and nested entity calls do not split arguments.
		argIndex := 0
		argOpen := false
		depth := 0
		for j := i + 2; j < len(tokens); j++ {
			arg := &tokens[j]
			if arg.Image == ")" && depth == 0 {
				break
			}
			switch arg.Image {
			case "(":
				depth++
			case ")":
				depth--
			case ",":
				if depth == 0 {
					argIndex++
					argOpen = false
					continue
				}
			}
			if !argOpen {
				argOpen = true
				if hint, ok := ifcArgumentHint(doc, arg, argIndex, sigParams, rangeStart, rangeEnd); ok {
					hints = append(hints, hint)
				}
			}
		}
	}
	return hints, nil
}

func ifcArgumentHint(doc *core.Document, arg *core.Token, argIndex int, sigParams []entitySignatureParameter, rangeStart, rangeEnd int) (lsp.InlayHint, bool) {
	if argIndex >= len(sigParams) {
		return lsp.InlayHint{}, false
	}
	start := int(arg.TextRange().Start)
	if start < rangeStart || start > rangeEnd {
		return lsp.InlayHint{}, false
	}
	return lsp.InlayHint{
		Position:    doc.TextDoc.PositionAt(start),
		Label:       []lsp.InlayHintLabelPart{{Value: sigParams[argIndex].Name + ": "}},
		Kind:        lsp.Parameter,
		PaddingLeft: true,
	}, true
}
