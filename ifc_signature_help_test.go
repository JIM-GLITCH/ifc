package ifc

import (
	"strings"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/lsp"
)

func signatureHelpAt(t *testing.T, fixture *fbtest.Fixture, doc *fbtest.Doc, label string) *lsp.SignatureHelp {
	t.Helper()
	var offset int
	found := false
	for _, idx := range doc.Indices {
		if idx.Label == label {
			offset = idx.Offset
			found = true
		}
	}
	if !found {
		t.Fatalf("fbtest: no index marker with label %q", label)
	}
	position := doc.Document.TextDoc.PositionAt(offset)
	server, err := service.Get[lsp.Server](fixture.Services())
	if err != nil {
		t.Fatalf("fbtest: no language server available: %v", err)
	}
	result, err := server.SignatureHelp(doc.Ctx(), &lsp.SignatureHelpParams{
		TextDocumentPositionParams: lsp.TextDocumentPositionParams{
			TextDocument: lsp.TextDocumentIdentifier{URI: doc.Document.URI.DocumentURI()},
			Position:     position,
		},
	})
	if err != nil {
		t.Fatalf("fbtest: signature help request failed: %v", err)
	}
	return result
}

func expectSignatureHelp(t *testing.T, fixture *fbtest.Fixture, doc *fbtest.Doc, label, wantSignature string, wantActiveParam int) {
	t.Helper()
	result := signatureHelpAt(t, fixture, doc, label)
	if result == nil {
		t.Fatalf("fbtest: expected signature help at %q, got nil", label)
	}
	if len(result.Signatures) != 1 {
		t.Fatalf("fbtest: expected exactly one signature at %q, got %d", label, len(result.Signatures))
	}
	sig := result.Signatures[0]
	if sig.Label != wantSignature {
		t.Errorf("fbtest: expected signature label %q at %q, got %q", wantSignature, label, sig.Label)
	}
	if result.ActiveParameter == nil || int(*result.ActiveParameter) != wantActiveParam {
		got := -1
		if result.ActiveParameter != nil {
			got = int(*result.ActiveParameter)
		}
		t.Errorf("fbtest: expected active parameter %d at %q, got %d", wantActiveParam, label, got)
	}
}

func parseSignatureFixture(t *testing.T, fixture *fbtest.Fixture, body string) *fbtest.Doc {
	t.Helper()
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4X3'));
ENDSEC;
DATA;
` + body + `
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()
	return doc
}

func TestSignatureHelpFirstArgument(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := parseSignatureFixture(t, fixture, `#1=IFCAXIS2PLACEMENT3D(<|p1>$,$);`)

	expectSignatureHelp(t, fixture, doc, "p1",
		"IFCAXIS2PLACEMENT3D(Location: IfcPoint, Axis: IfcDirection, RefDirection: IfcDirection)", 0)
}

func TestSignatureHelpLaterArgument(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := parseSignatureFixture(t, fixture, `#1=IFCDIRECTION((0.,0.,1.));
#2=IFCAXIS2PLACEMENT3D(#1,$,<|p2>$);`)

	expectSignatureHelp(t, fixture, doc, "p2",
		"IFCAXIS2PLACEMENT3D(Location: IfcPoint, Axis: IfcDirection, RefDirection: IfcDirection)", 2)
}

func TestSignatureHelpInsideNestedList(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := parseSignatureFixture(t, fixture, `#1=IFCCARTESIANPOINT((0.,<|p1>0.,0.));`)

	// The cursor is inside the nested coordinate list, but signature help
	// still describes the enclosing entity's single Coordinates argument.
	expectSignatureHelp(t, fixture, doc, "p1",
		"IFCCARTESIANPOINT(Coordinates: LIST [1:3] OF IfcLengthMeasure)", 0)
}

func TestSignatureHelpDerivedAttributeDocumentsStar(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := parseSignatureFixture(t, fixture, `#1=IFCSIUNIT(<|p1>*,.LENGTHUNIT.,.MILLI.,.METRE.);`)

	result := signatureHelpAt(t, fixture, doc, "p1")
	if result == nil || len(result.Signatures) != 1 {
		t.Fatalf("fbtest: expected signature help at %q", "p1")
	}
	sig := result.Signatures[0]
	if sig.Label != "IFCSIUNIT(Dimensions: IfcDimensionalExponents, UnitType: IfcUnitEnum, Prefix: IfcSIPrefix, Name: IfcSIUnitName)" {
		t.Errorf("fbtest: unexpected signature label %q", sig.Label)
	}
	if len(sig.Parameters) != 4 {
		t.Fatalf("fbtest: expected 4 parameters, got %d", len(sig.Parameters))
	}
	if sig.Parameters[0].Label != "Dimensions: IfcDimensionalExponents" {
		t.Errorf("fbtest: expected first parameter Dimensions: IfcDimensionalExponents, got %q", sig.Parameters[0].Label)
	}
	if got := sig.Parameters[0].Documentation; got == "" || !strings.Contains(got, "*") {
		t.Errorf("fbtest: expected '*' documentation for derived attribute, got %q", got)
	}
	if got := sig.Parameters[2].Documentation; got == "" || !strings.Contains(got, "Optional") {
		t.Errorf("fbtest: expected optional documentation, got %q", got)
	}
	if sig.Documentation == nil {
		t.Fatal("fbtest: expected signature documentation with parameter types")
	}
	markup, ok := sig.Documentation.Value.(lsp.MarkupContent)
	if !ok {
		t.Fatalf("fbtest: expected markup signature documentation, got %T", sig.Documentation.Value)
	}
	if !strings.Contains(markup.Value, "- `Dimensions`: `DERIVED IfcDimensionalExponents`") {
		t.Errorf("fbtest: expected derived parameter type in signature documentation, got %q", markup.Value)
	}
	if strings.Contains(markup.Value, "ENTITY IfcSIUnit") {
		t.Errorf("fbtest: expected no EXPRESS source in signature documentation, got %q", markup.Value)
	}
}

func TestSignatureHelpOutsideArgumentList(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := parseSignatureFixture(t, fixture, `#1=IFCDIRECTION((0.,0.,1.))<|out>;`)

	result := signatureHelpAt(t, fixture, doc, "out")
	if result != nil {
		t.Fatalf("fbtest: expected no signature help outside argument list, got %+v", result)
	}
}

func TestSignatureHelpInHeaderReturnsNil(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;<|p1>1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4X3'));
ENDSEC;
DATA;
#1=IFCDIRECTION((0.,0.,1.));
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	result := signatureHelpAt(t, fixture, doc, "p1")
	if result != nil {
		t.Fatalf("fbtest: expected no signature help in header entity, got %+v", result)
	}
}

func TestSignatureHelpCapabilityAdvertised(t *testing.T) {
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
	if result == nil || result.Capabilities.SignatureHelpProvider == nil {
		t.Fatalf("fbtest: signature help capability not advertised")
	}
	triggers := result.Capabilities.SignatureHelpProvider.TriggerCharacters
	if len(triggers) != 2 || triggers[0] != "(" || triggers[1] != "," {
		t.Errorf("fbtest: unexpected trigger characters %v", triggers)
	}
}
