package ifc

import (
	"context"
	"fmt"
	"strings"

	core "typefox.dev/fastbelt"
)

const evalUnknownCode = "ifc-eval-unknown"

func shouldSuppressEvalUnknownMessage(message string) bool {
	return strings.HasPrefix(strings.TrimSpace(message), "Unresolved entity reference")
}

func acceptEvalUnknown(
	accept core.ValidationAcceptor,
	node core.AstNode,
	token *core.Token,
	message string,
) {
	if accept == nil || strings.TrimSpace(message) == "" || shouldSuppressEvalUnknownMessage(message) {
		return
	}
	anchorNode, anchorToken := evalUnknownAnchor(node, token)
	if anchorNode == nil && anchorToken == nil {
		return
	}
	opts := append(diagnosticTokenOptions(anchorToken), core.WithCode(evalUnknownCode))
	accept(core.NewDiagnostic(core.SeverityError, message, anchorNode, opts...))
}

const maxEvalUnknownLineSpan = 50

func evalUnknownAnchor(node core.AstNode, token *core.Token) (core.AstNode, *core.Token) {
	if token != nil {
		return node, token
	}
	if node != nil && isDocumentRootNode(node) {
		return nil, nil
	}
	if inst, ok := node.(SIMPLE_ENTITY_INSTANCE); ok && inst != nil {
		if t := inst.NameToken(); t != nil {
			return inst, t
		}
	}
	if node != nil && diagnosticLineSpan(node) <= maxEvalUnknownLineSpan {
		return node, nil
	}
	return nil, nil
}

func isDocumentRootNode(node core.AstNode) bool {
	switch node.(type) {
	case EXCHANGE_FILE:
		return true
	default:
		return false
	}
}

func diagnosticLineSpan(node core.AstNode) int {
	if node == nil {
		return 0
	}
	seg := node.Segment()
	if seg == nil {
		return 0
	}
	return int(seg.Range.End.Line) - int(seg.Range.Start.Line)
}

func acceptEvalUnknowns(
	accept core.ValidationAcceptor,
	node core.AstNode,
	token *core.Token,
	messages []string,
) {
	for _, msg := range messages {
		acceptEvalUnknown(accept, node, token, msg)
	}
}

func parameterBlockedByUnresolvedEntityRef(ctx context.Context, pop *Population, param PARAMETER) bool {
	if param == nil {
		return false
	}
	switch p := param.(type) {
	case ENTITY_REFERENCE:
		if p.Value() == nil {
			return false
		}
		return resolveEntityReference(ctx, pop, p) == nil
	case LIST:
		for _, el := range p.Parameters() {
			if _, ok := parameterToWhereValueCtx(ctx, pop, el); !ok {
				return parameterBlockedByUnresolvedEntityRef(ctx, pop, el)
			}
		}
		return false
	case TYPED_PARAMETER:
		if inner := p.Parameter(); inner != nil {
			return parameterBlockedByUnresolvedEntityRef(ctx, pop, inner)
		}
		return false
	default:
		return false
	}
}

func parameterEvalFailureReason(ctx context.Context, pop *Population, param PARAMETER) string {
	if param == nil {
		return "Cannot evaluate parameter value"
	}
	switch p := param.(type) {
	case ENTITY_REFERENCE:
		ref := p.Value()
		if ref != nil {
			text := strings.TrimSpace(ref.Text())
			if text != "" {
				return fmt.Sprintf("Unresolved entity reference %s", text)
			}
		}
		return "Unresolved entity reference"
	case LIST:
		for i, el := range p.Parameters() {
			if _, ok := parameterToWhereValueCtx(ctx, pop, el); !ok {
				return fmt.Sprintf("Cannot evaluate list element %d for constraint check", i+1)
			}
		}
		return "Cannot evaluate list value for constraint check"
	case TYPED_PARAMETER:
		if inner := p.Parameter(); inner != nil {
			return parameterEvalFailureReason(ctx, pop, inner)
		}
		return "Cannot evaluate typed parameter value"
	default:
		return "Cannot evaluate parameter value"
	}
}
