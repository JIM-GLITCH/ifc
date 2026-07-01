package ifc

import (
	"strings"
	"testing"
)

func TestCheckTypeWhereCompoundPlaneAngle(t *testing.T) {
	v := wvFromList([]WhereValue{
		parseLiteralNumber("1"),
		parseLiteralNumber("90"),
		parseLiteralNumber("0"),
		parseLiteralNumber("0"),
	})
	ok, detail := CheckTypeWhereRules("IfcCompoundPlaneAngleMeasure", v)
	if ok {
		t.Fatal("expected compound plane angle violation")
	}
	if detail == "" {
		t.Fatal("expected detail message")
	}
}

func TestCheckTypeWherePositiveInteger(t *testing.T) {
	ok, _ := CheckTypeWhereRules("IfcPositiveInteger", parseLiteralNumber("0"))
	if ok {
		t.Fatal("expected positive integer violation for 0")
	}
	ok, _ = CheckTypeWhereRules("IfcPositiveInteger", parseLiteralNumber("1"))
	if !ok {
		t.Fatal("expected 1 to pass")
	}
}

func TestCheckTypeWhereCompoundPlaneAngleSiteCoordinates(t *testing.T) {
	lat := wvFromList([]WhereValue{
		parseLiteralNumber("42"),
		parseLiteralNumber("21"),
		parseLiteralNumber("31"),
		parseLiteralNumber("181945"),
	})
	if ok, detail := CheckTypeWhereRules("IfcCompoundPlaneAngleMeasure", lat); !ok {
		t.Fatalf("expected latitude to pass type WHERE, got: %s", detail)
	}

	long := wvFromList([]WhereValue{
		parseLiteralNumber("-71"),
		parseLiteralNumber("-3"),
		parseLiteralNumber("-24"),
		parseLiteralNumber("-263305"),
	})
	if ok, detail := CheckTypeWhereRules("IfcCompoundPlaneAngleMeasure", long); !ok {
		t.Fatalf("expected longitude to pass type WHERE, got: %s", detail)
	}
}

func TestWvCompareRelNumericBeforeLexicographic(t *testing.T) {
	left := parseLiteralNumber("181945")
	right := parseLiteralNumber("1000000")
	result := wvCompareRel(left, right, "<")
	if !wvTruthy(result) {
		t.Fatal("expected 181945 < 1000000 numerically")
	}
}

func TestWvQueryNonZeroDirectionRatios(t *testing.T) {
	ratios := wvFromList([]WhereValue{
		parseLiteralNumber("1."),
		parseLiteralNumber("0."),
		parseLiteralNumber("0."),
	})
	nonZero := wvQuery(ratios, func(tmp WhereValue) WhereValue {
		return wvCompareRel(tmp, parseLiteralNumber("0.0"), "<>")
	})
	size := wvSizeOf(nonZero)
	if !wvTruthy(wvCompareRel(size, parseLiteralNumber("0"), ">")) {
		t.Fatalf("expected one non-zero direction ratio, got size %v", size)
	}
}

func TestEntityTypeOfQualifiedNames(t *testing.T) {
	names := entityTypeOfQualifiedNames("IFCLINE")
	if len(names) == 0 {
		t.Fatal("expected TYPEOF names for IFCLINE")
	}
	found := false
	for _, name := range names {
		if strings.Contains(strings.ToUpper(name), "IFCLINE") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("IFCLINE not in TYPEOF set: %v", names)
	}
}

func TestWvCompareRelEnumMembers(t *testing.T) {
	left := wvFromEnum(".NOCHANGE.")
	right := wvFromEnum("IfcChangeActionEnum.NOCHANGE")
	if !wvTruthy(wvCompareRel(left, right, "=")) {
		t.Fatal("expected .NOCHANGE. = IfcChangeActionEnum.NOCHANGE")
	}
}

func TestWvQueryAllZeroDirectionRatios(t *testing.T) {
	ratios := wvFromList([]WhereValue{
		parseLiteralNumber("0."),
		parseLiteralNumber("0."),
		parseLiteralNumber("0."),
	})
	nonZero := wvQuery(ratios, func(tmp WhereValue) WhereValue {
		return wvCompareRel(tmp, parseLiteralNumber("0.0"), "<>")
	})
	size := wvSizeOf(nonZero)
	if wvTruthy(wvCompareRel(size, parseLiteralNumber("0"), ">")) {
		t.Fatal("expected no non-zero direction ratios")
	}
}

func TestWvTruthyUnknownIsFalse(t *testing.T) {
	if wvTruthy(wvUnknownVal()) {
		t.Fatal("expected unknown to be falsy in WHERE validation")
	}
	if wvTruthy(wvAbsentVal()) {
		t.Fatal("expected absent to be falsy in WHERE validation")
	}
	if wvTruthy(wvUnknownLogical()) {
		t.Fatal("expected unknown logical to be falsy")
	}
}

func TestWvAndUnknownDoesNotPass(t *testing.T) {
	trueLit := WhereValue{kind: wvKindLogical, log: 1}
	and := wvAnd(trueLit, wvUnknownVal())
	if wvTruthy(and) {
		t.Fatal("true AND unknown should not pass strict WHERE")
	}
}

func TestWvOrExistsShortCircuit(t *testing.T) {
	notExists := wvNot(wvExists(wvAbsentVal()))
	or := wvOr(notExists, wvUnknownVal())
	if !wvTruthy(or) {
		t.Fatal("NOT EXISTS(absent) OR unknown should pass")
	}
}

func TestWvNVLAbsentUsesFallback(t *testing.T) {
	fallback := parseLiteralNumber("1.0")
	got := wvNVL(wvAbsentVal(), fallback)
	if !wvTruthy(wvCompareRel(got, fallback, "=")) {
		t.Fatal("NVL(absent, fallback) should return fallback")
	}
}
