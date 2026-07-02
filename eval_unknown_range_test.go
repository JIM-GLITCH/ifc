package ifc

import (
	"context"
	"os"
	"testing"

	core "typefox.dev/fastbelt"
	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

func TestEvalUnknownAnchorSkipsFullDocumentNode(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	doc := fixture.Parse("ISO-10303-21;\nHEADER;\nFILE_SCHEMA(('IFC4'));\nENDSEC;\nDATA;\nENDSEC;\nEND-ISO-10303-21;")

	var diags []*core.Diagnostic
	acceptEvalUnknown(func(d *core.Diagnostic) { diags = append(diags, d) }, doc.Document.Root, nil, "Cannot evaluate parameter value")
	if len(diags) != 0 {
		t.Fatalf("expected full-document node without token to skip eval-unknown, got %v", diags)
	}
}

func TestTIfcParseErrorsDoNotPaintFullFileEvalUnknown(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	doc := fixture.Parse(string(raw))

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diags := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diags {
		if d.Code != evalUnknownCode {
			continue
		}
		span := int(d.Range.End.Line) - int(d.Range.Start.Line)
		if span > 50 {
			t.Fatalf("eval-unknown spans entire file: L%d-%d msg=%q", d.Range.Start.Line, d.Range.End.Line, d.Message)
		}
	}
}
