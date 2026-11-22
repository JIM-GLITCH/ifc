package ifc

import (
	"bytes"
	"context"
	"sort"
	"strings"
	"sync"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/server"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
	"typefox.dev/lsp"
)

// ifcHeaderKeywords are the STEP header entity names accepted by
// HEADER_ENTITY in place of an entity instance record.
var ifcHeaderKeywords = []string{"FILE_DESCRIPTION", "FILE_NAME", "FILE_SCHEMA"}

// ifcCompletionEntityData caches the entity-name candidates, their list
// details and their documentation markdown. All are derived from immutable
// generated schema data, so they are computed once per process. The heavy
// documentation is deliberately kept out of the completion items: it is
// attached lazily by ResolveCompletionItem when the user selects an entry,
// keeping each completion response small.
type ifcCompletionEntitiesData struct {
	names   []string
	details map[string]string
	docs    map[string]string
}

var ifcCompletionEntityData = sync.OnceValue(func() ifcCompletionEntitiesData {
	data := ifcCompletionEntitiesData{
		details: make(map[string]string, len(Entities)),
		docs:    make(map[string]string, len(Entities)),
	}
	for name, def := range Entities {
		if def.Abstract {
			continue
		}
		data.names = append(data.names, name)
		params, ok := entitySignatureParameters(name)
		if !ok {
			continue
		}
		data.details[name] = entitySignatureLabel(name, params, true)
		var b strings.Builder
		b.WriteString("`" + entitySignatureLabel(name, params, true) + "`")
		if types := entityParameterTypeMarkdown(params); types != "" {
			b.WriteString("\n\n" + types)
		}
		data.docs[name] = b.String()
	}
	sort.Strings(data.names)
	return data
})

// ifcEntityCompletionDoc returns the markdown documentation for an entity
// completion item label, or "" for non-entity items.
func ifcEntityCompletionDoc(label string) string {
	return ifcCompletionEntityData().docs[label]
}

// ---------------------------------------------------------------------------
// Completion provider
//
// IfcCompletionProvider replaces the framework's ATN-driven completion
// provider. The default provider re-runs the generated completion parser over
// the whole document prefix on every request (O(document) per keystroke),
// which is too slow for large STEP files. STEP's surface syntax is simple
// enough that the completion context can be derived from the text between
// the last newline and the cursor, so this provider never touches the parser.
// ---------------------------------------------------------------------------

type IfcCompletionProvider struct {
	sc *service.Container
}

func NewIfcCompletionProvider(sc *service.Container) server.CompletionProvider {
	return &IfcCompletionProvider{sc: sc}
}

func (p *IfcCompletionProvider) HandleCompletionRequest(ctx context.Context, params *lsp.CompletionParams) (*lsp.CompletionList, error) {
	empty := &lsp.CompletionList{IsIncomplete: false, Items: []lsp.CompletionItem{}}

	documentManager, err := service.Get[workspace.DocumentManager](p.sc)
	if err != nil {
		return empty, nil
	}
	doc := documentManager.Get(core.ParseURI(string(params.TextDocument.URI)))
	if doc == nil || doc.TextDoc == nil {
		return empty, nil
	}

	content := doc.TextDoc.Content()
	offset := doc.TextDoc.OffsetAt(params.Position)
	if offset < 0 || offset > len(content) {
		return empty, nil
	}

	lineStart := bytes.LastIndexByte(content[:offset], '\n') + 1
	linePrefix := string(content[lineStart:offset])

	inData := dataSectionStarted(content, offset)
	cc := analyzeLinePrefix(linePrefix, inData)
	if cc.kind == ctxNone {
		return empty, nil
	}

	var replace *lsp.Range
	if cc.typed != "" {
		replace = &lsp.Range{
			Start: doc.TextDoc.PositionAt(lineStart + cc.typedStart),
			End:   doc.TextDoc.PositionAt(offset),
		}
	}

	items := make([]lsp.CompletionItem, 0, 16)
	switch cc.kind {
	case ctxEntityKeyword:
		items = append(items, entityKeywordItems(cc.typed, replace)...)
	case ctxHeaderKeyword:
		items = append(items, headerKeywordItems(cc.typed, replace)...)
	case ctxParams:
		items = append(items, paramItems(doc, cc.typed, cc.attr, replace)...)
	}
	return &lsp.CompletionList{IsIncomplete: false, Items: items}, nil
}

// ---------------------------------------------------------------------------
// Line analysis
// ---------------------------------------------------------------------------

type completionContextKind int

const (
	ctxNone          completionContextKind = iota
	ctxEntityKeyword                       // record keyword position: #N=| or #N=(...|,|
	ctxHeaderKeyword                       // header entity keyword position
	ctxParams                              // inside a parameter list: entity refs + $/*
)

type lineContext struct {
	kind       completionContextKind
	typed      string      // partial token text being completed ("" = insert)
	typedStart int         // byte offset of typed within the line prefix
	attr       *SchemaAttr // expected attribute for reference filtering (nil = unfiltered)
}

type parenFrameKind int

const (
	frameEntity   parenFrameKind = iota // IFCDIRECTION( - parameter list of an entity record
	frameSubsuper                       // #N=( - complex entity instance, expects record keywords
	frameList                           // ( - nested parameter list
	frameHeader                         // FILE_NAME( - header entity parameters
	frameTyped                          // TIMESTAMP( - typed parameter
)

type parenFrame struct {
	kind          parenFrameKind
	recordKeyword string
	commas        int
}

// analyzeLinePrefix derives the completion context from the line text before
// the cursor. It scans the prefix once, tracking string literals and the
// stack of unmatched '(' so that commas and keywords are attributed to the
// correct nesting level. Only the current line is inspected; constructs
// that span lines degrade gracefully to the nearest line-local context.
func analyzeLinePrefix(prefix string, inData bool) lineContext {
	var stack []parenFrame
	expectKeyword := false // saw '=' at top level, record keyword not yet typed
	inString := false

	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if inString {
			if c == '\'' {
				if i+1 < len(prefix) && prefix[i+1] == '\'' {
					i++ // escaped quote ''
				} else {
					inString = false
				}
			}
			continue
		}
		switch {
		case c == '\'':
			inString = true
		case c == ';':
			stack = stack[:0]
			expectKeyword = false
		case c == '=' && len(stack) == 0:
			expectKeyword = true
		case c == '(':
			kind, keyword := classifyOpenParen(prefix, i)
			stack = append(stack, parenFrame{kind: kind, recordKeyword: keyword})
			expectKeyword = false
		case c == ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case c == ',' && len(stack) > 0:
			stack[len(stack)-1].commas++
		}
	}
	if inString {
		return lineContext{kind: ctxNone}
	}

	typed, typedStart := typedSuffix(prefix)

	if len(stack) > 0 {
		top := stack[len(stack)-1]
		switch top.kind {
		case frameEntity:
			return lineContext{kind: ctxParams, typed: typed, typedStart: typedStart,
				attr: entityAttrAt(top.recordKeyword, top.commas)}
		case frameList:
			return lineContext{kind: ctxParams, typed: typed, typedStart: typedStart,
				attr: enclosingEntityAttr(stack)}
		case frameSubsuper:
			return lineContext{kind: ctxEntityKeyword, typed: typed, typedStart: typedStart}
		case frameTyped:
			return lineContext{kind: ctxParams, typed: typed, typedStart: typedStart}
		case frameHeader:
			return lineContext{kind: ctxNone}
		}
	}

	if expectKeyword {
		return lineContext{kind: ctxEntityKeyword, typed: typed, typedStart: typedStart}
	}

	// Stack empty: statement level.
	if strings.HasPrefix(typed, "#") {
		// A reference token on a continuation line.
		return lineContext{kind: ctxParams, typed: typed, typedStart: typedStart}
	}
	if isBareKeyword(typed) {
		// A bare uppercase word at statement level is a record keyword
		// (SUBSUPER_RECORD continuation) in the data section or a header
		// entity keyword in the header section.
		if inData {
			return lineContext{kind: ctxEntityKeyword, typed: typed, typedStart: typedStart}
		}
		return lineContext{kind: ctxHeaderKeyword, typed: typed, typedStart: typedStart}
	}
	if strings.TrimSpace(prefix) == "" || strings.HasSuffix(strings.TrimSpace(prefix), ";") {
		// Fresh statement position. In the header section this is a header
		// entity keyword; in the data section a new statement starts with
		// '#', which has no completable prefix.
		if !inData {
			return lineContext{kind: ctxHeaderKeyword, typed: typed, typedStart: typedStart}
		}
	}
	return lineContext{kind: ctxNone}
}

// classifyOpenParen determines what a '(' opens by inspecting the
// significant token before it.
func classifyOpenParen(prefix string, openIdx int) (parenFrameKind, string) {
	j := openIdx - 1
	for j >= 0 && (prefix[j] == ' ' || prefix[j] == '\t') {
		j--
	}
	if j < 0 {
		return frameList, ""
	}
	c := prefix[j]
	switch {
	case c == '=':
		// #N=( - SUBSUPER_RECORD of a complex entity instance.
		return frameSubsuper, ""
	case c == '(' || c == ',' || c == ';':
		// '(' inside a parameter list or right after a separator: a nested
		// LIST parameter (or DATA(...) parameters).
		return frameList, ""
	}
	// Otherwise the '(' follows a keyword: extract the word before it.
	end := j + 1
	for j >= 0 && isKeywordChar(prefix[j]) {
		j--
	}
	word := strings.ToUpper(prefix[j+1 : end])
	word = strings.TrimPrefix(word, "!")
	if word == "" {
		return frameList, ""
	}
	if isHeaderKeyword(word) {
		return frameHeader, word
	}
	if _, ok := LookupEntity(word); ok {
		return frameEntity, word
	}
	return frameTyped, word
}

// typedSuffix returns the completable token text at the end of prefix and
// its start offset within prefix. The suffix covers STEP token characters:
// entity names, #references, enumerations (.X.), $ and *.
func typedSuffix(prefix string) (string, int) {
	i := len(prefix)
	for i > 0 && isCompletableChar(prefix[i-1]) {
		i--
	}
	return prefix[i:], i
}

func isCompletableChar(c byte) bool {
	return c == '#' || c == '.' || c == '$' || c == '*' || c == '!' ||
		(c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c == '_'
}

func isKeywordChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c == '_' || c == '!'
}

func isBareKeyword(s string) bool {
	if s == "" {
		return false
	}
	s = strings.TrimPrefix(s, "!")
	if s == "" || s[0] < 'A' || s[0] > 'Z' {
		return false
	}
	for i := 1; i < len(s); i++ {
		c := s[i]
		if !(c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_') {
			return false
		}
	}
	return true
}

func isHeaderKeyword(word string) bool {
	for _, kw := range ifcHeaderKeywords {
		if word == kw {
			return true
		}
	}
	return false
}

// dataSectionStarted reports whether the DATA section has begun before
// offset. It looks for a standalone DATA token followed by '(' or ';'.
func dataSectionStarted(content []byte, offset int) bool {
	search := content[:offset]
	for idx := bytes.LastIndex(search, []byte("DATA")); idx >= 0; {
		end := idx + len("DATA")
		beforeOK := idx == 0 || !isKeywordChar(search[idx-1])
		after := end
		for after < len(search) && (search[after] == ' ' || search[after] == '\t' || search[after] == '\n' || search[after] == '\r') {
			after++
		}
		afterOK := after < len(search) && (search[after] == '(' || search[after] == ';')
		if beforeOK && afterOK {
			return true
		}
		search = search[:idx]
		idx = bytes.LastIndex(search, []byte("DATA"))
	}
	return false
}

// ---------------------------------------------------------------------------
// Candidate builders
// ---------------------------------------------------------------------------

func entityKeywordItems(typed string, replace *lsp.Range) []lsp.CompletionItem {
	data := ifcCompletionEntityData()
	upper := strings.ToUpper(strings.TrimPrefix(typed, "!"))
	items := make([]lsp.CompletionItem, 0, 64)
	for _, name := range data.names {
		if upper != "" && !strings.Contains(name, upper) {
			continue
		}
		item := lsp.CompletionItem{
			Label:    name,
			Kind:     lsp.ClassCompletion,
			Detail:   data.details[name],
			SortText: "0-1-entity",
		}
		applyCompletionText(&item, replace, name)
		items = append(items, item)
	}
	return items
}

func headerKeywordItems(typed string, replace *lsp.Range) []lsp.CompletionItem {
	upper := strings.ToUpper(typed)
	items := make([]lsp.CompletionItem, 0, len(ifcHeaderKeywords)+1)
	for _, keyword := range ifcHeaderKeywords {
		if upper != "" && !strings.Contains(keyword, upper) {
			continue
		}
		item := lsp.CompletionItem{
			Label:    keyword,
			Kind:     lsp.KeywordCompletion,
			Detail:   "STEP header entity",
			SortText: "0-1-keyword",
		}
		applyCompletionText(&item, replace, keyword)
		items = append(items, item)
	}
	return items
}

// paramItems returns the candidates valid inside a parameter list: entity
// instance references (#N) filtered by the expected attribute type when it
// can be determined, plus the $ (default) and * (omitted) markers.
func paramItems(doc *core.Document, typed string, attr *SchemaAttr, replace *lsp.Range) []lsp.CompletionItem {
	items := make([]lsp.CompletionItem, 0, 8)
	for _, marker := range []struct{ label, detail string }{
		{"$", "default value"},
		{"*", "omitted parameter"},
	} {
		item := lsp.CompletionItem{
			Label:    marker.label,
			Kind:     lsp.KeywordCompletion,
			Detail:   marker.detail,
			SortText: "0-2-marker",
		}
		applyCompletionText(&item, replace, marker.label)
		items = append(items, item)
	}

	upper := strings.ToUpper(typed)
	for desc := range allIfcLocalSymbols(doc) {
		if desc == nil {
			continue
		}
		name := desc.Name.String()
		if upper != "" && !strings.Contains(strings.ToUpper(name), upper) {
			continue
		}
		if attr != nil && !symbolMatchesAttr(desc.Node, *attr) {
			continue
		}
		item := lsp.CompletionItem{
			Label:    name,
			Kind:     lsp.ReferenceCompletion,
			SortText: "0-0-ref",
		}
		applyCompletionText(&item, replace, name)
		items = append(items, item)
	}
	return items
}

// applyCompletionText sets either a TextEdit (when replace is non-nil,
// covering the partially typed token) or InsertText (plain insertion at the
// cursor).
func applyCompletionText(item *lsp.CompletionItem, replace *lsp.Range, newText string) {
	if replace != nil {
		item.TextEdit = &lsp.Or_CompletionItem_textEdit{
			Value: lsp.TextEdit{Range: *replace, NewText: newText},
		}
		return
	}
	item.InsertText = newText
}

// ---------------------------------------------------------------------------
// Attribute/type filtering for reference candidates
// ---------------------------------------------------------------------------

// entityAttrAt returns the schema attribute at the given parameter index of
// the entity record, or nil when the index is out of range.
func entityAttrAt(recordKeyword string, paramIndex int) *SchemaAttr {
	def, ok := LookupEntity(strings.ToUpper(recordKeyword))
	if !ok || paramIndex < 0 || paramIndex >= len(def.Attrs) {
		return nil
	}
	return &def.Attrs[paramIndex]
}

// enclosingEntityAttr resolves the expected attribute for a position inside
// a nested LIST: the list itself is the parameter value at the enclosing
// entity record's current parameter index.
func enclosingEntityAttr(stack []parenFrame) *SchemaAttr {
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i].kind == frameEntity {
			return entityAttrAt(stack[i].recordKeyword, stack[i].commas)
		}
	}
	return nil
}

// symbolMatchesAttr reports whether the entity instance declared by node may
// validly appear at a parameter with the given attribute type. Unknown
// symbol shapes are treated permissively.
func symbolMatchesAttr(node core.AstNode, attr SchemaAttr) bool {
	stepName, ok := symbolStepName(node)
	if !ok {
		return true
	}
	expected := attrEntityTypeNames(attr)
	if expected == nil {
		return true // cannot determine a constraint; stay permissive
	}
	for _, name := range expected {
		if entityIsOrSubtypes(stepName, name) {
			return true
		}
	}
	return false
}

// attrEntityTypeNames returns the EXPRESS type names an attribute accepts
// for entity references, or nil when no entity constraint applies.
// An empty non-nil result means the attribute cannot hold an entity at all.
func attrEntityTypeNames(attr SchemaAttr) []string {
	switch attr.Kind {
	case ParamEntity:
		if _, ok := LookupEntity(strings.ToUpper(attr.TypeName)); ok {
			return []string{attr.TypeName}
		}
		return nil
	case ParamSelect:
		typeDef, ok := LookupType(attr.TypeName)
		if !ok {
			return nil
		}
		var out []string
		for _, member := range typeDef.SelectMembers {
			if _, isEntity := LookupEntity(strings.ToUpper(member)); isEntity {
				out = append(out, member)
			}
		}
		return out
	case ParamAggregate:
		elem := aggregateElementType(attr.TypeName)
		if elem == "" {
			return nil
		}
		if _, isEntity := LookupEntity(strings.ToUpper(elem)); isEntity {
			return []string{elem}
		}
		if typeDef, ok := LookupType(elem); ok && typeDef.ValueKind == ParamSelect {
			var out []string
			for _, member := range typeDef.SelectMembers {
				if _, isEntity := LookupEntity(strings.ToUpper(member)); isEntity {
					out = append(out, member)
				}
			}
			return out
		}
		// Aggregate of a non-entity type (measures, strings, ...): no entity
		// reference is valid here.
		return []string{}
	default:
		return nil
	}
}

// aggregateElementType extracts the element type name from an aggregate type
// expression such as "SET [1:?] OF IfcLayeredItem" or "LIST [2:3] OF IfcReal".
func aggregateElementType(typeName string) string {
	idx := strings.LastIndex(typeName, " OF ")
	if idx < 0 {
		return ""
	}
	return strings.TrimSpace(typeName[idx+len(" OF "):])
}

// symbolStepName returns the STEP entity name declared by an entity instance
// AST node.
func symbolStepName(node core.AstNode) (string, bool) {
	switch n := node.(type) {
	case SIMPLE_ENTITY_INSTANCE:
		if record := n.Record(); record != nil {
			return strings.ToUpper(record.Keyword()), true
		}
	case COMPLEX_ENTITY_INSTANCE:
		if record := n.Record(); record != nil && len(record.Records()) > 0 && record.Records()[0] != nil {
			return strings.ToUpper(record.Records()[0].Keyword()), true
		}
	}
	return "", false
}
