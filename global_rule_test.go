package ifc

import (
	"context"
	"strings"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

func TestGlobalRuleSingleProjectInstance(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROJECT('p1',$,'P1',$,$,$,$,$,$);
#2=IFCPROJECT('p2',$,'P2',$,$,$,$,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustGlobalRuleValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	var found bool
	for _, d := range diagnostics {
		if d.Code == "ifc-global-rule" && strings.Contains(d.Message, "SIZEOF(IfcProject)") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected ifc-global-rule diagnostic for IfcSingleProjectInstance, got %v", diagnostics)
	}
}

func TestGlobalRuleSingleProjectInstancePass(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROJECT('p1',$,'P1',$,$,$,$,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustGlobalRuleValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-global-rule" {
			t.Fatalf("unexpected global RULE diagnostic: %s", d.Message)
		}
	}
}

func TestGlobalRuleRepresentationContextSameWCS(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#6=IFCCARTESIANPOINT((0.,0.,0.));
#117=IFCAXIS2PLACEMENT3D(#6,$,$);
#118=IFCDIRECTION((0.,1.));
#120=IFCGEOMETRICREPRESENTATIONCONTEXT($,'Model',3,0.01,#117,#118);
#121=IFCGEOMETRICREPRESENTATIONCONTEXT($,'Plan',3,0.01,#117,#118);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustGlobalRuleValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-global-rule" {
			t.Fatalf("unexpected global RULE diagnostic: %s", d.Message)
		}
	}
}

func TestGlobalRuleRepresentationContextDifferentWCS(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#6=IFCCARTESIANPOINT((0.,0.,0.));
#7=IFCCARTESIANPOINT((1.,0.,0.));
#117=IFCAXIS2PLACEMENT3D(#6,$,$);
#127=IFCAXIS2PLACEMENT3D(#7,$,$);
#118=IFCDIRECTION((0.,1.));
#120=IFCGEOMETRICREPRESENTATIONCONTEXT($,'Model',3,0.01,#117,#118);
#121=IFCGEOMETRICREPRESENTATIONCONTEXT($,'Plan',3,0.01,#127,#118);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustGlobalRuleValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	var found bool
	for _, d := range diagnostics {
		if d.Code == "ifc-global-rule" && strings.Contains(d.Message, "IsDifferent = FALSE") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected ifc-global-rule diagnostic for IfcRepresentationContextSameWCS, got %v", diagnostics)
	}
}

func serviceMustGlobalRuleValidator(sc *service.Container) workspace.DocumentValidator {
	return service.MustGet[workspace.DocumentValidator](sc)
}
