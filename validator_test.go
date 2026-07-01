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
func TestValidateEntityReferenceType(t *testing.T) {
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
#2=IFCPERSONANDORGANIZATION(#1,#1,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-entity-type" {
			found = true
		}
	}
	if !found {
		t.Fatal("expected ifc-entity-type diagnostic for wrong entity reference")
	}
}

func TestValidateDirectionAggregateType(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCDIRECTION('not-a-list');
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-param-type" {
			found = true
		}
	}
	if !found {
		t.Fatal("expected ifc-param-type diagnostic for IFCDIRECTION string parameter")
	}
}

func TestValidateInvalidEnumValue(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCOWNERHISTORY(#2,#3,$,.NOTAVALUE.,0,$,$,0);
#2=IFCPERSON($,$,$,$,$,$,$,$);
#3=IFCAPPLICATION(#4,$,$,$);
#4=IFCORGANIZATION($,$,$,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-enum-value" {
			found = true
		}
	}
	if !found {
		t.Fatal("expected ifc-enum-value diagnostic")
	}
}

func TestValidateTypedParameterKeyword(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCPROPERTYSINGLEVALUE('Name',$,IFCCARDINALPOINTREFERENCE(2),$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-typed-keyword" {
			found = true
		}
	}
	if !found {
		t.Fatal("expected ifc-typed-keyword diagnostic")
	}
}

func TestValidateMeasureWithUnitNestedSelect(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#53=IFCCONVERSIONBASEDUNIT(#54,.PLANEANGLEUNIT.,$,.RADIAN.);
#54=IFCDIMENSIONALEXPONENTS(0,0,0,0,0,0,0);
#55=IFCMEASUREWITHUNIT(IFCRATIOMEASURE(0.0174532925199433),#53);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	for _, d := range diagnostics {
		if d.Code == "ifc-typed-keyword" {
			t.Fatalf("unexpected ifc-typed-keyword: %s", d.Message)
		}
	}
}

func TestValidateListElementTypes(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCDIRECTION(('not','real'));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-param-type" {
			found = true
		}
	}
	if !found {
		t.Fatal("expected ifc-param-type diagnostic for list elements")
	}
}

func TestSIUnitDerivedDimensionsAndWR1(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#49=IFCSIUNIT(*,.LENGTHUNIT.,.MILLI.,.METRE.);`)
	rec, ok := pop.Lookup("#49")
	if !ok {
		t.Fatal("missing #49")
	}
	ec = NewEntityWhereEval(context.Background(), "IFCSIUNIT", rec.Instance, rec.Params, pop)
	dimDerived := ec.Derived("Dimensions")
	if dimDerived.kind != wvKindRecord {
		t.Fatalf("Dimensions kind=%d, want record", dimDerived.kind)
	}
	unitType := ec.QualAttr(ec.Self(), "UnitType", 0)
	dim := ec.QualAttr(ec.Self(), "Dimensions", 0)
	result := ec.CallSchema("IfcCorrectDimensions", []WhereValue{unitType, dim})
	if !wvTruthy(result) {
		t.Fatalf("IfcCorrectDimensions failed: unitType=%v dim=%v result=%v", unitType, dim, result)
	}
	okPass, detail := CheckEntityWhereRules(context.Background(), "IFCSIUNIT", rec.Instance, rec.Params, pop)
	if !okPass {
		t.Fatalf("CheckEntityWhereRules failed: %s", detail)
	}
}

func TestConversionBasedUnitExplicitDimensionsAndWR1(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#53=IFCSIUNIT(*,.PLANEANGLEUNIT.,$,.RADIAN.);
#54=IFCDIMENSIONALEXPONENTS(0,0,0,0,0,0,0);
#55=IFCMEASUREWITHUNIT(IFCRATIOMEASURE(0.0174532925199433),#53);
#56=IFCCONVERSIONBASEDUNIT(#54,.PLANEANGLEUNIT.,'DEGREE',#55);`)
	rec, ok := pop.Lookup("#56")
	if !ok {
		t.Fatal("missing #56")
	}
	ec = NewEntityWhereEval(context.Background(), "IFCCONVERSIONBASEDUNIT", rec.Instance, rec.Params, pop)
	unitType := ec.QualAttr(ec.Self(), "UnitType", 0)
	dim := ec.QualAttr(ec.Self(), "Dimensions", 0)
	if dim.entityInstance() == nil {
		t.Fatalf("Dimensions should be entity instance, got kind=%d", dim.kind)
	}
	result := ec.CallSchema("IfcCorrectDimensions", []WhereValue{unitType, dim})
	if !wvTruthy(result) {
		t.Fatalf("IfcCorrectDimensions failed: unitType=%v dim=%v result=%v", unitType, dim, result)
	}
	okPass, detail := CheckEntityWhereRules(context.Background(), "IFCCONVERSIONBASEDUNIT", rec.Instance, rec.Params, pop)
	if !okPass {
		t.Fatalf("CheckEntityWhereRules failed: %s", detail)
	}
}

func TestValidateConversionBasedUnitWR1(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#53=IFCSIUNIT(*,.PLANEANGLEUNIT.,$,.RADIAN.);
#54=IFCDIMENSIONALEXPONENTS(0,0,0,0,0,0,0);
#55=IFCMEASUREWITHUNIT(IFCRATIOMEASURE(0.0174532925199433),#53);
#56=IFCCONVERSIONBASEDUNIT(#54,.PLANEANGLEUNIT.,'DEGREE',#55);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	for _, d := range diagnostics {
		if d.Code == "ifc-entity-where" {
			t.Fatalf("unexpected ifc-entity-where: %s", d.Message)
		}
	}
}

func TestValidateSIUnitDerivedParameter(t *testing.T) {
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

	def, ok := LookupEntity("IFCSIUNIT")
	if !ok {
		t.Fatal("expected IFCSIUNIT in schema")
	}
	if len(def.Attrs) < 1 || !def.Attrs[0].Derived {
		t.Fatal("expected first IFCSIUNIT attribute to be marked derived")
	}

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	for _, d := range diagnostics {
		if d.Code == "ifc-param-required" || d.Code == "ifc-param-derived" || d.Code == "ifc-entity-where" {
			t.Fatalf("unexpected diagnostic (%s): %s", d.Code, d.Message)
		}
	}
}

func TestValidateSiteCompoundPlaneAngleMeasure(t *testing.T) {
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
#178=IFCLOCALPLACEMENT($,$);
#179=IFCSITE('2_ZvelcG11tRIbYc4wxe3f',#48,'74',$,$,#178,$,$,.ELEMENT.,(42,21,31,181945),(-71,-3,-24,-263305),0.,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	def, ok := LookupEntity("IFCSITE")
	if !ok {
		t.Fatal("expected IFCSITE in schema")
	}
	if len(def.Attrs) < 10 || def.Attrs[9].Kind != ParamAggregate {
		t.Fatalf("expected RefLatitude to be ParamAggregate, got %v", def.Attrs[9].Kind)
	}

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	for _, d := range diagnostics {
		if d.Code == "ifc-param-type" {
			t.Fatalf("unexpected diagnostic: %s", d.Message)
		}
		if d.Code == "ifc-type-where" {
			t.Fatalf("unexpected type WHERE diagnostic: %s", d.Message)
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
#206=IFCCARTESIANPOINTLIST2D(((5948.27662514854,19918.8467334511),(5510.01489411603,20894.9751343863),(4433.53684448654,20411.6584590421),(4871.79857551905,19435.5300581068),(5948.27662514854,19918.8467334511)));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	for _, d := range diagnostics {
		if d.Code == "ifc-param-type" || d.Code == "ifc-param-count" || d.Code == "ifc-unknown-entity" {
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

func TestValidateTypeWherePositiveInteger(t *testing.T) {
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
#4=IFCBUILDING('id',$,'',$,$,#3,$,$,.ELEMENT.,$,$,$);
#5=IFCBUILDINGSTOREY('id',$,'',$,$,#3,$,$,.ELEMENT.,0);
#6=IFCRELAGGREGATES('id',$,$,$,#4,(#5));
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	for _, d := range diagnostics {
		if d.Code == "ifc-type-where" {
			t.Fatalf("unexpected type WHERE diagnostic: %s", d.Message)
		}
	}
}

func TestValidateTypeWhereViolations(t *testing.T) {
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
#4=IFCSITE('id',$,'',$,$,#3,$,$,.ELEMENT.,(1,90,0,0),(0,0,0,0),0.,$,$);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diagnostics := validator.Validate(context.Background(), doc.Document, "test")

	found := false
	for _, d := range diagnostics {
		if d.Code == "ifc-type-where" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected ifc-type-where diagnostic for invalid compound plane angle")
	}
}
