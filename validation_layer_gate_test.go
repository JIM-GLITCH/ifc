package ifc

import (
	"context"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
)

func TestValidationLayerDisablesEntityWhere(t *testing.T) {
	sc := CreateServices()
	store, err := service.Get[*ValidationSettingsStore](sc)
	if err != nil {
		t.Fatal(err)
	}
	settings := DefaultValidationSettings()
	settings.Rules = false
	store.Set(settings)

	fixture := fbtest.New(t, sc)
	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCCARTESIANPOINT((0.));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()
	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" {
			t.Fatalf("entity WHERE should be disabled, got %q", d.Message)
		}
	}
}
