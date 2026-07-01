package ifc

import (
	"context"
	"fmt"
	"strings"

	core "typefox.dev/fastbelt"
)

const evalUnknownCode = "ifc-eval-unknown"

func acceptEvalUnknown(
	accept core.ValidationAcceptor,
	node core.AstNode,
	token *core.Token,
	message string,
) {
	if accept == nil || strings.TrimSpace(message) == "" {
		return
	}
	opts := append(diagnosticTokenOptions(token), core.WithCode(evalUnknownCode))
	accept(core.NewDiagnostic(core.SeverityError, message, node, opts...))
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
