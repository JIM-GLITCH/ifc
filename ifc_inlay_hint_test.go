package ifc

import (
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/lsp"
)

func inlayHintsForDoc(t *testing.T, fixture *fbtest.Fixture, doc *fbtest.Doc) []lsp.InlayHint {
	t.Helper()
	server, err := service.Get[lsp.Server](fixture.Services())
	if err != nil {
		t.Fatalf("fbtest: no language server available: %v", err)
	}
	content := doc.Document.TextDoc.Content()
	result, err := server.InlayHint(doc.Ctx(), &lsp.InlayHintParams{
		TextDocument: lsp.TextDocumentIdentifier{URI: doc.Document.URI.DocumentURI()},
		Range: lsp.Range{
			Start: doc.Document.TextDoc.PositionAt(0),
			End:   doc.Document.TextDoc.PositionAt(len(content)),
		},
	})
	if err != nil {
		t.Fatalf("fbtest: inlay hint request failed: %v", err)
	}
	return result
}

func inlayHintLabels(hints []lsp.InlayHint) []string {
	labels := make([]string, 0, len(hints))
	for _, hint := range hints {
		if len(hint.Label) == 1 {
			labels = append(labels, hint.Label[0].Value)
		}
	}
	return labels
}

func TestInlayHintEntityArguments(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := parseSignatureFixture(t, fixture, `#1=IFCDIRECTION((0.,0.,1.));
#2=IFCAXIS2PLACEMENT3D(#1,$,$);`)

	hints := inlayHintsForDoc(t, fixture, doc)
	labels := inlayHintLabels(hints)

	want := []string{
		"DirectionRatios: ",
		"Location: ", "Axis: ", "RefDirection: ",
	}
	if len(labels) != len(want) {
		t.Fatalf("fbtest: expected hints %v, got %v", want, labels)
	}
	for i := range want {
		if labels[i] != want[i] {
			t.Errorf("fbtest: hint %d: expected %q, got %q", i, want[i], labels[i])
		}
		if hints[i].Kind != lsp.Parameter {
			t.Errorf("fbtest: hint %d: expected Parameter kind, got %v", i, hints[i].Kind)
		}
	}
}

func TestInlayHintSkipsCommasInsideLists(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := parseSignatureFixture(t, fixture, `#1=IFCCARTESIANPOINT((1.,2.,3.));
#2=IFCAXIS2PLACEMENT3D(#1,$,$);`)

	hints := inlayHintsForDoc(t, fixture, doc)
	labels := inlayHintLabels(hints)

	// The coordinate list contains commas but must produce a single
	// Coordinates hint for IFCCARTESIANPOINT.
	want := []string{
		"Coordinates: ",
		"Location: ", "Axis: ", "RefDirection: ",
	}
	if len(labels) != len(want) {
		t.Fatalf("fbtest: expected hints %v, got %v", want, labels)
	}
	for i := range want {
		if labels[i] != want[i] {
			t.Errorf("fbtest: hint %d: expected %q, got %q", i, want[i], labels[i])
		}
	}
}

func TestInlayHintHeaderEntitiesIgnored(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4X3'));
ENDSEC;
DATA;
#1=IFCDIRECTION((0.,0.,1.));
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	hints := inlayHintsForDoc(t, fixture, doc)
	labels := inlayHintLabels(hints)

	want := []string{"DirectionRatios: "}
	if len(labels) != len(want) || labels[0] != want[0] {
		t.Fatalf("fbtest: expected only DATA entity hints %v, got %v", want, labels)
	}
}

func TestInlayHintCapabilityAdvertised(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := parseSignatureFixture(t, fixture, `#1=IFCDIRECTION((0.,0.,1.));`)

	server, err := service.Get[lsp.Server](fixture.Services())
	if err != nil {
		t.Fatalf("fbtest: no language server available: %v", err)
	}
	result, err := server.Initialize(doc.Ctx(), &lsp.ParamInitialize{})
	if err != nil {
		t.Fatalf("fbtest: initialize failed: %v", err)
	}
	if result == nil || result.Capabilities.InlayHintProvider == nil {
		t.Fatalf("fbtest: inlay hint capability not advertised")
	}
}
