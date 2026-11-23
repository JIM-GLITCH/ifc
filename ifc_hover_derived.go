package ifc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	core "typefox.dev/fastbelt"
	"typefox.dev/lsp"
)

type ifcPopulationEntry struct {
	root core.AstNode
	pop  *Population
}

func populationForHover(ctx context.Context, doc *core.Document) *Population {
	if doc == nil {
		return nil
	}
	if v, ok := doc.Data.Load("ifc.hover.population"); ok {
		if entry, ok := v.(*ifcPopulationEntry); ok && entry.root == doc.Root {
			return entry.pop
		}
	}
	pop := BuildPopulation(ctx, doc)
	doc.Data.Store("ifc.hover.population", &ifcPopulationEntry{root: doc.Root, pop: pop})
	return pop
}

func derivedHover(ctx context.Context, doc *core.Document, token *core.Token) *lsp.Hover {
	if token == nil || token.Type != Keyword_Asterisk {
		return nil
	}
	omitted, ok := token.Element.(OMITTED_PARAMETER)
	if !ok || omitted == nil {
		return nil
	}

	var record SIMPLE_RECORD
	var instance SIMPLE_ENTITY_INSTANCE
	for node := omitted.Container(); node != nil; node = node.Container() {
		switch n := node.(type) {
		case SIMPLE_RECORD:
			record = n
		case SIMPLE_ENTITY_INSTANCE:
			instance = n
		case COMPLEX_ENTITY_INSTANCE:
			if record != nil {
				instance = &simpleFromComplex{complex: n, record: record}
			}
		}
		if record != nil && instance != nil {
			break
		}
	}
	if record == nil || instance == nil {
		return nil
	}

	params := instanceParams(instance)
	index := -1
	for i, param := range params {
		if parameterInfoFrom(param).node == omitted {
			index = i
			break
		}
	}
	if index < 0 {
		return nil
	}

	stepName := strings.ToUpper(strings.TrimSpace(record.Keyword()))
	attrName := entityAttrNameAt(stepName, index)
	if attrName == "" || !entityHasDerivedAttr(stepName, attrName) {
		return nil
	}

	pop := populationForHover(ctx, doc)
	evalCtx := WithEntityEvalCache(WithPopulation(ctx, pop))
	ec := NewEntityWhereEval(evalCtx, stepName, instance, params, pop)
	value := ec.Derived(attrName)
	if refs := collectWhereValueEntityRefs(value, doc); len(refs) > 0 {
		link := func(inst SIMPLE_ENTITY_INSTANCE) string {
			return ifcEntityHoverCommandLink(inst, doc)
		}
		formatted := formatWhereValueForHoverMarkdown(value, link)
		if formatted == "" {
			return nil
		}
		var b strings.Builder
		b.WriteString("**" + attrName + "** = " + formatted)
		for _, inst := range refs {
			preview := strings.TrimSpace(inst.Text())
			if preview == "" {
				continue
			}
			b.WriteString("\n\n```ifc\n" + preview + "\n```")
		}
		return &lsp.Hover{
			Contents: lsp.MarkupContent{
				Kind:  lsp.Markdown,
				Value: b.String(),
			},
			Range: token.TextRange().LspRange(doc.TextDoc),
		}
	}

	formatted := formatWhereValueForHover(value)
	if formatted == "" {
		return nil
	}

	return &lsp.Hover{
		Contents: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: fmt.Sprintf("**%s** = `%s`", attrName, formatted),
		},
		Range: token.TextRange().LspRange(doc.TextDoc),
	}
}

func formatWhereValueForHover(v WhereValue) string {
	switch v.kind {
	case wvKindUnknown, wvKindAbsent:
		return ""
	case wvKindInteger:
		return strconv.FormatInt(v.i, 10)
	case wvKindReal:
		return strconv.FormatFloat(v.r, 'f', -1, 64)
	case wvKindString:
		return "'" + strings.ReplaceAll(v.s, "'", "''") + "'"
	case wvKindEnum:
		member := normalizeEnumMember(v.s)
		if member == "" {
			return ""
		}
		return "." + member + "."
	case wvKindLogical:
		switch v.log {
		case 0:
			return "FALSE"
		case 1:
			return "TRUE"
		default:
			return "?"
		}
	case wvKindEnumType:
		return v.s
	case wvKindEntityInst:
		return v.entity.Name()
	case wvKindEntityList:
		parts := make([]string, 0, len(v.entities))
		for _, inst := range v.entities {
			parts = append(parts, inst.Name())
		}
		return "(" + strings.Join(parts, ",") + ")"
	case wvKindList:
		parts := make([]string, 0, len(v.list))
		for _, item := range v.list {
			part := formatWhereValueForHover(item)
			if part == "" {
				part = "?"
			}
			parts = append(parts, part)
		}
		return "(" + strings.Join(parts, ",") + ")"
	case wvKindRecord:
		return formatWhereRecordForHover(v)
	default:
		return ""
	}
}

func formatWhereRecordForHover(v WhereValue) string {
	step := expressNameToStep(v.recordType)
	names := entityAttrNames[step]
	args := make([]string, 0, len(names))
	for _, name := range names {
		attr, ok := v.recordAttrs[name]
		if !ok {
			attr = wvUnknownVal()
		}
		part := formatWhereValueForHover(attr)
		if part == "" {
			part = "?"
		}
		args = append(args, part)
	}
	if len(args) == 0 {
		return v.recordType
	}
	return v.recordType + "(" + strings.Join(args, ",") + ")"
}

type ifcHoverRefTarget struct {
	URI   string    `json:"uri"`
	Range lsp.Range `json:"range"`
}

func ifcEntityHoverCommandLink(inst SIMPLE_ENTITY_INSTANCE, fallbackDoc *core.Document) string {
	if inst == nil {
		return ""
	}
	name := strings.TrimSpace(inst.Name())
	if name == "" {
		return ""
	}
	targetDoc := inst.Document()
	if targetDoc == nil {
		targetDoc = fallbackDoc
	}
	if targetDoc == nil || targetDoc.URI == nil {
		return ""
	}
	var r core.TextRange
	if token := inst.NameToken(); token != nil {
		if rng := token.TextRange(); rng.End > rng.Start {
			r = rng
		}
	}
	if r.End <= r.Start {
		if rng := inst.TextRange(); rng.End > rng.Start {
			r = rng
		}
	}
	if r.End <= r.Start || targetDoc.TextDoc == nil {
		return ""
	}
	args, err := json.Marshal([]ifcHoverRefTarget{{URI: string(targetDoc.URI.DocumentURI()), Range: r.LspRange(targetDoc.TextDoc)}})
	if err != nil {
		return ""
	}
	return fmt.Sprintf("[%s](command:ifc.openEntityReference?%s)", name, url.QueryEscape(string(args)))
}

func collectWhereValueEntityRefs(v WhereValue, doc *core.Document) []SIMPLE_ENTITY_INSTANCE {
	var refs []SIMPLE_ENTITY_INSTANCE
	seen := map[string]bool{}
	var add func(SIMPLE_ENTITY_INSTANCE)
	add = func(inst SIMPLE_ENTITY_INSTANCE) {
		if inst == nil {
			return
		}
		name := strings.TrimSpace(inst.Name())
		if name == "" {
			return
		}
		targetDoc := inst.Document()
		if targetDoc == nil {
			targetDoc = doc
		}
		uri := ""
		if targetDoc != nil && targetDoc.URI != nil {
			uri = string(targetDoc.URI.DocumentURI())
		}
		key := uri + "|" + name
		if seen[key] {
			return
		}
		seen[key] = true
		refs = append(refs, inst)
	}
	var walk func(WhereValue)
	walk = func(value WhereValue) {
		switch value.kind {
		case wvKindEntityInst:
			add(value.entity)
		case wvKindEntityList:
			for _, inst := range value.entities {
				add(inst)
			}
		case wvKindList:
			for _, item := range value.list {
				walk(item)
			}
		case wvKindRecord:
			step := expressNameToStep(value.recordType)
			used := map[string]bool{}
			for _, name := range entityAttrNames[step] {
				used[name] = true
				if attr, ok := value.recordAttrs[name]; ok {
					walk(attr)
				}
			}
			extra := make([]string, 0, len(value.recordAttrs))
			for name := range value.recordAttrs {
				if !used[name] {
					extra = append(extra, name)
				}
			}
			sort.Strings(extra)
			for _, name := range extra {
				walk(value.recordAttrs[name])
			}
		}
	}
	walk(v)
	return refs
}

func formatWhereValueForHoverMarkdown(v WhereValue, link func(SIMPLE_ENTITY_INSTANCE) string) string {
	switch v.kind {
	case wvKindUnknown, wvKindAbsent:
		return ""
	case wvKindEntityInst:
		if link != nil {
			if linked := link(v.entity); linked != "" {
				return linked
			}
		}
		return v.entity.Name()
	case wvKindEntityList:
		parts := make([]string, 0, len(v.entities))
		for _, inst := range v.entities {
			part := ""
			if link != nil {
				part = link(inst)
			}
			if part == "" {
				part = inst.Name()
			}
			parts = append(parts, part)
		}
		return "(" + strings.Join(parts, ",") + ")"
	case wvKindList:
		parts := make([]string, 0, len(v.list))
		for _, item := range v.list {
			part := formatWhereValueForHoverMarkdown(item, link)
			if part == "" {
				part = "?"
			}
			parts = append(parts, part)
		}
		return "(" + strings.Join(parts, ",") + ")"
	case wvKindRecord:
		return formatWhereRecordForHoverMarkdown(v, link)
	default:
		return formatWhereValueForHover(v)
	}
}

func formatWhereRecordForHoverMarkdown(v WhereValue, link func(SIMPLE_ENTITY_INSTANCE) string) string {
	step := expressNameToStep(v.recordType)
	names := entityAttrNames[step]
	args := make([]string, 0, len(names))
	for _, name := range names {
		attr, ok := v.recordAttrs[name]
		if !ok {
			attr = wvUnknownVal()
		}
		part := formatWhereValueForHoverMarkdown(attr, link)
		if part == "" {
			part = "?"
		}
		args = append(args, part)
	}
	if len(args) == 0 {
		return v.recordType
	}
	return v.recordType + "(" + strings.Join(args, ",") + ")"
}
