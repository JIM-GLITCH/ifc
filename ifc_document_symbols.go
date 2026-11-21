package ifc

import (
	"context"
	"sort"
	"strings"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/server"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
	"typefox.dev/lsp"
)

type IfcDocumentSymbolProvider struct {
	sc *service.Container
}

func NewIfcDocumentSymbolProvider(sc *service.Container) server.DocumentSymbolProvider {
	return &IfcDocumentSymbolProvider{sc: sc}
}

func (p *IfcDocumentSymbolProvider) HandleDocumentSymbolRequest(ctx context.Context, params *lsp.DocumentSymbolParams) ([]lsp.DocumentSymbol, error) {
	documentManager := service.MustGet[workspace.DocumentManager](p.sc)
	doc := documentManager.Get(core.ParseURI(string(params.TextDocument.URI)))
	if doc == nil || doc.Root == nil {
		return nil, nil
	}

	pop := populationForHover(ctx, doc)
	if pop == nil {
		return nil, nil
	}
	return buildIfcDocumentSymbols(ctx, pop), nil
}

type ifcSymbolNode struct {
	rec      *InstanceRecord
	symbol   lsp.DocumentSymbol
	children []*ifcSymbolNode
}

func buildIfcDocumentSymbols(ctx context.Context, pop *Population) []lsp.DocumentSymbol {
	if pop == nil {
		return nil
	}
	pop.ensureOutlineParents(ctx)
	var ordered []*InstanceRecord
	for _, rec := range pop.byName {
		if rec != nil && rec.Instance != nil {
			ordered = append(ordered, rec)
		}
	}
	if len(ordered) == 0 {
		return nil
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ifcInstanceStart(ordered[i].Instance) < ifcInstanceStart(ordered[j].Instance)
	})

	nodes := make(map[string]*ifcSymbolNode, len(ordered))
	var outline []*ifcSymbolNode
	for _, rec := range ordered {
		if !ifcOutlineStep(rec.StepName) {
			continue
		}
		node := &ifcSymbolNode{rec: rec}
		nodes[rec.Instance.Name()] = node
		outline = append(outline, node)
	}
	if len(outline) == 0 {
		return nil
	}

	var roots []*ifcSymbolNode
	for _, node := range outline {
		name := node.rec.Instance.Name()
		step := strings.ToUpper(strings.TrimSpace(node.rec.StepName))
		parent := ""
		if isSpatialContainerStep(step) {
			parent = pop.outlineAggParent[name]
			if parent == "" {
				parent = pop.outlineContainedParent[name]
			}
			if parent == "" {
				parent = pop.outlineHostParent[name]
			}
		} else {
			parent = pop.outlineContainedParent[name]
			if parent == "" {
				parent = pop.outlineAggParent[name]
			}
			if parent == "" {
				parent = pop.outlineHostParent[name]
			}
		}
		if parentNode, ok := nodes[parent]; ok && parentNode != node {
			parentNode.children = append(parentNode.children, node)
			continue
		}
		roots = append(roots, node)
	}

	seen := map[*ifcSymbolNode]bool{}
	var symbols []lsp.DocumentSymbol
	for _, node := range roots {
		symbols = append(symbols, finalizeIfcSymbolNode(node, seen)...)
	}
	for _, node := range outline {
		if !seen[node] {
			symbols = append(symbols, finalizeIfcSymbolNode(node, seen)...)
		}
	}
	return symbols
}

func finalizeIfcSymbolNode(node *ifcSymbolNode, seen map[*ifcSymbolNode]bool) []lsp.DocumentSymbol {
	if node == nil || seen[node] {
		return nil
	}
	seen[node] = true
	node.symbol = ifcDocumentSymbolFor(node.rec)
	for _, child := range node.children {
		node.symbol.Children = append(node.symbol.Children, finalizeIfcSymbolNode(child, seen)...)
	}
	return []lsp.DocumentSymbol{node.symbol}
}

func ifcDocumentSymbolFor(rec *InstanceRecord) lsp.DocumentSymbol {
	inst := rec.Instance
	r := ifcInstanceRange(inst)
	selection := r
	if token := inst.NameToken(); token != nil && token.Segment() != nil {
		selection = token.Segment().Range.LspRange()
	}
	return lsp.DocumentSymbol{
		Name:           strings.TrimSpace(rec.StepName + " " + inst.Name()),
		Detail:         ifcSymbolName(rec),
		Kind:           ifcSymbolKind(rec.StepName),
		Range:          r,
		SelectionRange: selection,
	}
}

func ifcSymbolName(rec *InstanceRecord) string {
	if rec == nil {
		return ""
	}
	for _, attr := range []string{"LongName", "Name", "ObjectType", "Tag", "PredefinedType"} {
		if value := ifcStringAttr(rec, attr); value != "" {
			return value
		}
	}
	return ""
}

func ifcStringAttr(rec *InstanceRecord, attrName string) string {
	if rec == nil {
		return ""
	}
	index := attrIndex(rec.StepName, attrName)
	if index < 0 || index >= len(rec.Params) {
		return ""
	}
	switch param := rec.Params[index].(type) {
	case STRING_LITERAL:
		return unquoteIFCString(param.Value())
	case ENUMERATION_LITERAL:
		return strings.Trim(param.Value(), ".")
	case TYPED_PARAMETER:
		if inner := param.Parameter(); inner != nil {
			switch str := inner.(type) {
			case STRING_LITERAL:
				return unquoteIFCString(str.Value())
			case ENUMERATION_LITERAL:
				return strings.Trim(str.Value(), ".")
			}
		}
	}
	return ""
}

func ifcSymbolKind(stepName string) lsp.SymbolKind {
	stepName = strings.ToUpper(strings.TrimSpace(stepName))
	if isSpatialContainerStep(stepName) {
		return lsp.Namespace
	}
	if entityIsOrSubtypes(stepName, "IfcElement") || entityIsOrSubtypes(stepName, "IfcProduct") {
		return lsp.Object
	}
	return lsp.Field
}

func ifcOutlineStep(stepName string) bool {
	stepName = strings.ToUpper(strings.TrimSpace(stepName))
	return stepName == "IFCPROJECT" || isSpatialContainerStep(stepName) || isSpatialElementStep(stepName)
}

func isSpatialContainerStep(stepName string) bool {
	stepName = strings.ToUpper(strings.TrimSpace(stepName))
	switch stepName {
	case "IFCPROJECT", "IFCSITE", "IFCBUILDING", "IFCBUILDINGSTOREY", "IFCSPACE", "IFCFACILITY", "IFCBRIDGE", "IFCROAD", "IFCRAILWAY", "IFCMARINEFACILITY":
		return true
	}
	return entityIsOrSubtypes(stepName, "IfcSpatialStructureElement")
}

func isSpatialElementStep(stepName string) bool {
	stepName = strings.ToUpper(strings.TrimSpace(stepName))
	return entityIsOrSubtypes(stepName, "IfcElement") || entityIsOrSubtypes(stepName, "IfcSpatialElement")
}

func ifcInstanceSegment(inst SIMPLE_ENTITY_INSTANCE) *core.TextSegment {
	if inst == nil {
		return nil
	}
	if segment := inst.Segment(); segment != nil {
		return segment
	}
	if token := inst.NameToken(); token != nil {
		return token.Segment()
	}
	if record := inst.Record(); record != nil {
		return record.Segment()
	}
	return nil
}

func ifcInstanceRange(inst SIMPLE_ENTITY_INSTANCE) lsp.Range {
	if segment := ifcInstanceSegment(inst); segment != nil {
		return segment.Range.LspRange()
	}
	return lsp.Range{}
}

func ifcInstanceStart(inst SIMPLE_ENTITY_INSTANCE) int32 {
	if segment := ifcInstanceSegment(inst); segment != nil {
		return int32(segment.Indices.Start)
	}
	return int32(^uint32(0) >> 1)
}
