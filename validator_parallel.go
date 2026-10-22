package ifc

import (
	"context"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"

	core "typefox.dev/fastbelt"
)

// parallelValidateThreshold is the number of validator nodes below which
// validation runs sequentially; goroutine setup would dominate the work.
const parallelValidateThreshold = 64

// parallelValidate collects every core.Validator node of the document and
// validates them across worker goroutines.
//
// Entity validation is embarrassingly parallel with respect to results: each
// entity's checks read only its own AST parameters, the read-only Population
// (name index + inverse links), other entities' raw AST via reference
// resolution, and the immutable schema. No entity consumes another entity's
// diagnostics, so any work distribution is correct.
//
// The shared entity eval cache is the only mutable state: it is a plain map
// keyed by (entity, attribute) whose values are deterministic, so each worker
// gets its own cache. The within-entity hit rate (several rules of one entity
// reading the same attribute) is fully preserved; cross-entity hits are lost,
// which only re-runs a cheap parameter conversion.
//
// Diagnostics from all workers are merged in a deterministic order (by source
// position, then code, then message) so output does not depend on scheduling.
func parallelValidate(ctx context.Context, doc *core.Document, level, source string) []*core.Diagnostic {
	diagnostics := []*core.Diagnostic{}
	if doc == nil || doc.Root == nil {
		return diagnostics
	}

	var nodes []core.Validator
	for node := range core.AllNodes(doc.Root) {
		if validator, ok := node.(core.Validator); ok {
			nodes = append(nodes, validator)
		}
	}
	if len(nodes) == 0 {
		return diagnostics
	}

	newAccept := func(out *[]*core.Diagnostic) core.ValidationAcceptor {
		return func(d *core.Diagnostic) {
			if d.Source == "" {
				d.Source = source
			}
			*out = append(*out, d)
		}
	}

	workers := runtime.GOMAXPROCS(0)
	if workers > len(nodes) {
		workers = len(nodes)
	}
	if len(nodes) < parallelValidateThreshold || workers <= 1 {
		for _, node := range nodes {
			if ctx.Err() != nil {
				break
			}
			accept := newAccept(&diagnostics)
			node.Validate(ctx, level, accept)
		}
		sortDiagnostics(diagnostics)
		return diagnostics
	}

	perWorker := make([][]*core.Diagnostic, workers)
	var next atomic.Int64
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			// Per-worker eval cache: the shared map is not safe for concurrent
			// writes, and cached values are deterministic so isolation costs
			// nothing but cross-entity memo hits.
			workerCtx := context.WithValue(ctx, entityEvalCacheKey{}, newEntityEvalCache())
			out := perWorker[w]
			accept := newAccept(&out)
			for {
				i := int(next.Add(1)) - 1
				if i >= len(nodes) {
					break
				}
				if ctx.Err() != nil {
					break
				}
				nodes[i].Validate(workerCtx, level, accept)
			}
			perWorker[w] = out
		}(w)
	}
	wg.Wait()

	for _, workerDiags := range perWorker {
		diagnostics = append(diagnostics, workerDiags...)
	}
	sortDiagnostics(diagnostics)
	return diagnostics
}

// sortDiagnostics orders diagnostics by source position, then code and message
// as tiebreakers, making the merged output independent of worker scheduling.
func sortDiagnostics(diags []*core.Diagnostic) {
	sort.SliceStable(diags, func(i, j int) bool {
		a, b := diags[i], diags[j]
		if a.Range.Start.Line != b.Range.Start.Line {
			return a.Range.Start.Line < b.Range.Start.Line
		}
		if a.Range.Start.Column != b.Range.Start.Column {
			return a.Range.Start.Column < b.Range.Start.Column
		}
		if a.Range.End.Line != b.Range.End.Line {
			return a.Range.End.Line < b.Range.End.Line
		}
		if a.Range.End.Column != b.Range.End.Column {
			return a.Range.End.Column < b.Range.End.Column
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		return a.Message < b.Message
	})
}
