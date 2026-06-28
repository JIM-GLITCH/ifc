package ifc

import (
	"context"
	"strings"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

func TestSchemaLookup(t *testing.T) {
	def, ok := LookupEntity("IFCPERSON")
	if !ok {
		t.Fatal("expected IFCPERSON in generated schema")
	}
	if def.ParamCount != 8 {
		t.Fatalf("expected 8 parameters for IFCPERSON, got %d", def.ParamCount)
	}

	dir, ok := LookupEntity("IFCDIRECTION")
	if !ok {
		t.Fatal("expected IFCDIRECTION in generated schema")
	}
	if dir.ParamCount != 1 {
		t.Fatalf("expected 1 parameter for IFCDIRECTION (derived attrs omitted in STEP), got %d", dir.ParamCount)
	}
}

func TestValidatePersonInstance(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPERSON('a','b','c',$,$,$,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	for _, d := range diagnostics {
		if d.Code == "ifc-param-count" || d.Code == "ifc-unknown-entity" {
			t.Fatalf("unexpected diagnostic: %s (%s)", d.Message, d.Code)
		}
	}
}

func TestValidateCartesianPointList2DTrailingOptional(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCCARTESIANPOINTLIST2D(((1.,2.),(3.,4.)));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	for _, d := range diagnostics {
		if d.Code == "ifc-param-count" || d.Code == "ifc-unknown-entity" {
			t.Fatalf("unexpected diagnostic: %s (%s)", d.Message, d.Code)
		}
	}
}

func TestValidateWrongParameterCount(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCDIRECTION();
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-param-count" {
			found = true
			if !strings.Contains(d.Message, "missing required parameter") {
				t.Fatalf("unexpected message: %s", d.Message)
			}
		}
	}
	if !found {
		t.Fatal("expected ifc-param-count diagnostic")
	}
}
