package ifc

import (
	"context"

	"typefox.dev/fastbelt/server"
	"typefox.dev/lsp"
)

// noopFoldingProvider disables folding ranges. IFC files are dominated by very
// long lines, and the default provider converts every AST node's byte range to
// an LSP position by rescanning the whole line in UTF-16 units. That O(nodes ×
// lineLength) work monopolizes a read lock and makes hover/diagnostics stall.
type noopFoldingProvider struct{}

var _ server.FoldingRangeProvider = (*noopFoldingProvider)(nil)

func (*noopFoldingProvider) HandleFoldingRangeRequest(context.Context, *lsp.FoldingRangeParams) ([]lsp.FoldingRange, error) {
	return nil, nil
}
