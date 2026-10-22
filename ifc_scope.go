package ifc

import (
	"context"
	"reflect"
	"sync"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/linking"
)

// ifcRefScopeDataKey caches the per-document reference scope on Document.Data.
const ifcRefScopeDataKey = "ifc.referenceScope"

// ifcScopeBuildMu serializes scope construction. The linker resolves all
// references of a document in parallel, so without this guard every goroutine
// that misses the cache at the same time would build the full symbol index.
var ifcScopeBuildMu sync.Mutex

// ifcScopeEntry pairs a cached scope with the LocalSymbols it was built from,
// so a rebuild that recomputes LocalSymbols on the same Document invalidates
// the cache. It is stored as a pointer in Document.Data: type-asserting a
// pointer out of an interface does not allocate, unlike a struct value.
type ifcScopeEntry struct {
	localSymbols core.LocalSymbols
	scope        core.Scope
}

// ifcScopeProvider resolves #N entity references through a per-document name
// index instead of the fastbelt default (linking.DefaultScopeOfType).
//
// The default walks the container chain and linearly scans each level's
// symbols, comparing every symbol name for every reference. STEP files have a
// flat namespace: every entity instance is a symbol in one container, so with
// ~10^5 references against ~10^4 symbols this is O(refs x symbols) string
// comparisons and dominates the linking phase.
//
// This provider builds a MapScope (name-indexed multimap) once per document
// build and caches it on Document.Data:
//   - LocalSymbols are merged across all containers into one flat index.
//     #N names are unique within a STEP file, so flattening is semantically
//     equivalent to the container-chain walk for this grammar.
//   - Imported symbols (cross-file references) form the outer scope.
//
// The cache is safe across rebuilds: workspace updates create a fresh
// *core.Document per build (fresh Data), and the LocalSymbols pointer check
// catches in-place resets that recompute symbols on the same Document.
type ifcScopeProvider struct{}

// ScopeENTITY_REFERENCEValue implements IfcScopeProvider.
func (ifcScopeProvider) ScopeENTITY_REFERENCEValue(_ context.Context, reference *core.Reference[SIMPLE_ENTITY_INSTANCE]) core.Scope {
	doc := reference.Owner().Document()
	if doc == nil {
		return core.EmptyScope
	}
	// Fast path: lock-free cache read.
	if v, ok := doc.Data.Load(ifcRefScopeDataKey); ok {
		if entry, ok := v.(*ifcScopeEntry); ok && entry.scope != nil && entry.localSymbols == doc.LocalSymbols {
			return entry.scope
		}
	}
	// Slow path: serialize construction so concurrent reference resolution
	// (the linker resolves references in parallel) builds the index once.
	ifcScopeBuildMu.Lock()
	defer ifcScopeBuildMu.Unlock()
	if v, ok := doc.Data.Load(ifcRefScopeDataKey); ok {
		if entry, ok := v.(*ifcScopeEntry); ok && entry.scope != nil && entry.localSymbols == doc.LocalSymbols {
			return entry.scope
		}
	}
	scope := buildIfcDocumentScope(doc)
	doc.Data.Store(ifcRefScopeDataKey, &ifcScopeEntry{localSymbols: doc.LocalSymbols, scope: scope})
	return scope
}

func buildIfcDocumentScope(doc *core.Document) core.Scope {
	var outer core.Scope = core.EmptyScope
	if doc.ImportedSymbols != nil {
		outer = core.NewMapScopeFromSeq(
			doc.ImportedSymbols.ForType(reflect.TypeFor[SIMPLE_ENTITY_INSTANCE]()),
			nil,
		)
	}
	return core.NewMapScopeFromSeq(allIfcLocalSymbols(doc), outer)
}

// allIfcLocalSymbols yields every local symbol of the document across all
// symbol containers.
func allIfcLocalSymbols(doc *core.Document) core.SymbolSeq {
	return func(yield func(*core.SymbolDescription) bool) {
		if doc == nil || doc.LocalSymbols == nil {
			return
		}
		// Fast path: the default provider stores containers in an exported map.
		if defaults, ok := doc.LocalSymbols.(*linking.DefaultLocalSymbols); ok {
			for _, container := range defaults.Symbols {
				if container == nil {
					continue
				}
				for desc := range container.All() {
					if !yield(desc) {
						return
					}
				}
			}
			return
		}
		// Fallback: walk the AST and collect each distinct container once.
		if doc.Root == nil {
			return
		}
		seen := make(map[core.SymbolContainer]struct{})
		for node := range core.AllNodes(doc.Root) {
			container := doc.LocalSymbols.For(node)
			if container == nil {
				continue
			}
			if _, dup := seen[container]; dup {
				continue
			}
			seen[container] = struct{}{}
			for desc := range container.All() {
				if !yield(desc) {
					return
				}
			}
		}
	}
}
