package ifc

import (
	"os"
	"testing"
)

func TestZZRoots(t *testing.T) {
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	fixture, sc := newIfcSymbolFixture(t)
	doc := fixture.Parse(string(raw))
	if doc.Document.Root == nil {
		t.Fatal("no AST")
	}
	symbols := documentSymbols(t, sc, doc)
	for i, s := range symbols {
		t.Logf("root[%d] name=%q detail=%q children=%d", i, s.Name, s.Detail, len(s.Children))
	}
}
