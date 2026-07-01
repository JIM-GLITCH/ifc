package ifc

import (
	"context"
	"testing"

	core "typefox.dev/fastbelt"
	fbtest "typefox.dev/fastbelt/test"
)

func TestShouldSkipTypeWhereOnlyUnset(t *testing.T) {
	if !shouldSkipTypeWhere(NewDEFAULT_PARAMETER()) {
		t.Fatal("expected $ to skip type WHERE")
	}
	if shouldSkipTypeWhere(NewOMITTED_PARAMETER()) {
		t.Fatal("expected * not to skip type WHERE at shouldSkipTypeWhere layer")
	}
}

func TestValidateTypeWhereOnWhereValueDimensionCount(t *testing.T) {
	attr := SchemaAttr{Kind: ParamInteger, TypeName: "IfcDimensionCount"}
	node := NewINTEGER_LITERAL()
	var diags []*core.Diagnostic
	accept := func(d *core.Diagnostic) { diags = append(diags, d) }

	validateTypeWhereOnWhereValue(attr, parseLiteralNumber("0"), node, nil, accept)
	if len(diags) != 1 || diags[0].Code != "ifc-type-where" {
		t.Fatalf("expected ifc-type-where for Dim=0, got %v", diags)
	}

	diags = nil
	validateTypeWhereOnWhereValue(attr, parseLiteralNumber("2"), node, nil, accept)
	if len(diags) != 0 {
		t.Fatalf("expected no diagnostic for Dim=2, got %v", diags)
	}
}

func TestTypeWhereSkipsUnsetOptional(t *testing.T) {
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
#3=IFCLOCALPLACEMENT($,#2);
#4=IFCSITE('id',$,'',$,$,#3,$,$,.ELEMENT.,$,$,0.,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-type-where" {
			t.Fatalf("unexpected ifc-type-where on optional $: %s", d.Message)
		}
	}
}

func TestDerivedParameterStillValidates(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#49=IFCSIUNIT(*,.LENGTHUNIT.,.MILLI.,.METRE.);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-param-required" || d.Code == "ifc-param-derived" || d.Code == "ifc-entity-where" || d.Code == "ifc-type-where" {
			t.Fatalf("unexpected diagnostic (%s): %s", d.Code, d.Message)
		}
	}
}
