package ifc

import (
	"context"
	"testing"

	"typefox.dev/fastbelt/server"
	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/lsp"
)

func newIfcSymbolFixture(t *testing.T) (*fbtest.Fixture, *service.Container) {
	t.Helper()
	sc := service.NewContainer()
	SetupServices(sc)
	server.SetupDefaultServices(sc)
	SetupIfcLanguageServer(sc)
	sc.Seal()
	return fbtest.New(t, sc), sc
}

func documentSymbols(t *testing.T, sc *service.Container, doc *fbtest.Doc) []lsp.DocumentSymbol {
	t.Helper()
	provider, err := service.Get[server.DocumentSymbolProvider](sc)
	if err != nil {
		t.Fatalf("document symbol provider: %v", err)
	}
	symbols, err := provider.HandleDocumentSymbolRequest(context.Background(), &lsp.DocumentSymbolParams{
		TextDocument: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI(doc.Document.URI.DocumentURI())},
	})
	if err != nil {
		t.Fatalf("document symbols: %v", err)
	}
	return symbols
}

func findChildSymbol(symbols []lsp.DocumentSymbol, name string) *lsp.DocumentSymbol {
	for i := range symbols {
		if symbols[i].Detail == name {
			return &symbols[i]
		}
	}
	return nil
}

func TestDocumentSymbolsSpatialHierarchy(t *testing.T) {
	fixture, sc := newIfcSymbolFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROJECT('g1',$,'Project',$,$,$,$,$,$);
#2=IFCSITE('g2',$,'Site',$,$,$,$,$,$,$,$,$,$);
#3=IFCBUILDING('g3',$,'Building',$,$,$,$,$,$,$,$,$);
#4=IFCBUILDINGSTOREY('g4',$,'Level 1',$,$,$,$,'Level 1',$,$);
#5=IFCSPACE('g5',$,'Space',$,$,$,$,'Room 101',$,.SPACE.,$);
#6=IFCWALL('g6',$,'Wall A',$,$,$,$,$,$);
#7=IFCOPENINGELEMENT('g7',$,'Opening',$,$,$,$,$,$);
#10=IFCRELAGGREGATES('r1',$,$,$,#1,(#2));
#11=IFCRELAGGREGATES('r2',$,$,$,#2,(#3));
#12=IFCRELAGGREGATES('r3',$,$,$,#3,(#4));
#13=IFCRELAGGREGATES('r4',$,$,$,#4,(#5));
#14=IFCRELCONTAINEDINSPATIALSTRUCTURE('r5',$,$,$,(#6),#5);
#15=IFCRELVOIDSELEMENT('r6',$,$,$,#6,#7);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	symbols := documentSymbols(t, sc, doc)
	if len(symbols) != 1 {
		t.Fatalf("expected 1 root symbol, got %d: %#v", len(symbols), symbols)
	}
	project := symbols[0]
	if project.Name != "IFCPROJECT #1" || project.Detail != "Project" || project.Kind != lsp.Namespace {
		t.Fatalf("unexpected project symbol: %#v", project)
	}

	site := findChildSymbol(project.Children, "Site")
	if site == nil {
		t.Fatalf("missing Site child: %#v", project.Children)
	}
	building := findChildSymbol(site.Children, "Building")
	if building == nil {
		t.Fatalf("missing Building child: %#v", site.Children)
	}
	storey := findChildSymbol(building.Children, "Level 1")
	if storey == nil {
		t.Fatalf("missing storey child: %#v", building.Children)
	}
	space := findChildSymbol(storey.Children, "Room 101")
	if space == nil {
		t.Fatalf("missing space child: %#v", storey.Children)
	}
	wall := findChildSymbol(space.Children, "Wall A")
	if wall == nil {
		t.Fatalf("missing wall child: %#v", space.Children)
	}
	if wall.Kind != lsp.Object {
		t.Fatalf("unexpected wall kind: %v", wall.Kind)
	}
	opening := findChildSymbol(wall.Children, "Opening")
	if opening == nil {
		t.Fatalf("missing opening child: %#v", wall.Children)
	}
}

func TestDocumentSymbolsOmitsNonSpatialEntities(t *testing.T) {
	fixture, sc := newIfcSymbolFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROJECT('g1',$,'Project',$,$,$,$,$,$);
#2=IFCWALLTYPE('g2',$,'WallType',$,$,$,$,'WallType',$,.STANDARD.);
#3=IFCPROPERTYSET('g3',$,'Pset',$,(#4));
#4=IFCPROPERTYSINGLEVALUE('Name',$,IFCLABEL('x'),$);
#5=IFCZONE('g5',$,'Zone',$,$,$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	symbols := documentSymbols(t, sc, doc)
	if len(symbols) != 1 || symbols[0].Detail != "Project" {
		t.Fatalf("expected only Project symbol, got %#v", symbols)
	}
	if len(symbols[0].Children) != 0 {
		t.Fatalf("expected no children, got %#v", symbols[0].Children)
	}
}

// Parents may be defined after their children in the file; attaching a node
// upward must not discard children that were attached earlier.
func TestDocumentSymbolsParentsDefinedAfterChildren(t *testing.T) {
	fixture, sc := newIfcSymbolFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#10=IFCWALL('g10',$,'Wall A',$,$,$,$,$,$);
#11=IFCBUILDINGSTOREY('g11',$,'Level 1',$,$,$,$,'Level 1',$,$);
#12=IFCBUILDING('g12',$,'Building',$,$,$,$,$,$,$,$,$);
#13=IFCSITE('g13',$,'Site',$,$,$,$,$,$,$,$,$,$);
#1=IFCPROJECT('g1',$,'Project',$,$,$,$,$,$);
#19=IFCRELAGGREGATES('r0',$,$,$,#1,(#13));
#20=IFCRELAGGREGATES('r1',$,$,$,#13,(#12));
#21=IFCRELAGGREGATES('r2',$,$,$,#12,(#11));
#22=IFCRELCONTAINEDINSPATIALSTRUCTURE('r3',$,$,$,(#10),#11);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	symbols := documentSymbols(t, sc, doc)
	if len(symbols) != 1 {
		t.Fatalf("expected 1 root symbol, got %d: %#v", len(symbols), symbols)
	}
	project := symbols[0]
	site := findChildSymbol(project.Children, "Site")
	if site == nil {
		t.Fatalf("missing Site child: %#v", project.Children)
	}
	building := findChildSymbol(site.Children, "Building")
	if building == nil {
		t.Fatalf("missing Building child: %#v", site.Children)
	}
	storey := findChildSymbol(building.Children, "Level 1")
	if storey == nil {
		t.Fatalf("missing storey child: %#v", building.Children)
	}
	wall := findChildSymbol(storey.Children, "Wall A")
	if wall == nil {
		t.Fatalf("missing wall child: %#v", storey.Children)
	}
}
