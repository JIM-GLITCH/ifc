package ifc

import (
	"context"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
)

func TestEntityEvalCacheReusesDerivedAttr(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)

	source := `ISO-10303-21;
HEADER;
FILE_DESCRIPTION((),'2;1');
FILE_NAME('','2022-01-01',(),(),'','','');
FILE_SCHEMA(('IFC4'));
ENDSEC;
DATA;
#130=IFCSIUNIT(*,.LENGTHUNIT.,$,.METRE.);
ENDSEC;
END-ISO-10303-21;`

	doc := fixture.Parse(source)
	doc.AssertNoParseErrors()
	pop := BuildPopulation(context.Background(), doc.Document)
	ctx := WithPopulation(context.Background(), pop)
	ctx = WithEntityEvalCache(ctx)

	rec, ok := pop.Lookup("#130")
	if !ok || rec.Instance == nil {
		t.Fatal("expected #130 in population")
	}
	inst := rec.Instance
	ec1 := NewEntityWhereEval(ctx, "IFCSIUNIT", inst, instanceParams(inst), pop)
	first := evalDerivedAttr(ec1, inst, "Dimensions")
	if first.kind == wvKindUnknown {
		t.Fatalf("expected derived Dimensions, got unknown")
	}

	ec2 := NewEntityWhereEval(ctx, "IFCSIUNIT", inst, instanceParams(inst), pop)
	second := evalDerivedAttr(ec2, inst, "Dimensions")
	if second.kind != first.kind {
		t.Fatalf("cache returned different kind: first=%v second=%v", first.kind, second.kind)
	}

	cache := entityEvalCacheFromContext(ctx)
	if cache == nil || len(cache.values) == 0 {
		t.Fatal("expected cache entries after derived eval")
	}
}
