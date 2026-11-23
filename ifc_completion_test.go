package ifc

import (
	"testing"

	"typefox.dev/fastbelt/server"
	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/lsp"
)

func newIfcCompletionFixture(t *testing.T) *fbtest.Fixture {
	t.Helper()
	sc := service.NewContainer()
	SetupServices(sc)
	server.SetupDefaultServices(sc)
	SetupIfcLanguageServer(sc)
	sc.Seal()
	return fbtest.New(t, sc)
}

func completionLabels(items []lsp.CompletionItem) map[string]lsp.CompletionItem {
	out := make(map[string]lsp.CompletionItem, len(items))
	for _, item := range items {
		out[item.Label] = item
	}
	return out
}

func TestCompletionEntityNamesAtRecordPosition(t *testing.T) {
	fixture := newIfcCompletionFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCDIRECTION((0.,0.,1.));
#2=IFCAXIS2PLACEMENT3D(#1,$,$);
#3=<|cursor>IFCDIRECTION((0.,0.,1.));
ENDSEC;
END-ISO-10303-21;`)
	doc.AssertNoParseErrors()

	items := completionLabels(doc.CompletionItems("cursor"))
	for _, want := range []string{"IFCDIRECTION", "IFCCARTESIANPOINT", "IFCAXIS2PLACEMENT3D", "IFCWALL"} {
		if _, ok := items[want]; !ok {
			t.Errorf("expected completion %q; got %d items", want, len(items))
		}
	}
	if item, ok := items["IFCDIRECTION"]; ok {
		if item.Kind != lsp.ClassCompletion {
			t.Errorf("IFCDIRECTION kind = %v, want ClassCompletion", item.Kind)
		}
		if item.Detail == "" {
			t.Error("IFCDIRECTION has no detail")
		}
		if item.Documentation != nil {
			t.Error("IFCDIRECTION documentation should be resolved lazily, not sent upfront")
		}
	}
}

func TestCompletionResolveAttachesEntityDocumentation(t *testing.T) {
	fixture := newIfcCompletionFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=<|cursor>IFCDIRECTION((0.,0.,1.));
ENDSEC;
END-ISO-10303-21;`)

	items := doc.CompletionItems("cursor")
	var direction *lsp.CompletionItem
	for i := range items {
		if items[i].Label == "IFCDIRECTION" {
			direction = &items[i]
			break
		}
	}
	if direction == nil {
		t.Fatal("IFCDIRECTION completion item not found")
	}

	server := fixture.Services().MustGet[lsp.Server]()
	resolved, err := server.ResolveCompletionItem(doc.Ctx(), direction)
	if err != nil {
		t.Fatalf("ResolveCompletionItem: %v", err)
	}
	if resolved == nil || resolved.Documentation == nil {
		t.Fatal("resolved IFCDIRECTION item has no documentation")
	}
}

func TestCompletionEntityNamesFilterByTypedPrefix(t *testing.T) {
	fixture := newIfcCompletionFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCAXIS2PLACEMENT3D($,$,$);
#2=IFCDI<|cursor>((0.,0.,1.));
ENDSEC;
END-ISO-10303-21;`)

	items := completionLabels(doc.CompletionItems("cursor"))
	if _, ok := items["IFCDIRECTION"]; !ok {
		t.Error("expected IFCDIRECTION for typed prefix IFCDI")
	}
	if _, ok := items["IFCWALL"]; ok {
		t.Error("IFCWALL should be filtered out by typed prefix")
	}
}

func TestCompletionHeaderKeywords(t *testing.T) {
	fixture := newIfcCompletionFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
<|cursor>FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
ENDSEC;
END-ISO-10303-21;`)

	items := completionLabels(doc.CompletionItems("cursor"))
	for _, want := range []string{"FILE_DESCRIPTION", "FILE_NAME", "FILE_SCHEMA"} {
		if _, ok := items[want]; !ok {
			t.Errorf("expected header keyword %q", want)
		}
	}
}

func TestCompletionEntityReferenceCandidates(t *testing.T) {
	fixture := newIfcCompletionFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCDIRECTION((0.,0.,1.));
#2=IFCCARTESIANPOINT((0.,0.,0.));
#3=IFCAXIS2PLACEMENT3D(<|cursor>#1,$,$);
ENDSEC;
END-ISO-10303-21;`)

	items := completionLabels(doc.CompletionItems("cursor"))
	// Param 0 expects IfcCartesianPoint, so only #2 is a valid candidate.
	if _, ok := items["#2"]; !ok {
		t.Error("expected entity reference #2")
	}
	if _, ok := items["#1"]; ok {
		t.Error("#1 (IFCDIRECTION) should be filtered out for a IfcCartesianPoint parameter")
	}
}

func TestCompletionEntityReferenceSecondParam(t *testing.T) {
	fixture := newIfcCompletionFixture(t)
	doc := fixture.Parse(`ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#1=IFCDIRECTION((0.,0.,1.));
#2=IFCCARTESIANPOINT((0.,0.,0.));
#3=IFCAXIS2PLACEMENT3D(#2,<|cursor>$,$);
ENDSEC;
END-ISO-10303-21;`)

	items := completionLabels(doc.CompletionItems("cursor"))
	// Param 1 (Axis) expects OPTIONAL IfcDirection, so only #1 is valid.
	if _, ok := items["#1"]; !ok {
		t.Error("expected entity reference #1 for IfcDirection parameter")
	}
	if _, ok := items["#2"]; ok {
		t.Error("#2 (IFCCARTESIANPOINT) should be filtered out for a IfcDirection parameter")
	}
}
