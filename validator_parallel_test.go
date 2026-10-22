package ifc

import (
	"context"
	"fmt"
	"os"
	"testing"

	core "typefox.dev/fastbelt"
	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

// TestValidateDeterministicUnderParallelism runs the full validation of the
// reference model several times and requires identical diagnostics every
// time. Parallel validation merges per-worker results, so any scheduling
// dependence (or shared-state race that corrupts output) shows up as a
// difference between runs.
func TestValidateDeterministicUnderParallelism(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	doc := fixture.Parse(string(raw))
	validator := service.MustGet[workspace.DocumentValidator](sc)

	// t.ifc has a known parse error; diagnostics still include the parallel
	// entity-validation results, which is what this test compares.
	render := func(diags []*core.Diagnostic) string {
		parts := make([]string, 0, len(diags))
		for _, d := range diags {
			parts = append(parts, fmt.Sprintf("%d:%d-%d:%d|%s|%s|%s",
				d.Range.Start.Line, d.Range.Start.Column,
				d.Range.End.Line, d.Range.End.Column,
				d.Code, d.Severity, d.Message))
		}
		return fmt.Sprintf("%d diagnostics:\n%s", len(diags), joinStrings(parts, "\n"))
	}

	first := render(validator.Validate(context.Background(), doc.Document, "test"))
	t.Logf("baseline: %d diagnostics", countLines(first))
	for i := 0; i < 7; i++ {
		got := render(validator.Validate(context.Background(), doc.Document, "test"))
		if got != first {
			t.Fatalf("validation output changed between runs (iteration %d)", i+2)
		}
	}
}

func countLines(s string) int {
	n := 1
	for _, c := range s {
		if c == '\n' {
			n++
		}
	}
	return n
}

func joinStrings(parts []string, sep string) string {
	out := ""
	for i, p := range parts {
		if i > 0 {
			out += sep
		}
		out += p
	}
	return out
}
