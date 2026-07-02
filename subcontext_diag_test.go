package ifc

import (
	"context"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
)

func TestSubContext124Diagnostics(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#118= IFCDIRECTION((6.12303176911189E-17,1.));
#120= IFCGEOMETRICREPRESENTATIONCONTEXT($,'Model',3,0.01,,#118);
#124= IFCGEOMETRICREPRESENTATIONSUBCONTEXT('Axis','Model',*,*,*,*,#120,$,.GRAPH_VIEW.,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	diagnostics := serviceMustDocumentValidator(sc).Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		t.Logf("L%d [%s] %s", d.Range.Start.Line, d.Code, d.Message)
	}
}
