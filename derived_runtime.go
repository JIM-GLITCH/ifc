package ifc

import "strings"

func evalDerivedAttr(ec *EntityWhereEval, inst SIMPLE_ENTITY_INSTANCE, attrName string) WhereValue {
	step, ok := instanceStepName(inst)
	if !ok {
		return wvUnknownVal()
	}
	attrs, ok := entityDerivedEvaluators[strings.ToUpper(strings.TrimSpace(step))]
	if !ok {
		return wvUnknownVal()
	}
	for name, fn := range attrs {
		if strings.EqualFold(name, attrName) {
			return fn(ec, inst)
		}
	}
	return wvUnknownVal()
}

func (ec *EntityWhereEval) withInstance(inst SIMPLE_ENTITY_INSTANCE) *EntityWhereEval {
	step, ok := instanceStepName(inst)
	if !ok {
		return &EntityWhereEval{ctx: ec.ctx, pop: ec.pop}
	}
	return &EntityWhereEval{
		ctx:      ec.ctx,
		stepName: step,
		self:     inst,
		params:   instanceParams(inst),
		pop:      ec.pop,
	}
}
