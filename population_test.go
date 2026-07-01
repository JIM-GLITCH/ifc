package ifc

import (
	"context"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

func TestPopulationOfProductRepresentationInverseViaSubtype(t *testing.T) {
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
#126=IFCGEOMETRICREPRESENTATIONSUBCONTEXT('Body','Model',*,*,*,*,#120,$,.MODEL_VIEW.,$);
#20=IFCDIRECTION((0.,0.,1.));
#1=IFCCARTESIANPOINT((0.,0.));
#2=IFCCARTESIANPOINT((1.,0.));
#3=IFCCARTESIANPOINT((1.,1.));
#213=IFCPOLYLINE((#1,#2,#3,#1));
#214=IFCARBITRARYCLOSEDPROFILEDEF(.AREA.,$,#213);
#217=IFCAXIS2PLACEMENT3D(#6,$,$);
#218=IFCEXTRUDEDAREASOLID(#214,#217,#20,1.);
#219=IFCSHAPEREPRESENTATION(#126,'Body','SweptSolid',(#218));
#226=IFCPRODUCTDEFINITIONSHAPE($,$,(#219));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	pop := BuildPopulation(context.Background(), doc.Document)
	links := pop.InverseRelations("#219", "OfProductRepresentation")
	if len(links) != 1 {
		t.Fatalf("expected 1 OfProductRepresentation inverse link, got %d", len(links))
	}
	if links[0].Name() != "#226" {
		t.Fatalf("expected inverse source #226, got %s", links[0].Name())
	}

	rec, ok := pop.Lookup("#219")
	if !ok {
		t.Fatal("missing #219")
	}
	okPass, detail, _ := CheckEntityWhereRules(context.Background(), "IFCSHAPEREPRESENTATION", rec.Instance, rec.Params, pop)
	if !okPass {
		t.Fatalf("CheckEntityWhereRules IFCSHAPEREPRESENTATION WR11 failed: %s", detail)
	}
}

func TestPopulationHasAssociationsDeduped(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#48=IFCOWNERHISTORY($,$,$,$,$);
#3228=IFCMATERIAL('Mat',$);
#2912=IFCWALL('wall',#48,'W',$,$,$,$,$,.NOTDEFINED.);
#94687=IFCRELASSOCIATESMATERIAL('r2',#48,$,$,(#2912),#3228);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()
	pop := BuildPopulation(context.Background(), doc.Document)

	assoc := pop.InverseRelations("#2912", "HasAssociations")
	if len(assoc) != 1 {
		t.Fatalf("expected 1 HasAssociations inverse link, got %d", len(assoc))
	}
	if assoc[0].Name() != "#94687" {
		t.Fatalf("expected inverse source #94687, got %s", assoc[0].Name())
	}

	rec, ok := pop.Lookup("#2912")
	if !ok {
		t.Fatal("missing #2912")
	}
	okPass, detail, _ := CheckEntityWhereRules(context.Background(), "IFCWALL", rec.Instance, rec.Params, pop)
	if !okPass {
		t.Fatalf("CheckEntityWhereRules IFCWALL MaxOneMaterialAssociation failed: %s", detail)
	}
}

func TestCartesianTransformationOperator3DOmittedAxis3(t *testing.T) {
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
#62595=IFCCARTESIANTRANSFORMATIONOPERATOR3D($,$,#6,1.,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()
	pop := BuildPopulation(context.Background(), doc.Document)

	rec, ok := pop.Lookup("#62595")
	if !ok {
		t.Fatal("missing #62595")
	}
	okPass, detail, _ := CheckEntityWhereRules(context.Background(), "IFCCARTESIANTRANSFORMATIONOPERATOR3D", rec.Instance, rec.Params, pop)
	if !okPass {
		t.Fatalf("CheckEntityWhereRules failed: %s", detail)
	}
}

func TestRelAssociatesMaterialAllowedElementsWindow(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#48=IFCOWNERHISTORY($,$,$,$,$);
#94038=IFCMATERIAL('Glass',$);
#94039=IFCMATERIAL('Frame',$);
#94040=IFCMATERIALCONSTITUENTSET('MaterialConstituentSet',$,(#94038,#94039));
#94035=IFCWINDOW('win',#48,'WI',$,'desc',$,$,'2240803',733.,2233.,.WINDOW.,.NOTDEFINED.,$);
#95125=IFCRELASSOCIATESMATERIAL('rel',#48,$,$,(#94035),#94040);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()
	pop := BuildPopulation(context.Background(), doc.Document)

	rec, ok := pop.Lookup("#95125")
	if !ok {
		t.Fatal("missing #95125")
	}
	okPass, detail, _ := CheckEntityWhereRules(context.Background(), "IFCRELASSOCIATESMATERIAL", rec.Instance, rec.Params, pop)
	if !okPass {
		t.Fatalf("CheckEntityWhereRules IFCRELASSOCIATESMATERIAL AllowedElements failed: %s", detail)
	}
}

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
