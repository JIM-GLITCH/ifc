package ifc

import "testing"

// TestLookupTypeWhereRulesCasing pins the case-insensitive lookup contract:
// every spelling of a type name that carries WHERE rules must resolve, and
// every spelling of a type without them must miss.
func TestLookupTypeWhereRulesCasing(t *testing.T) {
	withRules := ""
	for name := range TypeWhereRules {
		withRules = name
		break
	}
	if withRules == "" {
		t.Fatal("TypeWhereRules is empty; the generated schema is missing")
	}
	withoutRules := ""
	for name := range Types {
		if _, ok := TypeWhereRules[name]; ok {
			continue
		}
		withoutRules = name
		break
	}
	if withoutRules == "" {
		t.Fatal("no type without WHERE rules found")
	}

	for _, spelling := range []string{withRules, upperTypeName(withRules), lowerFirstRune(withRules)} {
		rules, ok := LookupTypeWhereRules(spelling)
		if !ok || len(rules) == 0 {
			t.Errorf("LookupTypeWhereRules(%q) = (%d rules, %v); want the rules of %q", spelling, len(rules), ok, withRules)
		}
	}
	for _, spelling := range []string{withoutRules, upperTypeName(withoutRules), lowerFirstRune(withoutRules)} {
		if rules, ok := LookupTypeWhereRules(spelling); ok {
			t.Errorf("LookupTypeWhereRules(%q) = (%d rules, true); want a miss", spelling, len(rules))
		}
	}
}

func lowerFirstRune(s string) string {
	if s == "" {
		return s
	}
	return string(s[0]|0x20) + s[1:]
}

func TestUpperTypeName(t *testing.T) {
	cases := map[string]string{
		"":                             "",
		"IfcLabel":                     "IFCLABEL",
		"IFCLABEL":                     "IFCLABEL",
		"ifcCompoundPlaneAngleMeasure": "IFCCOMPOUNDPLANEANGLEMEASURE",
		"LIST [1:?] OF IfcLengthMeasure": "LIST [1:?] OF IFCLENGTHMEASURE",
		"already_upper_123":            "ALREADY_UPPER_123",
	}
	for in, want := range cases {
		if got := upperTypeName(in); got != want {
			t.Errorf("upperTypeName(%q) = %q; want %q", in, got, want)
		}
	}
}
