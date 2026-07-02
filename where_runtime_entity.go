package ifc

import (
	"context"
	"fmt"
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

func whereValueExpressType(v WhereValue) (string, bool) {
	if inst := v.entityInstance(); inst != nil {
		return instanceStepName(inst)
	}
	if v.kind == wvKindRecord && v.recordType != "" {
		return expressNameToStep(v.recordType), true
	}
	return "", false
}

func wvEntityInstanceAsRecord(inst SIMPLE_ENTITY_INSTANCE) (WhereValue, bool) {
	step, ok := instanceStepName(inst)
	if !ok {
		return WhereValue{}, false
	}
	def, ok := LookupEntity(step)
	if !ok {
		return WhereValue{}, false
	}
	names := entityAttrNames[step]
	params := instanceParams(inst)
	if len(names) == 0 || len(params) == 0 {
		return WhereValue{}, false
	}
	attrs := make(map[string]WhereValue, len(names))
	for i, name := range names {
		if i >= len(params) {
			break
		}
		value, ok := parameterToWhereValue(params[i])
		if !ok {
			return WhereValue{}, false
		}
		attrs[name] = value
	}
	return WhereValue{
		kind:        wvKindRecord,
		recordType:  def.ExpressName,
		recordAttrs: attrs,
	}, true
}

func wvNormalizeForCompare(v WhereValue) WhereValue {
	if inst := v.entityInstance(); inst != nil {
		if rec, ok := wvEntityInstanceAsRecord(inst); ok {
			return rec
		}
	}
	return v
}

func (v WhereValue) entityInstance() SIMPLE_ENTITY_INSTANCE {
	if v.kind == wvKindEntityInst {
		return v.entity
	}
	return nil
}

type EntityWhereEval struct {
	ctx          context.Context
	stepName     string
	self         SIMPLE_ENTITY_INSTANCE
	params       []PARAMETER
	pop          *Population
	evalUnknowns *[]string
	whereKind    string
	ruleLabel    string
}

func (ec *EntityWhereEval) noteEvalUnknown(message string) {
	if ec == nil || ec.evalUnknowns == nil || strings.TrimSpace(message) == "" || shouldSuppressEvalUnknownMessage(message) {
		return
	}
	message = formatWhereEvalUnknownMessage(ec.stepName, ec.ruleLabel, ec.whereKind, message)
	for _, existing := range *ec.evalUnknowns {
		if existing == message {
			return
		}
	}
	*ec.evalUnknowns = append(*ec.evalUnknowns, message)
}

func (ec *EntityWhereEval) hasEvalUnknowns() bool {
	return ec != nil && ec.evalUnknowns != nil && len(*ec.evalUnknowns) > 0
}

func formatWhereEvalUnknownMessage(stepName, ruleLabel, whereKind, message string) string {
	stepName = strings.TrimSpace(stepName)
	ruleLabel = strings.TrimSpace(ruleLabel)
	if ruleLabel == "" {
		return message
	}
	switch whereKind {
	case "global":
		if stepName != "" {
			return fmt.Sprintf("Global RULE %q %q: %s", stepName, ruleLabel, message)
		}
		return fmt.Sprintf("Global RULE %q: %s", ruleLabel, message)
	default:
		if stepName != "" {
			return fmt.Sprintf("Entity %q WHERE %q: %s", stepName, ruleLabel, message)
		}
		return fmt.Sprintf("WHERE %q: %s", ruleLabel, message)
	}
}

func (ec *EntityWhereEval) unknown(message string) WhereValue {
	if ec != nil && ec.hasEvalUnknowns() {
		return wvUnknownVal()
	}
	ec.noteEvalUnknown(message)
	return wvUnknownVal()
}

func (ec *EntityWhereEval) WvOrLazy(a, b func() WhereValue) WhereValue {
	av := a()
	if ec != nil && ec.hasEvalUnknowns() {
		return wvUnknownVal()
	}
	if wvIsDefiniteTrue(av) {
		return WhereValue{kind: wvKindLogical, log: 1}
	}
	if wvIsDefiniteFalse(av) {
		bv := b()
		if ec != nil && ec.hasEvalUnknowns() {
			return wvUnknownVal()
		}
		return wvOr(av, bv)
	}
	if ec != nil && ec.hasEvalUnknowns() {
		return wvUnknownVal()
	}
	bv := b()
	return wvOr(av, bv)
}

func (ec *EntityWhereEval) WvAndLazy(a, b func() WhereValue) WhereValue {
	av := a()
	if ec != nil && ec.hasEvalUnknowns() {
		return wvUnknownVal()
	}
	if wvIsDefiniteFalse(av) {
		return WhereValue{kind: wvKindLogical, log: 0}
	}
	if wvIsDefiniteTrue(av) {
		bv := b()
		if ec != nil && ec.hasEvalUnknowns() {
			return wvUnknownVal()
		}
		return wvAnd(av, bv)
	}
	if ec != nil && ec.hasEvalUnknowns() {
		return wvUnknownVal()
	}
	bv := b()
	return wvAnd(av, bv)
}

func (ec *EntityWhereEval) takeEvalUnknowns() []string {
	if ec == nil || ec.evalUnknowns == nil {
		return nil
	}
	out := append([]string(nil), *ec.evalUnknowns...)
	*ec.evalUnknowns = (*ec.evalUnknowns)[:0]
	return out
}

func NewEntityWhereEval(
	ctx context.Context,
	stepName string,
	self SIMPLE_ENTITY_INSTANCE,
	params []PARAMETER,
	pop *Population,
) *EntityWhereEval {
	notes := []string{}
	return &EntityWhereEval{
		ctx:          ctx,
		stepName:     strings.ToUpper(strings.TrimSpace(stepName)),
		self:         self,
		params:       params,
		pop:          pop,
		evalUnknowns: &notes,
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
	if entityHasDerivedAttr(ec.stepName, name) {
		if ec.self == nil {
			return ec.unknown(fmt.Sprintf("Cannot evaluate attribute %q: missing entity instance", name))
		}
		return evalDerivedAttr(ec, ec.self, name)
	}
	idx := attrIndex(ec.stepName, name)
	if idx < 0 {
		if def, ok := LookupType(name); ok && def.ValueKind == ParamEnum {
			return wvFromEnumType(name)
		}
		return wvUnknownVal()
	}
	if idx >= len(ec.params) {
		return ec.unknown(fmt.Sprintf("Cannot evaluate attribute %q on %s: parameter missing", name, ec.stepName))
	}
	value, ok := ec.parameterToWhereValue(ec.params[idx])
	if !ok {
		if parameterBlockedByUnresolvedEntityRef(ec.ctx, ec.pop, ec.params[idx]) {
			return wvUnknownVal()
		}
		return ec.unknown(parameterEvalFailureReason(ec.ctx, ec.pop, ec.params[idx]))
	}
	return value
}

func (ec *EntityWhereEval) Derived(name string) WhereValue {
	if ec.self == nil {
		return ec.unknown(fmt.Sprintf("Cannot evaluate DERIVE attribute %q: missing entity instance", name))
	}
	return evalDerivedAttr(ec, ec.self, name)
}

func (ec *EntityWhereEval) QualEntity(v WhereValue, entityName string) WhereValue {
	if v.kind == wvKindRecord && v.recordType != "" {
		if entityIsOrSubtypes(v.recordType, entityName) {
			return v
		}
		return wvUnknownVal()
	}
	if inst := v.entityInstance(); inst != nil {
		if step, ok := instanceStepName(inst); ok && entityIsOrSubtypes(step, entityName) {
			return v
		}
		return wvUnknownVal()
	}
	if ec.self != nil && entityIsOrSubtypes(ec.stepName, entityName) {
		return wvFromEntityInstance(ec.self)
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
	if v.kind == wvKindRecord {
		if attr, ok := v.recordAttrs[attrName]; ok {
			if index > 0 {
				return ec.QualIndex(attr, index)
			}
			return attr
		}
		if strings.EqualFold(attrName, "Dim") {
			if entityIsOrSubtypes(v.recordType, "IfcVector") {
				if orientation, ok := v.recordAttrs["Orientation"]; ok {
					return ec.QualAttr(orientation, "Dim", 0)
				}
			}
			if ratios, ok := v.recordAttrs["DirectionRatios"]; ok && ratios.kind == wvKindList {
				return wvFromInt(len(ratios.list))
			}
		}
		return wvUnknownVal()
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
	if entityHasDerivedAttr(step, attrName) {
		return evalDerivedAttr(ec, inst, attrName)
	}
	idx := attrIndex(step, attrName)
	if idx < 0 {
		if entityHasInverse(step, attrName) {
			return ec.inverseOn(inst, attrName)
		}
		return wvUnknownVal()
	}
	params := instanceParams(inst)
	if idx >= len(params) {
		return ec.unknown(fmt.Sprintf("Cannot evaluate attribute %q on %s: parameter missing", attrName, step))
	}
	value, ok := ec.parameterToWhereValue(params[idx])
	if !ok {
		if parameterBlockedByUnresolvedEntityRef(ec.ctx, ec.pop, params[idx]) {
			return wvUnknownVal()
		}
		return ec.unknown(parameterEvalFailureReason(ec.ctx, ec.pop, params[idx]))
	}
	return value
}

func (ec *EntityWhereEval) CallSchema(name string, args []WhereValue) WhereValue {
	result := dispatchSchemaFunction(ec, name, args)
	if result.kind != wvKindUnknown {
		return result
	}
	result = constructEntityRecord(name, args)
	if result.kind == wvKindUnknown {
		return ec.unknown(fmt.Sprintf("Schema function %q could not be evaluated", name))
	}
	return result
}

func constructEntityRecord(expressName string, args []WhereValue) WhereValue {
	step := expressNameToStep(expressName)
	names := entityAttrNames[step]
	if len(names) == 0 || len(args) == 0 {
		return wvUnknownVal()
	}
	return wvNewEntityRecord(expressName, names, args)
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
) (ok bool, detail string, evalUnknowns []string) {
	rules := mergedEntityWhereRules(stepName)
	if len(rules) == 0 {
		return true, "", nil
	}
	for _, rule := range rules {
		if rule.Check == nil {
			continue
		}
		ec := NewEntityWhereEval(ctx, stepName, self, params, pop)
		ec.ruleLabel = rule.Label
		pass, msg := rule.Check(ec)
		ruleUnknowns := ec.takeEvalUnknowns()
		evalUnknowns = append(evalUnknowns, ruleUnknowns...)
		if len(ruleUnknowns) > 0 {
			break
		}
		if !pass && detail == "" {
			detail = msg
		}
	}
	if detail != "" {
		return false, detail, evalUnknowns
	}
	if len(evalUnknowns) > 0 {
		return false, "", evalUnknowns
	}
	return true, "", nil
}
