package ifc

import (
	"context"
	"strings"

	core "typefox.dev/fastbelt"
)

var entityAttrNames map[string][]string

type InstanceRecord struct {
	Instance SIMPLE_ENTITY_INSTANCE
	StepName string
	Params   []PARAMETER
}

type Population struct {
	byName       map[string]*InstanceRecord
	inverseLinks map[string]map[string][]SIMPLE_ENTITY_INSTANCE
}

func BuildPopulation(ctx context.Context, doc *core.Document) *Population {
	pop := &Population{
		byName:       make(map[string]*InstanceRecord),
		inverseLinks: make(map[string]map[string][]SIMPLE_ENTITY_INSTANCE),
	}
	if doc == nil || doc.Root == nil {
		return pop
	}
	file := findExchangeFile(doc.Root)
	if file == nil {
		return pop
	}
	for _, section := range file.DataSections() {
		if section == nil {
			continue
		}
		for _, inst := range section.EntityInstances() {
			simple := asSimpleInstance(inst)
			if simple == nil {
				continue
			}
			record := instanceRecordFrom(simple)
			if record == nil {
				continue
			}
			pop.byName[simple.Name()] = record
		}
	}
	pop.buildInverseLinks(ctx)
	return pop
}

func findExchangeFile(root core.AstNode) EXCHANGE_FILE {
	if root == nil {
		return nil
	}
	if file, ok := root.(StepFile); ok {
		return file.File()
	}
	if file, ok := root.(EXCHANGE_FILE); ok {
		return file
	}
	for node := range core.AllNodes(root) {
		if file, ok := node.(EXCHANGE_FILE); ok {
			return file
		}
	}
	return nil
}

func asSimpleInstance(inst ENTITY_INSTANCE) SIMPLE_ENTITY_INSTANCE {
	if inst == nil {
		return nil
	}
	if simple, ok := inst.(SIMPLE_ENTITY_INSTANCE); ok {
		return simple
	}
	if complex, ok := inst.(COMPLEX_ENTITY_INSTANCE); ok {
		record := complex.Record()
		if record == nil || len(record.Records()) == 0 {
			return nil
		}
		first := record.Records()[0]
		if first == nil {
			return nil
		}
		return &simpleFromComplex{complex: complex, record: first}
	}
	return nil
}

type simpleFromComplex struct {
	core.AstNodeBase
	complex COMPLEX_ENTITY_INSTANCE
	record  SIMPLE_RECORD
}

func (s *simpleFromComplex) IsENTITY_INSTANCE()        {}
func (s *simpleFromComplex) IsSIMPLE_ENTITY_INSTANCE() {}
func (s *simpleFromComplex) Name() string {
	if s.complex != nil {
		return s.complex.Name()
	}
	return ""
}
func (s *simpleFromComplex) NameToken() *core.Token {
	if s.complex != nil {
		return s.complex.NameToken()
	}
	return nil
}
func (s *simpleFromComplex) SetName(value *core.Token) {}
func (s *simpleFromComplex) Record() SIMPLE_RECORD     { return s.record }
func (s *simpleFromComplex) SetRecord(value SIMPLE_RECORD) {
	s.record = value
}
func (s *simpleFromComplex) ForEachNode(fn func(core.AstNode)) {
	if s.record != nil {
		fn(s.record)
	}
}
func (s *simpleFromComplex) ForEachReference(fn func(core.UntypedReference)) {}

func instanceRecordFrom(simple SIMPLE_ENTITY_INSTANCE) *InstanceRecord {
	record := simple.Record()
	if record == nil {
		return nil
	}
	stepName := strings.ToUpper(strings.TrimSpace(record.Keyword()))
	var params []PARAMETER
	if paramList := record.Parameters(); paramList != nil {
		params = paramList.Parameters()
	}
	return &InstanceRecord{
		Instance: simple,
		StepName: stepName,
		Params:   params,
	}
}

func mergedForwardInverses(stepName string) map[int][]ForwardInverseTarget {
	stepName = strings.ToUpper(strings.TrimSpace(stepName))
	seen := map[string]bool{}
	result := make(map[int][]ForwardInverseTarget)
	var walk func(string)
	walk = func(step string) {
		step = strings.ToUpper(strings.TrimSpace(step))
		if step == "" || seen[step] {
			return
		}
		seen[step] = true
		if forwardMap := ForwardInverses[step]; forwardMap != nil {
			for attrIndex, targets := range forwardMap {
				if _, ok := result[attrIndex]; !ok {
					result[attrIndex] = targets
				}
			}
		}
		def, ok := LookupEntity(step)
		if !ok {
			return
		}
		for _, super := range def.SuperTypes {
			walk(super)
		}
	}
	walk(stepName)
	if len(result) == 0 {
		return nil
	}
	return result
}

func (pop *Population) buildInverseLinks(ctx context.Context) {
	for _, rec := range pop.byName {
		forwardMap := mergedForwardInverses(rec.StepName)
		if forwardMap == nil {
			continue
		}
		for attrIndex, targets := range forwardMap {
			if attrIndex < 0 || attrIndex >= len(rec.Params) {
				continue
			}
			targetNames := entityRefTargetNames(ctx, rec.Params[attrIndex])
			for _, targetName := range targetNames {
				for _, target := range targets {
					pop.addInverseLink(targetName, target.InverseName, rec.Instance)
				}
			}
		}
	}
}

func (pop *Population) addInverseLink(targetName, inverseName string, source SIMPLE_ENTITY_INSTANCE) {
	if targetName == "" || inverseName == "" || source == nil {
		return
	}
	byInverse, ok := pop.inverseLinks[targetName]
	if !ok {
		byInverse = make(map[string][]SIMPLE_ENTITY_INSTANCE)
		pop.inverseLinks[targetName] = byInverse
	}
	sourceName := source.Name()
	for _, existing := range byInverse[inverseName] {
		if existing != nil && existing.Name() == sourceName {
			return
		}
	}
	byInverse[inverseName] = append(byInverse[inverseName], source)
}

func (pop *Population) InverseRelations(targetName, inverseName string) []SIMPLE_ENTITY_INSTANCE {
	if pop == nil {
		return nil
	}
	byInverse := pop.inverseLinks[targetName]
	if byInverse == nil {
		return nil
	}
	return byInverse[inverseName]
}

func (pop *Population) Lookup(name string) (*InstanceRecord, bool) {
	if pop == nil {
		return nil, false
	}
	rec, ok := pop.byName[name]
	return rec, ok
}

func entityRefTargetNames(ctx context.Context, param PARAMETER) []string {
	if param == nil {
		return nil
	}
	switch p := param.(type) {
	case ENTITY_REFERENCE:
		if name := entityRefTargetName(ctx, p); name != "" {
			return []string{name}
		}
	case LIST:
		var names []string
		for _, el := range p.Parameters() {
			names = append(names, entityRefTargetNames(ctx, el)...)
		}
		return names
	case TYPED_PARAMETER:
		return entityRefTargetNames(ctx, p.Parameter())
	}
	return nil
}

func entityRefTargetName(ctx context.Context, ref ENTITY_REFERENCE) string {
	value := ref.Value()
	if value == nil {
		return ""
	}
	if target := value.Ref(ctx); target != nil {
		if simple, ok := target.(SIMPLE_ENTITY_INSTANCE); ok {
			return simple.Name()
		}
	}
	return strings.TrimSpace(value.Text())
}

func attrIndex(stepName, attrName string) int {
	names := entityAttrNames[stepName]
	lower := strings.ToLower(attrName)
	for i, name := range names {
		if strings.ToLower(name) == lower {
			return i
		}
	}
	return -1
}

func entityAttrNameAt(stepName string, index int) string {
	if stepName == "" || index < 0 {
		return ""
	}
	names := entityAttrNames[stepName]
	if index >= len(names) {
		return ""
	}
	return names[index]
}

type populationContextKey struct{}

func WithPopulation(ctx context.Context, pop *Population) context.Context {
	if pop == nil {
		return ctx
	}
	return context.WithValue(ctx, populationContextKey{}, pop)
}

func PopulationFromContext(ctx context.Context) *Population {
	if ctx == nil {
		return nil
	}
	pop, _ := ctx.Value(populationContextKey{}).(*Population)
	return pop
}
