package ifc

import (
	"context"
	"os"
	"strings"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

func TestTIfcBuilding150UniquePropertySetNames(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	doc := fixture.Parse(string(raw))
	doc.AssertNoParseErrors()
	pop := BuildPopulation(context.Background(), doc.Document)
	rec, ok := pop.Lookup("#150")
	if !ok {
		t.Fatal("#150 not found")
	}
	isDefinedBy := pop.InverseRelations("#150", "IsDefinedBy")
	t.Logf("IsDefinedBy count: %d", len(isDefinedBy))
	for _, rel := range isDefinedBy {
		t.Logf("  relation: %s", rel.Name())
	}
	okPass, detail, evalUnknowns := CheckEntityWhereRules(context.Background(), "IFCBUILDING", rec.Instance, rec.Params, pop)
	if len(evalUnknowns) > 0 {
		for _, msg := range evalUnknowns {
			t.Logf("eval-unknown: %s", msg)
		}
	}
	if !okPass {
		t.Fatalf("CheckEntityWhereRules failed: %q evalUnknowns=%v", detail, evalUnknowns)
	}
}

func TestTIfcPropertySetGlobalIdCaseDistinct(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	doc := fixture.Parse(string(raw))
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-unique" && strings.Contains(d.Message, "#64609") && strings.Contains(d.Message, "#65761") {
			t.Fatalf("false duplicate GlobalId between #64609 and #65761: %s", d.Message)
		}
	}
}
