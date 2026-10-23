package ifc

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"

	core "typefox.dev/fastbelt"
)

// Incremental entity validation.
//
// Every didChange re-parses the document and hands the validator a brand new
// core.Document with a brand new AST, so a naive validator pays the full
// validation cost of the whole file even when a single entity changed. This
// file keeps a per-document cache of entity instance fingerprints plus the
// diagnostics they produced, and re-runs validation only for the entities a
// change can actually reach.
//
// An entity's diagnostics are a pure function of:
//
//   - its own source text (captured by the fingerprint),
//   - the entities it can reach by reference, including their inverse sets,
//   - the schema header and the active validation settings.
//
// So an entity may be re-used unless the change is observable from it. A
// change to entity E is observable by:
//
//   - E itself, because its text changed or it appeared;
//   - anything E points at, because E moving changes their inverse sets;
//   - anything pointing at E, or sharing a referent with E, because a WHERE
//     rule or a DERIVE clause may read an inverse attribute of an entity it
//     only points at. Co-referencing is therefore treated as a dependency.
//
// This is deliberately coarser than a true dependency graph: it never reports
// a smaller set than the real one, and building it costs a single pass over
// the document text. Entity validation itself stays results-independent (see
// parallelValidate), so the re-validated subset can be distributed across
// workers exactly like the full walk.

// incrementalEntityThreshold is the number of entity instances below which a
// document is always validated in full: the per-entity bookkeeping costs more
// than the validation it would avoid.
const incrementalEntityThreshold = 256

// incrementalChangedBudget is 1/N: when more than 1/N of the entities have to
// be re-validated, a full validation is cheaper than diffing plus reuse.
const incrementalChangedBudget = 4

// incrementalCacheMaxDocs bounds how many documents keep validation state.
const incrementalCacheMaxDocs = 4

// fnvOffset64 and fnvPrime64 are the FNV-1a 64 bit constants.
const (
	fnvOffset64 = uint64(0xcbf29ce484222325)
	fnvPrime64  = uint64(0x100000001b3)
)

// entityUnit is one entity instance as seen by validation.
type entityUnit struct {
	node    core.Validator
	ordinal int
	hash    uint64
	segment core.TextSegment
}

// text returns the source bytes covered by the unit.
func (u *entityUnit) text(content []byte) []byte {
	return segmentBytes(content, &u.segment)
}

// cachedDiagnostic stores a diagnostic in coordinates relative to the start of
// the entity instance it was reported on, so it can be re-anchored after the
// surrounding text moved.
//
// Columns on the entity's own first line are relative to the entity start
// column; columns on any later line are absolute, because those lines belong
// to the entity's own, unchanged, text.
type cachedDiagnostic struct {
	startLine int32
	startCol  int32
	endLine   int32
	endCol    int32
	severity  core.DiagnosticSeverity
	message   string
	source    string
	code      string
	tags      []core.DiagnosticTag
	codeDesc  *core.DiagnosticCodeDescription
	data      any
}

// diagnostic re-anchors the cached diagnostic at the entity's current start.
func (c cachedDiagnostic) diagnostic(start core.TextLocation) *core.Diagnostic {
	startCol := core.TextColumn(c.startCol)
	endCol := core.TextColumn(c.endCol)
	if c.startLine == 0 {
		startCol += start.Column
	}
	if c.endLine == 0 {
		endCol += start.Column
	}
	return &core.Diagnostic{
		Range: core.TextRange{
			Start: core.TextLocation{
				Line:   core.TextLine(int32(start.Line) + c.startLine),
				Column: startCol,
			},
			End: core.TextLocation{
				Line:   core.TextLine(int32(start.Line) + c.endLine),
				Column: endCol,
			},
		},
		Severity:       c.severity,
		Message:        c.message,
		Source:         c.source,
		Code:           c.code,
		CodeDescription: c.codeDesc,
		Tags:           c.tags,
		Data:           c.data,
	}
}

type cachedEntity struct {
	fingerprint uint64
	diags       []cachedDiagnostic
}

// documentValidationState is everything the next validation of the same
// document needs in order to skip the entities a change cannot reach.
type documentValidationState struct {
	header      uint64
	settings    ValidationSettings
	level       string
	lexErrors   int
	parseErrors int
	entities    map[int]cachedEntity
}

// compatible reports whether the state may be reused for a document that now
// has the given header fingerprint, settings, level and front-end error counts.
func (s *documentValidationState) compatible(
	header uint64,
	settings ValidationSettings,
	level string,
	lexErrors, parseErrors int,
) bool {
	if s == nil {
		return false
	}
	return s.header == header &&
		s.settings == settings &&
		s.level == level &&
		s.lexErrors == lexErrors &&
		s.parseErrors == parseErrors
}

// validationStateCache keeps the validation state of the last few documents.
type validationStateCache struct {
	mu    sync.Mutex
	order []string
	docs  map[string]*documentValidationState
}

func (c *validationStateCache) get(uri string) *documentValidationState {
	if uri == "" {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.docs[uri]
}

func (c *validationStateCache) put(uri string, state *documentValidationState) {
	if uri == "" || state == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.docs == nil {
		c.docs = make(map[string]*documentValidationState, incrementalCacheMaxDocs)
	}
	if _, ok := c.docs[uri]; !ok {
		c.order = append(c.order, uri)
	}
	c.docs[uri] = state
	for len(c.order) > incrementalCacheMaxDocs {
		oldest := c.order[0]
		c.order = c.order[1:]
		delete(c.docs, oldest)
	}
}

// reset drops every cached document. Benchmarks use it to keep measuring a
// full validation instead of the unchanged-document fast path.
func (c *validationStateCache) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.order = nil
	c.docs = nil
}

// collectEntityUnits walks the document and returns every entity instance that
// takes part in validation, in document order, together with the document text.
//
// ok is false when the document cannot be validated entity by entity: any
// validator node that is not a "#<number>" entity instance makes the whole
// document fall back to a full validation.
func collectEntityUnits(doc *core.Document) (units []entityUnit, content []byte, ok bool) {
	if doc == nil || doc.Root == nil || doc.TextDoc == nil {
		return nil, nil, false
	}
	content = doc.TextDoc.Content()
	for node := range core.AllNodes(doc.Root) {
		validator, isValidator := node.(core.Validator)
		if !isValidator {
			continue
		}
		var name string
		switch inst := node.(type) {
		case SIMPLE_ENTITY_INSTANCE:
			name = inst.Name()
		case COMPLEX_ENTITY_INSTANCE:
			name = inst.Name()
		default:
			return nil, nil, false
		}
		ordinal := parseInstanceOrdinal(name)
		if ordinal <= 0 {
			return nil, nil, false
		}
		segment := node.Segment()
		if segment == nil {
			return nil, nil, false
		}
		units = append(units, entityUnit{
			node:    validator,
			ordinal: ordinal,
			hash:    hashBytes(segmentBytes(content, segment)),
			segment: *segment,
		})
	}
	return units, content, true
}

func segmentBytes(content []byte, segment *core.TextSegment) []byte {
	start := int(segment.Indices.Start)
	end := int(segment.Indices.End)
	if start < 0 {
		start = 0
	}
	if end > len(content) {
		end = len(content)
	}
	if end <= start {
		return nil
	}
	return content[start:end]
}

// hashBytes is FNV-1a 64 bit. It runs at roughly one byte per cycle, which
// keeps the per-validation fingerprint pass over the whole file far below the
// cost of the validation it avoids.
func hashBytes(b []byte) uint64 {
	h := fnvOffset64
	for i := 0; i < len(b); i++ {
		h ^= uint64(b[i])
		h *= fnvPrime64
	}
	return h
}

// forEachEntityReference calls fn for every "#<number>" ordinal appearing in
// text, stopping as soon as fn returns false.
//
// Scanning the raw text instead of the AST keeps this a single allocation-free
// pass over the document, and it cannot miss a reference. The only false
// positives are "#<number>" sequences inside string literals, which merely add
// entities to the re-validated set.
func forEachEntityReference(text []byte, fn func(ordinal int) bool) {
	for i := 0; i < len(text); i++ {
		if text[i] != '#' {
			continue
		}
		ordinal := 0
		digits := 0
		j := i + 1
		for j < len(text) && text[j] >= '0' && text[j] <= '9' {
			ordinal = ordinal*10 + int(text[j]-'0')
			digits++
			j++
			if digits >= 9 {
				break
			}
		}
		if digits == 0 {
			continue
		}
		if !fn(ordinal) {
			return
		}
		i = j - 1
	}
}

// referencesAny reports whether text points at any ordinal in targets.
func referencesAny(text []byte, targets map[int]bool) bool {
	if len(targets) == 0 {
		return false
	}
	hit := false
	forEachEntityReference(text, func(ordinal int) bool {
		if targets[ordinal] {
			hit = true
			return false
		}
		return true
	})
	return hit
}

// computeAffected determines which entity instances a change can be observed
// from. ok is false when the state cannot be used at all.
func computeAffected(units []entityUnit, content []byte, state *documentValidationState) (recompute []bool, count int, ok bool) {
	recompute = make([]bool, len(units))
	if len(units) == 0 {
		return recompute, 0, true
	}
	index := make(map[int]int, len(units))
	for i, u := range units {
		if _, duplicate := index[u.ordinal]; duplicate {
			return nil, 0, false
		}
		index[u.ordinal] = i
	}
	// An instance that disappeared can have moved any inverse set it used to
	// contribute to, and its text is no longer available to find out where.
	// Recomputing that neighbourhood is a guess, so validate everything.
	if len(state.entities) != len(index) {
		for ordinal := range state.entities {
			if _, present := index[ordinal]; !present {
				return nil, 0, false
			}
		}
	}

	var changedIdx []int
	for i, u := range units {
		previous, seen := state.entities[u.ordinal]
		if !seen || previous.fingerprint != u.hash {
			changedIdx = append(changedIdx, i)
			recompute[i] = true
		}
	}
	if len(changedIdx) == 0 {
		return recompute, 0, true
	}

	// Related = changed instances plus everything a changed instance points
	// at. Anything in that set moved some inverse set.
	related := make(map[int]bool, len(changedIdx)+8)
	for _, i := range changedIdx {
		related[units[i].ordinal] = true
	}
	for _, i := range changedIdx {
		forEachEntityReference(units[i].text(content), func(ordinal int) bool {
			related[ordinal] = true
			return true
		})
	}
	// Every instance that points at something related has to be re-validated:
	// it may read an inverse attribute of an entity that only it points at.
	for i, u := range units {
		if recompute[i] {
			continue
		}
		if related[u.ordinal] || referencesAny(u.text(content), related) {
			recompute[i] = true
		}
	}
	for _, mark := range recompute {
		if mark {
			count++
		}
	}
	return recompute, count, true
}

// validateEntityUnits validates the entities at the given indices and returns
// the diagnostics produced per entity unit. ok is false when the run was
// interrupted, in which case the result must not be published.
func validateEntityUnits(
	ctx context.Context,
	units []entityUnit,
	indices []int,
	level, source string,
) (out [][]*core.Diagnostic, ok bool) {
	out = make([][]*core.Diagnostic, len(units))
	if len(indices) == 0 {
		return out, true
	}
	if len(indices) < parallelValidateThreshold || runtime.GOMAXPROCS(0) <= 1 {
		var sink *[]*core.Diagnostic
		accept := func(d *core.Diagnostic) {
			if d.Source == "" {
				d.Source = source
			}
			*sink = append(*sink, d)
		}
		for _, i := range indices {
			if ctx.Err() != nil {
				return nil, false
			}
			sink = &out[i]
			units[i].node.Validate(ctx, level, accept)
		}
		return out, ctx.Err() == nil
	}

	var next atomic.Int64
	var wg sync.WaitGroup
	workers := runtime.GOMAXPROCS(0)
	if workers > len(indices) {
		workers = len(indices)
	}
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workerCtx := context.WithValue(ctx, entityEvalCacheKey{}, newEntityEvalCache())
			var sink *[]*core.Diagnostic
			accept := func(d *core.Diagnostic) {
				if d.Source == "" {
					d.Source = source
				}
				*sink = append(*sink, d)
			}
			for {
				k := int(next.Add(1)) - 1
				if k >= len(indices) {
					break
				}
				if workerCtx.Err() != nil {
					break
				}
				i := indices[k]
				sink = &out[i]
				units[i].node.Validate(workerCtx, level, accept)
			}
		}()
	}
	wg.Wait()
	return out, ctx.Err() == nil
}

// buildDocumentState converts a validation run into cacheable state. ok is
// false when a diagnostic does not fit inside its own entity instance, which
// means the relative encoding below would not be sound.
func buildDocumentState(
	header uint64,
	settings ValidationSettings,
	level string,
	lexErrors, parseErrors int,
	units []entityUnit,
	diags [][]*core.Diagnostic,
) (*documentValidationState, bool) {
	state := &documentValidationState{
		header:      header,
		settings:    settings,
		level:       level,
		lexErrors:   lexErrors,
		parseErrors: parseErrors,
		entities:    make(map[int]cachedEntity, len(units)),
	}
	for i, u := range units {
		entry := cachedEntity{
			fingerprint: u.hash,
		}
		if len(diags[i]) > 0 {
			entry.diags = make([]cachedDiagnostic, 0, len(diags[i]))
			for _, d := range diags[i] {
				relative, fits := relativeDiagnostic(d, u.segment.Range)
				if !fits {
					return nil, false
				}
				entry.diags = append(entry.diags, relative)
			}
		}
		if _, duplicate := state.entities[u.ordinal]; duplicate {
			return nil, false
		}
		state.entities[u.ordinal] = entry
	}
	return state, true
}

// relativeDiagnostic re-bases a diagnostic onto the start of its entity.
func relativeDiagnostic(d *core.Diagnostic, entity core.TextRange) (cachedDiagnostic, bool) {
	if d == nil {
		return cachedDiagnostic{}, false
	}
	startLine := int32(d.Range.Start.Line) - int32(entity.Start.Line)
	endLine := int32(d.Range.End.Line) - int32(entity.Start.Line)
	if startLine < 0 || endLine < 0 {
		return cachedDiagnostic{}, false
	}
	maxLine := int32(entity.End.Line) - int32(entity.Start.Line)
	if endLine > maxLine || (endLine == maxLine && int32(d.Range.End.Column) > int32(entity.End.Column)) {
		return cachedDiagnostic{}, false
	}
	startCol := int32(d.Range.Start.Column)
	endCol := int32(d.Range.End.Column)
	if startLine == 0 {
		if startCol < int32(entity.Start.Column) {
			return cachedDiagnostic{}, false
		}
		startCol -= int32(entity.Start.Column)
	}
	if endLine == 0 {
		if endCol < int32(entity.Start.Column) {
			return cachedDiagnostic{}, false
		}
		endCol -= int32(entity.Start.Column)
	}
	return cachedDiagnostic{
		startLine: startLine,
		startCol:  startCol,
		endLine:   endLine,
		endCol:    endCol,
		severity:  d.Severity,
		message:   d.Message,
		source:    d.Source,
		code:      d.Code,
		tags:      d.Tags,
		codeDesc:  d.CodeDescription,
		data:      d.Data,
	}, true
}

// hashHeaderSection fingerprints the schema header, so a FILE_SCHEMA edit
// invalidates every cached diagnostic.
func hashHeaderSection(doc *core.Document, content []byte) uint64 {
	if doc == nil || doc.Root == nil {
		return 0
	}
	file, ok := doc.Root.(EXCHANGE_FILE)
	if !ok {
		for node := range core.AllNodes(doc.Root) {
			if found, isFile := node.(EXCHANGE_FILE); isFile {
				file = found
				break
			}
		}
	}
	if file == nil {
		return 0
	}
	header := file.HeaderSection()
	if header == nil {
		return 0
	}
	segment := header.Segment()
	if segment == nil {
		return 0
	}
	return hashBytes(segmentBytes(content, segment))
}

// incrementalReuseRuns and incrementalFullRuns count how often validation
// reused cached entity diagnostics versus validating everything. They are
// read by tests to prove that a scenario really took the incremental path
// instead of silently falling back to a full walk.
var (
	incrementalReuseRuns atomic.Int64
	incrementalFullRuns  atomic.Int64
)

// validateEntities runs entity validation, reusing cached diagnostics for the
// instances a change cannot reach.
func (v *IfcDocumentValidator) validateEntities(
	ctx context.Context,
	doc *core.Document,
	level, source string,
	settings ValidationSettings,
) []*core.Diagnostic {
	units, content, scannable := collectEntityUnits(doc)
	if !scannable || len(units) < incrementalEntityThreshold {
		incrementalFullRuns.Add(1)
		return parallelValidate(ctx, doc, level, source)
	}
	uri := doc.URI.StringUnencoded()
	header := hashHeaderSection(doc, content)
	state := v.cache.get(uri)

	var recompute []bool
	reusable := false
	if state.compatible(header, settings, level, len(doc.LexerErrors), len(doc.ParserErrors)) {
		if marks, count, ok := computeAffected(units, content, state); ok {
			if count == 0 {
				// Nothing the cache tracks moved at all.
				incrementalReuseRuns.Add(1)
				return cachedDiagnostics(units, state, source)
			}
			if count*incrementalChangedBudget <= len(units) {
				recompute, reusable = marks, true
			}
		}
	}

	indices := make([]int, 0, len(units))
	if reusable {
		for i, mark := range recompute {
			if mark {
				indices = append(indices, i)
			}
		}
	} else {
		for i := range units {
			indices = append(indices, i)
		}
	}
	if reusable {
		incrementalReuseRuns.Add(1)
	} else {
		incrementalFullRuns.Add(1)
	}
	diags, ok := validateEntityUnits(ctx, units, indices, level, source)
	if !ok {
		// Interrupted by a newer edit: publish nothing, the builder discards
		// the result anyway once it sees the cancelled context.
		return nil
	}

	if ctx.Err() == nil {
		if fresh, buildable := buildDocumentState(
			header, settings, level,
			len(doc.LexerErrors), len(doc.ParserErrors),
			units, diags,
		); buildable {
			v.cache.put(uri, fresh)
		}
	}

	diagnostics := make([]*core.Diagnostic, 0, 64)
	if reusable {
		for i, u := range units {
			if recompute[i] {
				diagnostics = append(diagnostics, diags[i]...)
				continue
			}
			entry := state.entities[u.ordinal]
			for _, cached := range entry.diags {
				diagnostics = append(diagnostics, cached.diagnostic(u.segment.Range.Start))
			}
		}
	} else {
		for i := range units {
			diagnostics = append(diagnostics, diags[i]...)
		}
	}
	sortDiagnostics(diagnostics)
	return diagnostics
}

// cachedDiagnostics re-anchors every cached diagnostic of an unchanged
// document.
func cachedDiagnostics(units []entityUnit, state *documentValidationState, source string) []*core.Diagnostic {
	diagnostics := make([]*core.Diagnostic, 0, 64)
	for _, u := range units {
		entry := state.entities[u.ordinal]
		for _, cached := range entry.diags {
			diagnostics = append(diagnostics, cached.diagnostic(u.segment.Range.Start))
		}
	}
	sortDiagnostics(diagnostics)
	return diagnostics
}
