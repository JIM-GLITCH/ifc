package ifc

import (
	"context"
	"os"
	"testing"

	core "typefox.dev/fastbelt"
	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

// Sinks to keep benchmark results alive across iterations.
var (
	sinkDiags     []*core.Diagnostic
	sinkPop       *Population
	sinkViolation []UniqueViolation
	sinkBool      bool
	sinkString    string
	sinkInt       int
	sinkKind      ParamKind
	sinkKinds     []ParamKind
	sinkForward   map[int][]ForwardInverseTarget
	sinkKey       string
)

// benchTIfc parses the t.ifc fixture through the full build pipeline
// (parse -> symbols -> link -> validate) and returns the service container
// together with the built document.
//
// If preset is non-nil, the validation tiers are applied to the container's
// settings store before the build runs; they also govern later Validate calls
// made on the returned container. A nil preset keeps the default (all tiers).
//
// All setup happens before b.ResetTimer, so timings cover only the measured
// operation, not file IO or the initial build.
func benchTIfc(b *testing.B, preset *ValidationSettings) (*service.Container, *fbtest.Doc) {
	b.Helper()
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		b.Fatal(err)
	}
	sc := CreateServices()
	if preset != nil {
		ApplyValidationSettings(sc, *preset)
	}
	doc := fbtest.New(b, sc).Parse(string(raw))
	// Note: t.ifc triggers a pre-existing parser recovery error (see
	// TestTIfcBuilding150UniquePropertySetNames), so only require an AST.
	if doc.Document.Root == nil {
		b.Fatal("t.ifc produced no AST")
	}
	b.ResetTimer()
	return sc, doc
}

// BenchmarkValidate measures the phase-3 validation hot path for a large IFC
// document, broken down by validation tier. This is the number targeted by the
// memoize optimizations.
//
// The no-tiers case isolates the AST walk itself; structure+types adds
// parameter and type checks; all-tiers additionally runs entity WHERE rules,
// global RULEs and UNIQUE constraints.
func BenchmarkValidate(b *testing.B) {
	noTiers := ValidationSettings{}
	structureTypes := ValidationSettings{Structure: true, Types: true}
	cases := []struct {
		name   string
		preset *ValidationSettings
	}{
		{"no-tiers", &noTiers},
		{"structure+types", &structureTypes},
		{"all-tiers", nil},
	}
	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			sc, doc := benchTIfc(b, tc.preset)
			validator := service.MustGet[workspace.DocumentValidator](sc)
			ctx := context.Background()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sinkDiags = validator.Validate(ctx, doc.Document, "bench")
			}
		})
	}
}

// BenchmarkParse isolates the lexer + parser cost (phase 1) for comparison
// with the validation numbers.
func BenchmarkParse(b *testing.B) {
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		b.Fatal(err)
	}
	sc := CreateServices()
	parser := service.MustGet[workspace.DocumentParser](sc)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doc, err := core.NewDocumentFromString("inmemory://test.ifc", "ifc", string(raw))
		if err != nil {
			b.Fatal(err)
		}
		parser.Parse(doc)
		if doc.Root == nil {
			b.Fatal("parse produced no AST")
		}
	}
}

// BenchmarkBuildPopulation measures the population and inverse-link
// construction performed at the start of every Validate call.
func BenchmarkBuildPopulation(b *testing.B) {
	_, doc := benchTIfc(b, nil)
	ctx := context.Background()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkPop = BuildPopulation(ctx, doc.Document)
	}
}

// BenchmarkCheckGlobalRules measures global RULE evaluation, including the
// entity-set construction it performs.
func BenchmarkCheckGlobalRules(b *testing.B) {
	_, doc := benchTIfc(b, nil)
	pop := BuildPopulation(context.Background(), doc.Document)
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkBool, sinkString, _ = CheckGlobalRules(ctx, pop)
	}
}

// BenchmarkCheckUniqueRules measures UNIQUE constraint evaluation over the
// owner entity sets (IfcRoot/GlobalId dominates real files).
func BenchmarkCheckUniqueRules(b *testing.B) {
	_, doc := benchTIfc(b, nil)
	pop := BuildPopulation(context.Background(), doc.Document)
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkViolation, _ = CheckUniqueRules(ctx, pop)
	}
}

// BenchmarkMergedForwardInverses measures the per-instance forward-inverse
// merge performed while building inverse links.
func BenchmarkMergedForwardInverses(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkForward = mergedForwardInverses("IFCPROJECT")
	}
}

// BenchmarkAttrIndex measures attribute-name resolution used by every WHERE
// rule attribute access.
func BenchmarkAttrIndex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkInt = attrIndex("IFCPROJECT", "GlobalId")
	}
}

// BenchmarkAllowedSelectKinds measures SELECT type member resolution performed
// per parameter during type validation.
func BenchmarkAllowedSelectKinds(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkKinds = allowedSelectKinds("IfcActorSelect")
	}
}

// BenchmarkResolvedTypeValueKind measures type alias chain resolution.
func BenchmarkResolvedTypeValueKind(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkKind = resolvedTypeValueKind("IfcPositiveLengthMeasure")
	}
}

// BenchmarkEntityAttrCacheKey measures the eval-cache key construction done on
// every attribute access during WHERE rule evaluation.
func BenchmarkEntityAttrCacheKey(b *testing.B) {
	_, doc := benchTIfc(b, nil)
	pop := BuildPopulation(context.Background(), doc.Document)
	rec, ok := pop.Lookup("#1")
	if !ok || rec.Instance == nil {
		b.Fatal("#1 not found in population")
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkKey = entityAttrCacheKey(rec.Instance, "GlobalId")
	}
}
