package ifc

import (
	"fmt"
	"os"
	"strings"
	"testing"

	core "typefox.dev/fastbelt"
	fbtest "typefox.dev/fastbelt/test"
)

const incrementalTestURI = "inmemory://incremental.ifc"

// renderDiagnostics flattens diagnostics into comparable strings. Everything
// that reaches the editor is covered: position, severity, code, source and
// message.
func renderDiagnostics(diags []*core.Diagnostic) []string {
	out := make([]string, 0, len(diags))
	for _, d := range diags {
		out = append(out, fmt.Sprintf("%d:%d-%d:%d|%d|%s|%s|%s",
			d.Range.Start.Line, d.Range.Start.Column,
			d.Range.End.Line, d.Range.End.Column,
			d.Severity, d.Code, d.Source, d.Message))
	}
	return out
}

// nthEntityStart returns the byte offset of the nth "#<n>=" instance start.
func nthEntityStart(text string, n int) int {
	count := 0
	for i := 0; i+2 < len(text); i++ {
		if text[i] != '#' {
			continue
		}
		j := i + 1
		for j < len(text) && text[j] >= '0' && text[j] <= '9' {
			j++
		}
		if j > i+1 && j < len(text) && text[j] == '=' {
			count++
			if count == n {
				return i
			}
		}
	}
	return -1
}

// entityTextEnd returns the offset just past the ";" terminating the instance
// that starts at start.
func entityTextEnd(text string, start int) int {
	for i := start; i < len(text); i++ {
		if text[i] == ';' {
			return i + 1
		}
	}
	return len(text)
}

func firstDigitAfter(text string, from int) int {
	for i := from; i < len(text); i++ {
		if text[i] >= '0' && text[i] <= '9' {
			return i
		}
	}
	return -1
}

// incrementalEdits are single-document edits. reuse states whether the edit is
// expected to take the incremental path; a false entry must fall back to a
// full validation. Either way the published diagnostics have to match a cold
// full validation exactly.
func incrementalEdits() []struct {
	name  string
	reuse bool
	apply func(string) string
} {
	return []struct {
		name  string
		reuse bool
		apply func(string) string
	}{
		{
			// Every data instance moves down one line, no instance text
			// changes: pure re-anchoring of cached diagnostics.
			name:  "shift-all",
			reuse: true,
			apply: func(text string) string {
				idx := strings.Index(text, "ENDSEC;")
				if idx < 0 {
					return text
				}
				return text[:idx+len("ENDSEC;")] + "\n" + text[idx+len("ENDSEC;"):]
			},
		},
		{
			// Insert whitespace in front of a mid-file instance: everything
			// after it moves, nothing else changes.
			name:  "shift-from-middle",
			reuse: true,
			apply: func(text string) string {
				idx := nthEntityStart(text, 12000)
				if idx < 0 {
					return text
				}
				return text[:idx] + " " + text[idx:]
			},
		},
		{
			// Edit one digit inside a single instance.
			name:  "edit-one-instance",
			reuse: false,
			apply: func(text string) string {
				start := nthEntityStart(text, 6000)
				if start < 0 {
					return text
				}
				end := entityTextEnd(text, start)
				digit := firstDigitAfter(text, start)
				if digit < 0 || digit >= end {
					return text
				}
				replacement := byte('0')
				if text[digit] == '0' {
					replacement = '1'
				}
				return text[:digit] + string(replacement) + text[digit+1:]
			},
		},
		{
			// Edit the instance that carries IFC4's root entity: it is
			// referenced by a large share of the file, so this exercises the
			// dependency closure.
			name:  "edit-root-entity",
			reuse: true,
			apply: func(text string) string {
				idx := strings.Index(text, "IFCPROJECT(")
				if idx < 0 {
					return text
				}
				digit := firstDigitAfter(text, idx)
				if digit < 0 {
					return text
				}
				replacement := byte('0')
				if text[digit] == '0' {
					replacement = '1'
				}
				return text[:digit] + string(replacement) + text[digit+1:]
			},
		},
		{
			// Remove a whole instance: the entity set shrinks, which must
			// fall back to a full validation rather than guess.
			name:  "delete-instance",
			reuse: false,
			apply: func(text string) string {
				start := nthEntityStart(text, 4000)
				if start < 0 {
					return text
				}
				lineStart := strings.LastIndex(text[:start], "\n") + 1
				lineEnd := strings.Index(text[start:], "\n")
				if lineEnd < 0 {
					return text
				}
				return text[:lineStart] + text[start+lineEnd+1:]
			},
		},
		{
			// Append a renamed duplicate of an existing instance: the entity
			// set grows and a previously dangling reference may resolve.
			name:  "add-instance",
			reuse: true,
			apply: func(text string) string {
				start := nthEntityStart(text, 2000)
				if start < 0 {
					return text
				}
				end := entityTextEnd(text, start)
				body := strings.TrimSuffix(text[start:end], ";")
				body = body[strings.Index(body, "=")+1:]
				insert := "#999999=" + body + ";\n"
				lastEndsec := strings.LastIndex(text, "ENDSEC;")
				if lastEndsec < 0 {
					return text
				}
				return text[:lastEndsec] + insert + text[lastEndsec:]
			},
		},
	}
}

// TestIncrementalValidationMatchesFull validates the same document twice: once
// through a warm validator that already validated the previous revision, and
// once through a cold validator that only sees the new revision. The published
// diagnostics must be identical.
func TestIncrementalValidationMatchesFull(t *testing.T) {
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	original := string(raw)

	for _, edit := range incrementalEdits() {
		t.Run(edit.name, func(t *testing.T) {
			edited := edit.apply(original)
			if edited == original {
				t.Fatalf("edit %q did not change the document", edit.name)
			}

			warmContainer := CreateServices()
			warmFixture := fbtest.New(t, warmContainer)
			warmFixture.ParseURI(original, incrementalTestURI)
			incrementalReuseRuns.Store(0)
			incrementalFullRuns.Store(0)
			warm := warmFixture.ParseURI(edited, incrementalTestURI)

			// Read the counters before the cold run: the cold container
			// validates too and would add its own full run to them.
			reuseAfterEdit, fullAfterEdit := incrementalReuseRuns.Load(), incrementalFullRuns.Load()

			coldContainer := CreateServices()
			coldFixture := fbtest.New(t, coldContainer)
			cold := coldFixture.ParseURI(edited, incrementalTestURI+"-cold")

			want := renderDiagnostics(cold.Document.Diagnostics)
			got := renderDiagnostics(warm.Document.Diagnostics)
			if len(want) != len(got) {
				t.Fatalf("diagnostic count: full=%d incremental=%d\nfirst extra: %v",
					len(want), len(got), firstDifference(want, got))
			}
			for i := range want {
				if want[i] != got[i] {
					t.Fatalf("diagnostic %d differs:\n full       %q\n incremental %q", i, want[i], got[i])
				}
			}
			if edit.reuse {
				if reuseAfterEdit != 1 || fullAfterEdit != 0 {
					t.Fatalf("edit %q: expected the incremental path, got reuse=%d full=%d",
						edit.name, reuseAfterEdit, fullAfterEdit)
				}
			} else if reuseAfterEdit != 0 || fullAfterEdit != 1 {
				t.Fatalf("edit %q: expected a full fallback, got reuse=%d full=%d",
					edit.name, reuseAfterEdit, fullAfterEdit)
			}
			t.Logf("%s: %d diagnostics, reuse=%d full=%d", edit.name, len(got),
				reuseAfterEdit, fullAfterEdit)
		})
	}
}

func firstDifference(want, got []string) string {
	for i := 0; i < len(want) && i < len(got); i++ {
		if want[i] != got[i] {
			return fmt.Sprintf("index %d: full=%q incremental=%q", i, want[i], got[i])
		}
	}
	return ""
}

// TestIncrementalAffectedSet checks the dependency closure directly: an edit
// must mark the edited instance, everything it points at, and everything that
// points at any of those.
func TestIncrementalAffectedSet(t *testing.T) {
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	doc := fbtest.New(t, CreateServices()).ParseURI(string(raw), incrementalTestURI+"-affected").Document
	units, text, ok := collectEntityUnits(doc)
	if !ok {
		t.Fatal("collectEntityUnits failed")
	}

	state := &documentValidationState{entities: map[int]cachedEntity{}}
	for _, u := range units {
		state.entities[u.ordinal] = cachedEntity{fingerprint: u.hash}
	}
	if _, count, ok := computeAffected(units, text, state); !ok || count != 0 {
		t.Fatalf("unchanged document: count=%d ok=%v; want 0, true", count, ok)
	}

	// Flip one instance's fingerprint: that instance must always be marked,
	// plus its dependency closure.
	victim := units[len(units)/2]
	state.entities[victim.ordinal] = cachedEntity{fingerprint: victim.hash ^ 0xff}
	recompute, count, ok := computeAffected(units, text, state)
	if !ok || count < 1 {
		t.Fatalf("single edit: count=%d ok=%v; want >= 1", count, ok)
	}
	marked := false
	for i, u := range units {
		if u.ordinal == victim.ordinal {
			marked = recompute[i]
		}
	}
	if !marked {
		t.Fatal("the edited instance itself was not marked for re-validation")
	}
	t.Logf("single-instance edit marks %d of %d instances", count, len(units))

	// An instance that disappears cannot be handled: its inverse
	// contributions are unknown, and its text is gone.
	shrunk := make([]entityUnit, 0, len(units)-1)
	victimIndex := -1
	for i, u := range units {
		if u.ordinal == victim.ordinal {
			victimIndex = i
			continue
		}
		shrunk = append(shrunk, u)
	}
	if victimIndex < 0 {
		t.Fatal("victim not found in units")
	}
	full := &documentValidationState{entities: map[int]cachedEntity{}}
	for _, u := range units {
		full.entities[u.ordinal] = cachedEntity{fingerprint: u.hash}
	}
	if _, _, ok := computeAffected(shrunk, text, full); ok {
		t.Fatal("a removed instance must disable the incremental path")
	}
}
