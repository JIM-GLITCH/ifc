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

func TestHoverEntityNameShowsExpressDeclaration(t *testing.T) {
	fixture := newIfcHoverFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4X3'));
ENDSEC;
DATA;
#1=<|entity:IFCWALL|>('id',$,$,$,$,$,$,$,$);
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	doc.ExpectHoverAt("entity", "```express\n"+entityExpressDocs["IFCWALL"]+"\n```")
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
