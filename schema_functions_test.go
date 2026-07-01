package ifc

import (
	"context"
	"math"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
)

const schemaTestHeader = `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
`

func schemaTestEnv(t *testing.T, data string) (*Population, *EntityWhereEval) {
	t.Helper()
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	source := schemaTestHeader + "DATA;\n" + data + "\nENDSEC;\nEND-ISO-10303-21;"
	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()
	pop := BuildPopulation(context.Background(), doc.Document)
	ctx := WithPopulation(context.Background(), pop)
	ec := NewEntityWhereEval(ctx, "IFCDIRECTION", nil, nil, pop)
	return pop, ec
}

func entityWV(t *testing.T, pop *Population, name string) WhereValue {
	t.Helper()
	rec, ok := pop.Lookup(name)
	if !ok {
		t.Fatalf("entity %s not found", name)
	}
	return wvFromEntityInstance(rec.Instance)
}

func callSchema(t *testing.T, ec *EntityWhereEval, name string, args ...WhereValue) WhereValue {
	t.Helper()
	result := ec.CallSchema(name, args)
	if result.kind == wvKindUnknown {
		t.Fatalf("CallSchema(%s) returned unknown", name)
	}
	return result
}

func wvMustReal(t *testing.T, v WhereValue) float64 {
	t.Helper()
	n, ok := wvAsNumber(v)
	if !ok {
		t.Fatalf("expected numeric WhereValue, got kind %v", v.kind)
	}
	return n
}

func assertNear(t *testing.T, label string, got, want, eps float64) {
	t.Helper()
	if math.Abs(got-want) > eps {
		t.Fatalf("%s: got %g, want %g (eps %g)", label, got, want, eps)
	}
}

func directionRatios(t *testing.T, ec *EntityWhereEval, v WhereValue) []float64 {
	t.Helper()
	ratios := ec.QualAttr(v, "DirectionRatios", 0)
	if ratios.kind != wvKindList {
		t.Fatalf("expected DirectionRatios list, got kind %v", ratios.kind)
	}
	out := make([]float64, len(ratios.list))
	for i, item := range ratios.list {
		out[i] = wvMustReal(t, item)
	}
	return out
}

func TestSchemaIfcNormaliseDirection(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((3.,4.,0.));`)
	dir := entityWV(t, pop, "#1")
	result := callSchema(t, ec, "IfcNormalise", dir)
	ratios := directionRatios(t, ec, result)
	if len(ratios) != 3 {
		t.Fatalf("expected 3 ratios, got %d", len(ratios))
	}
	assertNear(t, "ratio[0]", ratios[0], 0.6, 1e-9)
	assertNear(t, "ratio[1]", ratios[1], 0.8, 1e-9)
	assertNear(t, "ratio[2]", ratios[2], 0.0, 1e-9)
}

func TestSchemaIfcNormaliseVector(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((3.,4.,0.));
#2=IFCVECTOR(#1,5.);`)
	vec := entityWV(t, pop, "#2")
	result := callSchema(t, ec, "IfcNormalise", vec)
	mag := wvMustReal(t, ec.QualAttr(result, "Magnitude", 0))
	assertNear(t, "Magnitude", mag, 1.0, 1e-9)
	ratios := directionRatios(t, ec, ec.QualAttr(result, "Orientation", 0))
	assertNear(t, "ratio[0]", ratios[0], 0.6, 1e-9)
	assertNear(t, "ratio[1]", ratios[1], 0.8, 1e-9)
	assertNear(t, "ratio[2]", ratios[2], 0.0, 1e-9)
}

func TestSchemaIfcDotProduct(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((1.,0.,0.));
#2=IFCDIRECTION((0.,1.,0.));
#3=IFCDIRECTION((2.,0.,0.));`)
	x := entityWV(t, pop, "#1")
	y := entityWV(t, pop, "#2")
	x2 := entityWV(t, pop, "#3")

	assertNear(t, "orthogonal", wvMustReal(t, callSchema(t, ec, "IfcDotProduct", x, y)), 0.0, 1e-9)
	assertNear(t, "parallel", wvMustReal(t, callSchema(t, ec, "IfcDotProduct", x, x)), 1.0, 1e-9)
	assertNear(t, "scaled parallel", wvMustReal(t, callSchema(t, ec, "IfcDotProduct", x, x2)), 1.0, 1e-9)
}

func TestSchemaIfcDotProductConstructedDirection(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#20=IFCDIRECTION((0.,0.,1.));`)
	zAxis := wvEntityIntersect(
		wvEntityIntersect(wvEntityTypeTag("IfcRepresentationItem"), wvEntityTypeTag("IfcGeometricRepresentationItem")),
		wvNewDirection(wvFromList([]WhereValue{parseLiteralNumber("0.0"), parseLiteralNumber("0.0"), parseLiteralNumber("1.0")})),
	)
	extruded := entityWV(t, pop, "#20")
	assertNear(t, "z dot extruded", wvMustReal(t, callSchema(t, ec, "IfcDotProduct", zAxis, extruded)), 1.0, 1e-9)
}

func TestExtrudedAreaSolidValidExtrusionDirection(t *testing.T) {
	pop, _ := schemaTestEnv(t, `#20=IFCDIRECTION((0.,0.,1.));
#214=IFCARBITRARYCLOSEDPROFILEDEF(.AREA.,$,#213);
#213=IFCPOLYLINE((#1,#2,#3,#4,#5));
#1=IFCCARTESIANPOINT((0.,0.));
#2=IFCCARTESIANPOINT((1.,0.));
#3=IFCCARTESIANPOINT((1.,1.));
#4=IFCCARTESIANPOINT((0.,1.));
#5=IFCCARTESIANPOINT((0.,0.));
#217=IFCAXIS2PLACEMENT3D(#6,$,$);
#6=IFCCARTESIANPOINT((0.,0.,0.));
#218=IFCEXTRUDEDAREASOLID(#214,#217,#20,3400.);`)
	rec, ok := pop.Lookup("#218")
	if !ok {
		t.Fatal("missing #218")
	}
	okPass, detail := CheckEntityWhereRules(context.Background(), "IFCEXTRUDEDAREASOLID", rec.Instance, rec.Params, pop)
	if !okPass {
		t.Fatalf("CheckEntityWhereRules failed: %s", detail)
	}
}

func TestSchemaIfcCrossProduct(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((1.,0.,0.));
#2=IFCDIRECTION((0.,1.,0.));`)
	x := entityWV(t, pop, "#1")
	y := entityWV(t, pop, "#2")
	result := callSchema(t, ec, "IfcCrossProduct", x, y)
	mag := wvMustReal(t, ec.QualAttr(result, "Magnitude", 0))
	assertNear(t, "Magnitude", mag, 1.0, 1e-9)
	ratios := directionRatios(t, ec, ec.QualAttr(result, "Orientation", 0))
	if len(ratios) != 3 {
		t.Fatalf("expected 3 ratios, got %d", len(ratios))
	}
	assertNear(t, "ratio[0]", ratios[0], 0.0, 1e-9)
	assertNear(t, "ratio[1]", ratios[1], 0.0, 1e-9)
	assertNear(t, "ratio[2]", ratios[2], 1.0, 1e-9)
}

func TestSchemaIfcConvertDirectionInto2D(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((1.,2.,3.));`)
	dir := entityWV(t, pop, "#1")
	result := callSchema(t, ec, "IfcConvertDirectionInto2D", dir)
	ratios := directionRatios(t, ec, result)
	if len(ratios) != 2 {
		t.Fatalf("expected 2 ratios, got %d", len(ratios))
	}
	assertNear(t, "ratio[0]", ratios[0], 1.0, 1e-9)
	assertNear(t, "ratio[1]", ratios[1], 2.0, 1e-9)
}

func TestSchemaIfcOrthogonalComplement(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((3.,4.));`)
	dir := entityWV(t, pop, "#1")
	result := callSchema(t, ec, "IfcOrthogonalComplement", dir)
	ratios := directionRatios(t, ec, result)
	if len(ratios) != 2 {
		t.Fatalf("expected 2 ratios, got %d", len(ratios))
	}
	assertNear(t, "ratio[0]", ratios[0], -4.0, 1e-9)
	assertNear(t, "ratio[1]", ratios[1], 3.0, 1e-9)
}

func TestSchemaIfcCrossProductReverse(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((0.,1.,0.));
#2=IFCDIRECTION((1.,0.,0.));`)
	y := entityWV(t, pop, "#1")
	x := entityWV(t, pop, "#2")
	result := callSchema(t, ec, "IfcCrossProduct", y, x)
	ratios := directionRatios(t, ec, ec.QualAttr(result, "Orientation", 0))
	assertNear(t, "ratio[0]", ratios[0], 0.0, 1e-9)
	assertNear(t, "ratio[1]", ratios[1], 0.0, 1e-9)
	assertNear(t, "ratio[2]", ratios[2], -1.0, 1e-9)
}

func TestSchemaIfcNormaliseAlreadyUnit(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((1.,0.,0.));`)
	dir := entityWV(t, pop, "#1")
	result := callSchema(t, ec, "IfcNormalise", dir)
	ratios := directionRatios(t, ec, result)
	assertNear(t, "ratio[0]", ratios[0], 1.0, 1e-9)
	assertNear(t, "ratio[1]", ratios[1], 0.0, 1e-9)
	assertNear(t, "ratio[2]", ratios[2], 0.0, 1e-9)
}

func TestSchemaIfcScalarTimesVectorOnDirection(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCDIRECTION((3.,4.,0.));`)
	dir := entityWV(t, pop, "#1")
	result := callSchema(t, ec, "IfcScalarTimesVector", parseLiteralNumber("2.0"), dir)
	mag := wvMustReal(t, ec.QualAttr(result, "Magnitude", 0))
	if math.Abs(mag-2.0) > 1e-9 {
		t.Fatalf("expected magnitude 2, got %g", mag)
	}
	ratios := directionRatios(t, ec, ec.QualAttr(result, "Orientation", 0))
	assertNear(t, "ratio[0]", ratios[0], 0.6, 1e-9)
	assertNear(t, "ratio[1]", ratios[1], 0.8, 1e-9)
}

func TestSchemaIfcUniquePropertySetNamesDetectsDuplicate(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCPROPERTYSET('g1',$,'A',$,(#2));
#2=IFCPROPERTYSINGLEVALUE('x',$,$,$);
#3=IFCPROPERTYSET('g2',$,'A',$,(#4));
#4=IFCPROPERTYSINGLEVALUE('y',$,$,$);`)
	set := wvFromList([]WhereValue{
		entityWV(t, pop, "#1"),
		entityWV(t, pop, "#3"),
	})
	result := callSchema(t, ec, "IfcUniquePropertySetNames", set)
	if wvTruthy(result) {
		t.Fatal("expected duplicate property set names to fail")
	}
}

func TestSchemaIfcUniquePropertySetNamesPass(t *testing.T) {
	pop, ec := schemaTestEnv(t, `#1=IFCPROPERTYSET('g1',$,'A',$,(#2));
#2=IFCPROPERTYSINGLEVALUE('x',$,$,$);
#3=IFCPROPERTYSET('g2',$,'B',$,(#4));
#4=IFCPROPERTYSINGLEVALUE('y',$,$,$);`)
	set := wvFromList([]WhereValue{
		entityWV(t, pop, "#1"),
		entityWV(t, pop, "#3"),
	})
	result := callSchema(t, ec, "IfcUniquePropertySetNames", set)
	if !wvTruthy(result) {
		t.Fatal("expected unique property set names to pass")
	}
}
