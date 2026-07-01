package ifc

import (
	"context"
	"strings"
	"testing"

	core "typefox.dev/fastbelt"
	fbtest "typefox.dev/fastbelt/test"
)

func TestValidateTypeWhereOnUnknownReportsEvalUnknown(t *testing.T) {
	attr := SchemaAttr{Kind: ParamInteger, TypeName: "IfcDimensionCount"}
	node := NewINTEGER_LITERAL()
	var diags []*core.Diagnostic
	accept := func(d *core.Diagnostic) { diags = append(diags, d) }

	validateTypeWhereOnWhereValue(attr, wvUnknownVal(), node, nil, accept)
	if len(diags) != 1 || diags[0].Code != evalUnknownCode {
		t.Fatalf("expected %s for unknown value, got %v", evalUnknownCode, diags)
	}

	diags = nil
	validateTypeWhereOnWhereValue(attr, wvAbsentVal(), node, nil, accept)
	if len(diags) != 0 {
		t.Fatalf("expected absent to skip type WHERE, got %v", diags)
	}
}

func TestUnresolvedEntityReferenceReportsEvalUnknown(t *testing.T) {
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
#2=IFCAXIS2PLACEMENT3D(#1,$,$);
#3=IFCLOCALPLACEMENT($,#999);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	found := false
	for _, d := range diagnostics {
		if d.Code == evalUnknownCode && strings.Contains(d.Message, "#999") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected %s for unresolved #999, got %v", evalUnknownCode, diagnostics)
	}
}
