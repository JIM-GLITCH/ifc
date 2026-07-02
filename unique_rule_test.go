package ifc

import (
	"context"
	"strings"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

func TestUniqueRuleGlobalIdDuplicate(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROJECT('dup-id',$,'P1',$,$,$,$,$,$);
#2=IFCWALL('dup-id',$,'W1',$,$,$,$,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustUniqueRuleValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	var found int
	for _, d := range diagnostics {
		if d.Code == "ifc-unique" && strings.Contains(d.Message, "GlobalId") {
			found++
		}
	}
	if found < 2 {
		t.Fatalf("expected ifc-unique diagnostics on duplicate GlobalId, got %v", diagnostics)
	}
}

func TestUniqueRuleGlobalIdPass(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROJECT('id-1',$,'P1',$,$,$,$,$,$);
#2=IFCWALL('id-2',$,'W1',$,$,$,$,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustUniqueRuleValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-unique" {
			t.Fatalf("unexpected UNIQUE diagnostic: %s", d.Message)
		}
	}
}

func TestUniqueRuleGlobalIdCaseSensitivePass(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROPERTYSET('2uphWCk9v7hw8QHlM1lYgR',#48,'A',$,(#2));
#2=IFCPROPERTYSINGLEVALUE('x',$,$,$);
#3=IFCPROPERTYSET('2uphWCk9v7hw8QHlM1lYGr',#48,'B',$,(#2));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustUniqueRuleValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-unique" {
			t.Fatalf("case-differing GlobalIds must not collide: %s", d.Message)
		}
	}
}

func TestUniqueAggregateEnumerationValues(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROPERTYSINGLEVALUE('x',$,$,$);
#2=IFCPROPERTYENUMERATION('colors',(#1,#1),$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustUniqueRuleValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	var found bool
	for _, d := range diagnostics {
		if d.Code == "ifc-unique-aggregate" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected ifc-unique-aggregate diagnostic, got %v", diagnostics)
	}
}

func serviceMustUniqueRuleValidator(sc *service.Container) workspace.DocumentValidator {
	return service.MustGet[workspace.DocumentValidator](sc)
}
