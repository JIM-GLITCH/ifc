package ifc

import (
	"context"
	"strings"
	"testing"

	core "typefox.dev/fastbelt"
	fbtest "typefox.dev/fastbelt/test"
)

func TestParseAggregateBounds(t *testing.T) {
	tests := []struct {
		typeName string
		ok       bool
		low      int
		high     int
	}{
		{"LIST [1:3] OF IfcLengthMeasure", true, 1, 3},
		{"SET [0:?] OF IfcPerson", true, 0, -1},
		{"ARRAY [1:2] OF REAL", true, 1, 2},
		{"LIST [2:?] OF IfcCartesianPoint", true, 2, -1},
		{"LIST [1:?] OF UNIQUE LIST [1:2] OF IfcLengthMeasure", true, 1, -1},
		{"IfcLabel", false, 0, 0},
		{"LIST [0:UpperIndex] OF IfcCartesianPoint", false, 0, 0},
	}

	for _, tc := range tests {
		bounds, ok := parseAggregateBounds(tc.typeName)
		if ok != tc.ok {
			t.Fatalf("parseAggregateBounds(%q): ok=%v, want %v", tc.typeName, ok, tc.ok)
		}
		if !ok {
			continue
		}
		if bounds.low != tc.low || bounds.high != tc.high {
			t.Fatalf("parseAggregateBounds(%q): got low=%d high=%d, want low=%d high=%d",
				tc.typeName, bounds.low, bounds.high, tc.low, tc.high)
		}
	}
}

func TestAggregateBoundsLowerBound(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCCARTESIANPOINT(());
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	if !hasDiagnosticCode(diagnostics, "ifc-aggregate-bounds") {
		t.Fatalf("expected ifc-aggregate-bounds for empty LIST [1:3], got %v", diagnostics)
	}
}

func TestAggregateBoundsUpperBound(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCCARTESIANPOINT((1.,2.,3.,4.));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	if !hasDiagnosticCode(diagnostics, "ifc-aggregate-bounds") {
		t.Fatalf("expected ifc-aggregate-bounds for 4 elements in LIST [1:3], got %v", diagnostics)
	}
}

func TestAggregateBoundsPass(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCCARTESIANPOINT((1.,2.,3.));
#2=IFCDIRECTION((1.,0.));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-aggregate-bounds" {
			t.Fatalf("unexpected aggregate bounds diagnostic: %s", d.Message)
		}
	}
}

func TestAggregateBoundsDirectionMinimum(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCDIRECTION((1.));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-aggregate-bounds" && strings.Contains(d.Message, "at least 2") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected ifc-aggregate-bounds for LIST [2:3] with 1 element, got %v", diagnostics)
	}
}

func hasDiagnosticCode(diagnostics []*core.Diagnostic, code string) bool {
	for _, d := range diagnostics {
		if d.Code == code {
			return true
		}
	}
	return false
}
