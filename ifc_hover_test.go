package ifc

import (
	"testing"

	"typefox.dev/fastbelt/server"
	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
)

func newIfcHoverFixture(t *testing.T) *fbtest.Fixture {
	t.Helper()
	sc := service.NewContainer()
	SetupServices(sc)
	server.SetupDefaultServices(sc)
	SetupIfcLanguageServer(sc)
	sc.Seal()
	return fbtest.New(t, sc)
}

func TestHoverEntityReferencePreviewsDefinition(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCDIRECTION((0.,0.,1.));
#2=IFCAXIS2PLACEMENT3D(<|ref:#1|>,$,$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectHoverAt("ref", "```ifc\n#1=IFCDIRECTION((0.,0.,1.));\n```")
}

func TestHoverEntityDeclarationHasNoPreview(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
<|decl:#1|>=IFCDIRECTION((0.,0.,1.));
#2=IFCAXIS2PLACEMENT3D(#1,$,$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectNoHoverAt("decl")
}

func TestHoverEntityNameShowsSignatureAndParameterTypes(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4X3'));
ENDSEC;
DATA;
#2=IFCCARTESIANPOINT((0.,0.,0.));
#1=<|entity:IFCPRESENTATIONLAYERASSIGNMENT|>('layer',$,(#2),$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectHoverAt("entity", "`IFCPRESENTATIONLAYERASSIGNMENT(Name, Description, AssignedItems, Identifier)`\n\n- `Name`: `IfcLabel`\n- `Description`: `OPTIONAL IfcText`\n- `AssignedItems`: `SET [1:?] OF IfcLayeredItem`\n- `Identifier`: `OPTIONAL IfcIdentifier`")
}

func TestHoverOmittedSiUnitDimensions(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCSIUNIT(<|dimensions:*|>,.LENGTHUNIT.,.MILLI.,.METRE.);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectHoverAt("dimensions", "**Dimensions** = `IfcDimensionalExponents(1,0,0,0,0,0,0)`")
}

func TestHoverInheritedSubContextDerivedValue(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCCARTESIANPOINT((0.,0.,0.));
#2=IFCAXIS2PLACEMENT3D(#1,$,$);
#3=IFCGEOMETRICREPRESENTATIONCONTEXT('Model','Model',3,1.E-5,#2,$);
#4=IFCGEOMETRICREPRESENTATIONSUBCONTEXT('Body','Model',<|coord:*|>,*,*,*,#3,0.5,.MODEL_VIEW.,$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectHoverAt("coord", "**CoordinateSpaceDimension** = `3`")
}

func TestHoverDerivedEntityReferenceShowsDefinitionAndLink(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#118=IFCDIRECTION((0.,1.));
#1=IFCCARTESIANPOINT((0.,0.,0.));
#2=IFCAXIS2PLACEMENT3D(#1,$,$);
#3=IFCGEOMETRICREPRESENTATIONCONTEXT('Model','Model',3,1.E-5,#2,#118);
#4=IFCGEOMETRICREPRESENTATIONSUBCONTEXT('Body','Model',*,*,*,<|north:*|>,#3,0.5,.MODEL_VIEW.,$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectHoverAt("north", "**TrueNorth** = [#118](command:ifc.openEntityReference?%5B%7B%22uri%22%3A%22inmemory%3A%2Ftest%22%2C%22range%22%3A%7B%22start%22%3A%7B%22line%22%3A7%2C%22character%22%3A0%7D%2C%22end%22%3A%7B%22line%22%3A7%2C%22character%22%3A4%7D%7D%7D%5D)\n\n```ifc\n#118=IFCDIRECTION((0.,1.));\n```")
}

func TestHoverEnumerationShowsTypeDeclaration(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCSIUNIT(*,<|unit:.LENGTHUNIT.|>,.MILLI.,.METRE.);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectHoverAt("unit", "**UnitType** : `IfcUnitEnum`\n\n```express\n"+enumTypeExpressDocs["IFCUNITENUM"]+"\n```")
}

func TestHoverEnumerationOnLaterAttribute(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4X3'));
ENDSEC;
DATA;
#1=IFCGEOMETRICREPRESENTATIONCONTEXT('Model','Model',3,1.E-5,$,$);
#2=IFCGEOMETRICREPRESENTATIONSUBCONTEXT('Axis','Model',*,*,*,*,#1,0.5,<|view:.MODEL_VIEW.|>,$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectHoverAt("view", "**TargetView** : `IfcGeometricProjectionEnum`\n\n```express\n"+enumTypeExpressDocs["IFCGEOMETRICPROJECTIONENUM"]+"\n```")
}

func TestHoverUnresolvedEntityReferenceHasNoPreview(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCAXIS2PLACEMENT3D(<|missing:#3|>,$,$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectNoHoverAt("missing")
}
