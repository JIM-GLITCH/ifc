package ifc

import (
	"context"
	"os"
	"testing"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/lexer"
	"typefox.dev/fastbelt/parser"
	"typefox.dev/fastbelt/server"
	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/lsp"
)

var sinkHover *lsp.Hover

// BenchmarkHoverEntityRef hovers over a "#N" entity reference inside t.ifc
// through the real HoverProvider path.
func benchHoverServices(b *testing.B) *service.Container {
	sc := service.NewContainer()
	SetupServices(sc)
	server.SetupDefaultServices(sc)
	sc.Seal()
	return sc
}

func BenchmarkHoverEntityRef(b *testing.B) {
	sc := benchHoverServices(b)
	fixture := fbtest.New(b, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		b.Fatal(err)
	}
	doc := fixture.Parse(string(raw))
	if doc.Document.Root == nil {
		b.Fatal("no AST")
	}

	// Find an ENTITY_REFERENCE-ish token: a "#<n>" inside a parameter list
	// (not the "#n=" definition itself).
	var offset int = -1
	content := doc.Document.TextDoc.Content()
	for i := 0; i+2 < len(content); i++ {
		if content[i] != '#' {
			continue
		}
		j := i + 1
		for j < len(content) && content[j] >= '0' && content[j] <= '9' {
			j++
		}
		if j > i+1 && j < len(content) && content[j] != '=' {
			offset = i + 1
			break
		}
	}
	if offset < 0 {
		b.Fatal("no entity reference found")
	}
	pos := doc.Document.TextDoc.PositionAt(offset)
	uri := doc.Document.URI.DocumentURI()

	provider := NewIfcHoverProvider(sc)
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkHover, _ = provider.HandleHoverRequest(ctx, &lsp.HoverParams{
			TextDocument: lsp.TextDocumentIdentifier{URI: uri},
			Position:     pos,
		})
	}
}

// BenchmarkHoverComponents isolates the pieces: token lookup, name finder,
// population build.
func BenchmarkHoverComponents(b *testing.B) {
	sc := benchHoverServices(b)
	fixture := fbtest.New(b, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		b.Fatal(err)
	}
	doc := fixture.Parse(string(raw))
	content := doc.Document.TextDoc.Content()
	offset := -1
	for i := 0; i+2 < len(content); i++ {
		if content[i] != '#' {
			continue
		}
		j := i + 1
		for j < len(content) && content[j] >= '0' && content[j] <= '9' {
			j++
		}
		if j > i+1 && j < len(content) && content[j] != '=' {
			offset = i + 1
			break
		}
	}
	first, second := doc.Document.Tokens.SearchOffset2(offset)
	if first == nil {
		b.Fatal("no token")
	}
	ctx := context.Background()

	b.Run("SearchOffset2", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			first, second = doc.Document.Tokens.SearchOffset2(offset)
		}
	})
	b.Run("NameFinder", func(b *testing.B) {
		nf := sc.MustGet[server.NameFinder]()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			found := nf.Find(ctx, first, second)
			_ = found
		}
	})
	b.Run("BuildPopulation", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkPop = BuildPopulation(ctx, doc.Document)
		}
	})
	b.Run("populationForHover", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkPop = populationForHover(ctx, doc.Document)
		}
	})
	var _ = core.TextRange{}
	var _ = service.NewContainer
}

// BenchmarkLexParse measures lexer + parser on t.ifc.
func BenchmarkLexParse(b *testing.B) {
	sc := benchHoverServices(b)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		b.Fatal(err)
	}
	lx := sc.MustGet[lexer.Lexer]()
	p := sc.MustGet[parser.Parser]()
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doc := core.NewDocumentFromString("inmemory://bench.ifc", "ifc", string(raw))
		lx.Exec(ctx, doc)
		p.Parse(ctx, doc)
		if doc.Root == nil {
			b.Fatal("no AST")
		}
	}
}
