package ifc

import (
	"context"
	"testing"
)

func TestEntityWhereEvalUnknownMessageIncludesRuleLabel(t *testing.T) {
	ec := NewEntityWhereEval(context.Background(), "IFCBUILDING", nil, nil, nil)
	ec.ruleLabel = "UniquePropertySetNames"
	ec.unknown("Cannot evaluate parameter value")
	notes := ec.takeEvalUnknowns()
	if len(notes) != 1 {
		t.Fatalf("expected 1 note, got %v", notes)
	}
	want := `Entity "IFCBUILDING" WHERE "UniquePropertySetNames": Cannot evaluate parameter value`
	if notes[0] != want {
		t.Fatalf("got %q want %q", notes[0], want)
	}
}

func TestEntityWhereEvalWvOrLazyStopsAfterEvalUnknown(t *testing.T) {
	ec := NewEntityWhereEval(context.Background(), "IFCAXIS2PLACEMENT3D", nil, nil, nil)
	ec.ruleLabel = "AxisIs3D"
	calls := 0
	result := ec.WvOrLazy(
		func() WhereValue {
			return ec.unknown("Cannot evaluate parameter value")
		},
		func() WhereValue {
			calls++
			return parseLiteralNumber("1")
		},
	)
	if calls != 0 {
		t.Fatalf("expected OR right branch to be skipped after eval-unknown, calls=%d", calls)
	}
	if result.kind != wvKindUnknown {
		t.Fatalf("expected unknown result, got %v", result.kind)
	}
	if len(*ec.evalUnknowns) != 1 {
		t.Fatalf("expected single eval-unknown note, got %v", *ec.evalUnknowns)
	}
}

func TestCheckEntityWhereRulesStopsAfterFirstEvalUnknown(t *testing.T) {
	rules := []EntityWhereRuleSpec{
		{Label: "WR1", Check: func(ec *EntityWhereEval) (bool, string) {
			ec.unknown("first failure")
			return false, "WR1 failed"
		}},
		{Label: "WR2", Check: func(ec *EntityWhereEval) (bool, string) {
			ec.unknown("second failure")
			return false, "WR2 failed"
		}},
	}
	var evalUnknowns []string
	for _, rule := range rules {
		ec := NewEntityWhereEval(context.Background(), "IFCTEST", nil, nil, nil)
		ec.ruleLabel = rule.Label
		_, _ = rule.Check(ec)
		ruleUnknowns := ec.takeEvalUnknowns()
		evalUnknowns = append(evalUnknowns, ruleUnknowns...)
		if len(ruleUnknowns) > 0 {
			break
		}
	}
	if len(evalUnknowns) != 1 {
		t.Fatalf("expected 1 eval-unknown after early stop, got %v", evalUnknowns)
	}
	if evalUnknowns[0] != `Entity "IFCTEST" WHERE "WR1": first failure` {
		t.Fatalf("unexpected message: %q", evalUnknowns[0])
	}
}
