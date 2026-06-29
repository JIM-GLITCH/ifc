package ifc

import (
	"context"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

func TestPopulationInverseIsActingUpon(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCACTOR('actor',$,$,$,$,$);
#2=IFCRELASSIGNSTOACTOR('rel',$,$,$,(),$,#1,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	pop := BuildPopulation(context.Background(), doc.Document)
	links := pop.InverseRelations("#1", "IsActingUpon")
	if len(links) != 1 {
		t.Fatalf("expected 1 IsActingUpon inverse link, got %d", len(links))
	}
	if links[0].Name() != "#2" {
		t.Fatalf("expected inverse source #2, got %s", links[0].Name())
	}
}

func TestEntityWhereIfcArbitraryClosedProfileDefIndexedPolyCurve(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#206=IFCCARTESIANPOINTLIST2D(((0.,0.),(1.,0.),(1.,1.),(0.,0.)));
#213=IFCINDEXEDPOLYCURVE(#206,$,.F.);
#214=IFCARBITRARYCLOSEDPROFILEDEF(.AREA.,$,#213);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustDocumentValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" {
			t.Fatalf("unexpected entity WHERE diagnostic: %s", d.Message)
		}
	}
}

func TestEntityWhereIfcGeometricRepresentationSubContext(t *testing.T) {
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
#133=IFCPROJECTEDCRS('EPSG:28992','CRS',$,$,$,$,$);
#135=IFCMAPCONVERSION(#120,#133,0.,0.,0.,1.,0.,$);
#124=IFCGEOMETRICREPRESENTATIONSUBCONTEXT('Axis','Model',*,*,*,*,#120,$,.GRAPH_VIEW.,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustDocumentValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" || d.Code == "ifc-type-where" {
			t.Fatalf("unexpected diagnostic (%s): %s", d.Code, d.Message)
		}
	}
}

func TestEntityWhereIfcOwnerHistoryNoChange(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#45=IFCPERSON($,'Family',$,$,$,$,$,$);
#46=IFCORGANIZATION($,'Org',$,$,$);
#47=IFCPERSONANDORGANIZATION(#45,#46,$);
#5=IFCAPPLICATION(#46,$,$,$);
#48=IFCOWNERHISTORY(#47,#5,$,.NOCHANGE.,$,$,$,1663155515);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustDocumentValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" {
			t.Fatalf("unexpected entity WHERE diagnostic: %s", d.Message)
		}
	}
}

func TestEntityWhereIfcAxis2Placement3DLocationIs3D(t *testing.T) {
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
#32=IFCAXIS2PLACEMENT3D(#6,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustDocumentValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" {
			t.Fatalf("unexpected entity WHERE diagnostic: %s", d.Message)
		}
	}
}

func TestEntityWhereIfcDirectionMagnitudeGreaterZero(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#12=IFCDIRECTION((1.,0.,0.));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustDocumentValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" {
			t.Fatalf("unexpected entity WHERE diagnostic: %s", d.Message)
		}
	}
}

func TestEntityWhereNotContainedFeatureElement(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCSURFACEFEATURE('feat',$,$,$,$,$,$,.MARK.);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustDocumentValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" {
			t.Fatalf("unexpected entity WHERE diagnostic: %s", d.Message)
		}
	}
}

func TestEntityWhereNotContainedViolation(t *testing.T) {
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
#4=IFCBUILDING('b',$,'',$,$,#3,$,$,.ELEMENT.,$,$,$);
#5=IFCSURFACEFEATURE('feat',$,$,$,$,$,$,.MARK.);
#6=IFCRELCONTAINEDINSPATIALSTRUCTURE('rel',$,$,$,(#5),$,#4);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := serviceMustDocumentValidator(sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected ifc-entity-where when feature element is contained in spatial structure")
	}
}

func serviceMustDocumentValidator(sc *service.Container) workspace.DocumentValidator {
	return service.MustGet[workspace.DocumentValidator](sc)
}
