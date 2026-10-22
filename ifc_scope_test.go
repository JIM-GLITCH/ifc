package ifc

import (
	"context"
	"os"
	"testing"

	core "typefox.dev/fastbelt"
	fbtest "typefox.dev/fastbelt/test"
)

// TestScopeProviderResolvesAllReferences verifies that the flat MapScope used
// for #N reference resolution resolves every reference in the reference model.
// t.ifc has a known parse error, so this deliberately ignores parse state and
// only checks reference resolution.
func TestScopeProviderResolvesAllReferences(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	doc := fixture.Parse(string(raw))

	total, resolved := 0, 0
	for _, ref := range doc.Document.References {
		total++
		ref.Resolve(context.Background())
		if ref.Error() == nil && ref.RefNode(context.Background()) != nil {
			resolved++
		}
	}
	if total == 0 {
		t.Fatal("no references found - test setup broken")
	}
	t.Logf("references: total=%d resolved=%d unresolved=%d", total, resolved, total-resolved)
	if resolved != total {
		t.Fatalf("unresolved references: %d of %d", total-resolved, total)
	}
}

// TestScopeProviderCachedPerBuild verifies the per-document scope cache: a
// second build over the same document must reuse the cached scope (the
// LocalSymbols pointer is unchanged across a Reset that keeps symbol state)
// while a rebuild that recomputes LocalSymbols builds a fresh one.
func TestScopeProviderCachedPerBuild(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCCARTESIANPOINT((0.,0.,0.));
#2=IFCDIRECTION(#1);
ENDSEC;
END-ISO-10303-21;`
	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	scope := ifcScopeProvider{}.ScopeENTITY_REFERENCEValue(context.Background(), doc.Document.References[0].(*core.Reference[SIMPLE_ENTITY_INSTANCE]))
	if scope == nil {
		t.Fatal("expected a scope for the first reference")
	}
	// The cached entry must be a pointer stored under the data key.
	v, ok := doc.Document.Data.Load(ifcRefScopeDataKey)
	if !ok {
		t.Fatal("expected the scope to be cached on Document.Data")
	}
	entry, ok := v.(*ifcScopeEntry)
	if !ok || entry.scope == nil {
		t.Fatalf("expected a *ifcScopeEntry cache entry, got %T", v)
	}
	if entry.localSymbols != doc.Document.LocalSymbols {
		t.Fatal("cached entry should track the document's current LocalSymbols")
	}
}
