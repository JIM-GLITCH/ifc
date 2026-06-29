package ifc

import (
	"context"
	"strings"
)

const (
	wvKindEntityInst wvKind = 100
	wvKindEntityList wvKind = 101
)

func wvFromEntityInstance(inst SIMPLE_ENTITY_INSTANCE) WhereValue {
	if inst == nil {
		return wvUnknownVal()
	}
	return WhereValue{kind: wvKindEntityInst, entity: inst}
}

func wvFromEntityInstances(items []SIMPLE_ENTITY_INSTANCE) WhereValue {
	if len(items) == 0 {
		return WhereValue{kind: wvKindEntityList, entities: []SIMPLE_ENTITY_INSTANCE{}}
	}
	copyItems := append([]SIMPLE_ENTITY_INSTANCE(nil), items...)
	return WhereValue{kind: wvKindEntityList, entities: copyItems}
}

func (v WhereValue) entityInstance() SIMPLE_ENTITY_INSTANCE {
	if v.kind == wvKindEntityInst {
		return v.entity
	}
	return nil
}

type EntityWhereEval struct {
	ctx      context.Context
	stepName string
	self     SIMPLE_ENTITY_INSTANCE
	params   []PARAMETER
	pop      *Population
}

func NewEntityWhereEval(
	ctx context.Context,
	stepName string,
	self SIMPLE_ENTITY_INSTANCE,
	params []PARAMETER,
	pop *Population,
) *EntityWhereEval {
	return &EntityWhereEval{
		ctx:      ctx,
		stepName: strings.ToUpper(strings.TrimSpace(stepName)),
		self:     self,
		params:   params,
		pop:      pop,
	}
}

func (ec *EntityWhereEval) Self() WhereValue {
	return wvFromEntityInstance(ec.self)
}

func (ec *EntityWhereEval) Inverse(name string) WhereValue {
	if ec.pop == nil || ec.self == nil {
		return wvFromEntityInstances(nil)
	}
	items := ec.pop.InverseRelations(ec.self.Name(), name)
	return wvFromEntityInstances(items)
}

func (ec *EntityWhereEval) parameterToWhereValue(param PARAMETER) (WhereValue, bool) {
	return parameterToWhereValueCtx(ec.ctx, ec.pop, param)
}

func (ec *EntityWhereEval) Attr(name string) WhereValue {
	idx := attrIndex(ec.stepName, name)
	if idx < 0 {
		if def, ok := LookupType(name); ok && def.ValueKind == ParamEnum {
			return wvFromEnumType(name)
		}
		return wvUnknownVal()
	}
	if idx >= len(ec.params) {
		return wvUnknownVal()
	}
	value, ok := ec.parameterToWhereValue(ec.params[idx])
	if !ok {
		return wvUnknownVal()
	}
	return value
}

func (ec *EntityWhereEval) Derived(name string) WhereValue {
	if ec.self == nil {
		return wvUnknownVal()
	}
	return evalDerivedAttr(ec, ec.self, name)
}

func (ec *EntityWhereEval) QualEntity(v WhereValue, entityName string) WhereValue {
	if ec.self != nil && entityIsOrSubtypes(ec.stepName, entityName) {
		return wvFromEntityInstance(ec.self)
	}
	if inst := v.entityInstance(); inst != nil {
		if step, ok := instanceStepName(inst); ok && entityIsOrSubtypes(step, entityName) {
			return v
		}
	}
	return wvUnknownVal()
}

func (ec *EntityWhereEval) QualIndex(v WhereValue, index int) WhereValue {
	if index < 1 {
		return wvUnknownVal()
	}
	if v.kind == wvKindEntityList {
		if index > len(v.entities) {
			return wvUnknownVal()
		}
		return wvFromEntityInstance(v.entities[index-1])
	}
	if v.kind == wvKindList {
		if index > len(v.list) {
			return wvUnknownVal()
		}
		return v.list[index-1]
	}
	return wvUnknownVal()
}

func (ec *EntityWhereEval) QualAttr(v WhereValue, attrName string, index int) WhereValue {
	if attrName == "" && index > 0 {
		return ec.QualIndex(v, index)
	}
	if v.kind == wvKindEnumType {
		return wvFromEnumStepLiteral(attrName)
	}
	inst := v.entityInstance()
	if inst == nil {
		return wvUnknownVal()
	}
	step, ok := instanceStepName(inst)
	if !ok {
		return wvUnknownVal()
	}
	idx := attrIndex(step, attrName)
	if idx < 0 {
		if entityHasDerivedAttr(step, attrName) {
			return evalDerivedAttr(ec, inst, attrName)
		}
		if entityHasInverse(step, attrName) {
			return ec.inverseOn(inst, attrName)
		}
		return wvUnknownVal()
	}
	params := instanceParams(inst)
	if idx >= len(params) {
		return wvUnknownVal()
	}
	value, ok := ec.parameterToWhereValue(params[idx])
	if !ok {
		return wvUnknownVal()
	}
	return value
}

func (ec *EntityWhereEval) CallSchema(name string, args []WhereValue) WhereValue {
	return dispatchSchemaFunction(ec, name, args)
}

func parseUsedInRole(rolePath string) (relationStep, forwardAttr string, ok bool) {
	parts := strings.Split(strings.TrimSpace(rolePath), ".")
	if len(parts) < 2 {
		return "", "", false
	}
	return strings.ToUpper(parts[len(parts)-2]), parts[len(parts)-1], true
}

func inverseNameForUsedInRole(ownerStep, relationStep, forwardAttr string) string {
	seen := map[string]bool{}
	var result string
	var walk func(string)
	walk = func(step string) {
		step = strings.ToUpper(strings.TrimSpace(step))
		if result != "" || seen[step] {
			return
		}
		seen[step] = true
		for _, inv := range EntityInverses[step] {
			if inv.RelationStepName == relationStep && strings.EqualFold(inv.ForwardAttrName, forwardAttr) {
				result = inv.Name
				return
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
	walk(ownerStep)
	return result
}

func (ec *EntityWhereEval) UsedIn(entity WhereValue, rolePath string) WhereValue {
	inst := entity.entityInstance()
	if inst == nil {
		return wvFromEntityInstances(nil)
	}
	step, ok := instanceStepName(inst)
	if !ok {
		return wvFromEntityInstances(nil)
	}
	relationStep, forwardAttr, ok := parseUsedInRole(rolePath)
	if !ok {
		return wvFromEntityInstances(nil)
	}
	inverseName := inverseNameForUsedInRole(step, relationStep, forwardAttr)
	if inverseName == "" {
		return wvFromEntityInstances(nil)
	}
	sub := ec.withInstance(inst)
	return sub.Inverse(inverseName)
}

func expressNameToStep(name string) string {
	return strings.ToUpper(strings.TrimSpace(name))
}

func instanceStepName(inst SIMPLE_ENTITY_INSTANCE) (string, bool) {
	if inst == nil {
		return "", false
	}
	record := inst.Record()
	if record == nil {
		return "", false
	}
	step := strings.ToUpper(strings.TrimSpace(record.Keyword()))
	if step == "" {
		return "", false
	}
	return step, true
}

func instanceParams(inst SIMPLE_ENTITY_INSTANCE) []PARAMETER {
	if inst == nil {
		return nil
	}
	record := inst.Record()
	if record == nil {
		return nil
	}
	paramList := record.Parameters()
	if paramList == nil {
		return nil
	}
	return paramList.Parameters()
}

func mergedEntityWhereRules(stepName string) []EntityWhereRuleSpec {
	def, ok := LookupEntity(stepName)
	if !ok {
		return nil
	}
	seen := map[string]bool{}
	var rules []EntityWhereRuleSpec
	collect := func(name string) {
		if seen[name] {
			return
		}
		seen[name] = true
		if entityRules, ok := EntityWhereRules[name]; ok {
			rules = append(rules, entityRules...)
		}
	}
	collect(stepName)
	for _, super := range def.SuperTypes {
		collect(strings.ToUpper(super))
	}
	return rules
}

func CheckEntityWhereRules(
	ctx context.Context,
	stepName string,
	self SIMPLE_ENTITY_INSTANCE,
	params []PARAMETER,
	pop *Population,
) (ok bool, detail string) {
	rules := mergedEntityWhereRules(stepName)
	if len(rules) == 0 {
		return true, ""
	}
	ec := NewEntityWhereEval(ctx, stepName, self, params, pop)
	for _, rule := range rules {
		if rule.Check == nil {
			continue
		}
		if pass, msg := rule.Check(ec); !pass {
			return false, msg
		}
	}
	return true, ""
}
