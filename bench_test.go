package ifc

import (
	"context"
	"os"
	"strings"
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
	sinkUnits     []entityUnit
	sinkContent   []byte
	sinkScanOK    bool
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
			// The validator caches per-document entity state; drop it so the
			// benchmark keeps measuring a full validation instead of the
			// unchanged-document fast path.
			if incremental, ok := validator.(*IfcDocumentValidator); ok {
				defer incremental.cache.reset()
			}
			for i := 0; i < b.N; i++ {
				sinkDiags = validator.Validate(ctx, doc.Document, "bench")
				if incremental, ok := validator.(*IfcDocumentValidator); ok {
					incremental.cache.reset()
				}
			}
		})
	}
}

// BenchmarkValidateIncrementalEdit measures what a didChange costs when only a
// single entity instance changed. Both revisions are parsed before the timer
// starts, so the number covers validation alone: population build, fingerprint
// pass, dependency closure and the re-validated subset.
//
// The two documents alternate, so every iteration after the first two
// re-validates a revision the cache has not seen yet, which is exactly the
// steady state of typing in a large file.
//
// midFile edits an instance with almost no referrers, so the dependency
// closure is tight and the number shows what the reuse itself is worth.
// root edits the schema root instance, which a large share of the file points
// at; that closure is inherently large and is the worst case for reuse.
func BenchmarkValidateIncrementalEdit(b *testing.B) {
	for _, target := range []string{"midFile", "root"} {
		b.Run(target, func(b *testing.B) {
			benchmarkIncrementalEdit(b, target)
		})
	}
}

func benchmarkIncrementalEdit(b *testing.B, target string) {
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		b.Fatal(err)
	}
	original := string(raw)
	edited := original
	editAt := -1
	if target == "root" {
		editAt = strings.Index(edited, "IFCPROJECT(")
	} else {
		// The instance roughly in the middle of the file.
		editAt = nthEntityOffset(edited, len(edited)/2)
	}
	if editAt < 0 {
		b.Fatalf("target %q not found in t.ifc", target)
	}
	// Skip the "#<number>=" prefix: editing the instance name would rename the
	// instance, which is a different (and much more expensive) scenario.
	for editAt < len(edited) && edited[editAt] != '=' {
		editAt++
	}
	for i := editAt; i < len(edited); i++ {
		if edited[i] >= '0' && edited[i] <= '9' {
			replacement := byte('0')
			if edited[i] == '0' {
				replacement = '1'
			}
			edited = edited[:i] + string(replacement) + edited[i+1:]
			break
		}
	}
	if edited == original {
		b.Fatal("edit did not change the document")
	}

	sc := CreateServices()
	fixture := fbtest.New(b, sc)
	documents := []*core.Document{
		fixture.ParseURI(original, "inmemory://incremental-bench.ifc").Document,
		fixture.ParseURI(edited, "inmemory://incremental-bench.ifc").Document,
	}
	validator := service.MustGet[workspace.DocumentValidator](sc)
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkDiags = validator.Validate(ctx, documents[i%len(documents)], "bench")
	}
}

// nthEntityOffset returns the offset of the entity instance that contains the
// byte offset at.
func nthEntityOffset(text string, at int) int {
	if at <= 0 || at >= len(text) {
		return -1
	}
	best := -1
	for i := 0; i+2 < len(text); i++ {
		if text[i] != '#' {
			continue
		}
		j := i + 1
		for j < len(text) && text[j] >= '0' && text[j] <= '9' {
			j++
		}
		if j > i+1 && j < len(text) && text[j] == '=' {
			if i > at {
				break
			}
			best = i
		}
	}
	return best
}

// BenchmarkEntityScan measures the per-validation bookkeeping of the
// incremental path: collecting every entity instance plus fingerprinting its
// source text. This is the fixed cost that every incremental validation pays
// before it can decide what to re-validate.
func BenchmarkEntityScan(b *testing.B) {
	_, doc := benchTIfc(b, nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkUnits, sinkContent, sinkScanOK = collectEntityUnits(doc.Document)
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

// BenchmarkLink measures build phase 2 (symbol tables + reference linking)
// alone, by resetting a fully built document back to the parsed state and
// rebuilding it with all validation tiers off. The remaining cost is phases
// 1-2 plus the tier-off validation walk (population build + AST traversal).
//
// Note: tiers must be disabled through the ValidationSettingsStore, not the
// context - IfcDocumentValidator.Validate re-reads settings from the store and
// overrides whatever the incoming context carries.
func BenchmarkLink(b *testing.B) {
	sc := CreateServices()
	store := service.MustGet[*ValidationSettingsStore](sc)
	store.Set(ValidationSettings{})
	fixture := fbtest.New(b, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		b.Fatal(err)
	}
	doc := fixture.Parse(string(raw))
	lock := service.MustGet[workspace.Lock](sc)
	builder := service.MustGet[workspace.Builder](sc)
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lock.Write(ctx, func(ctx context.Context, downgrade func()) {
			builder.Reset(doc.Document, core.DocStateParsed)
			if err := builder.Build(ctx, []*core.Document{doc.Document}, downgrade); err != nil {
				b.Fatal(err)
			}
		})
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
