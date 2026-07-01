package ifc

import (
	"fmt"
	"strings"
)

func evalDerivedAttr(ec *EntityWhereEval, inst SIMPLE_ENTITY_INSTANCE, attrName string) WhereValue {
	step, ok := instanceStepName(inst)
	if !ok {
		if ec != nil {
			return ec.unknown(fmt.Sprintf("Cannot evaluate DERIVE attribute %q: invalid entity instance", attrName))
		}
		return wvUnknownVal()
	}
	attrs, ok := entityDerivedEvaluators[strings.ToUpper(strings.TrimSpace(step))]
	if !ok {
		if ec != nil {
			return ec.unknown(fmt.Sprintf("Cannot evaluate DERIVE attribute %q on %s: no evaluator", attrName, step))
		}
		return wvUnknownVal()
	}
	for name, fn := range attrs {
		if strings.EqualFold(name, attrName) {
			v := fn(ec, inst)
			if v.kind == wvKindUnknown && ec != nil {
				ec.noteEvalUnknown(fmt.Sprintf("Cannot evaluate DERIVE attribute %q on %s", attrName, step))
			}
			return v
		}
	}
	if ec != nil {
		return ec.unknown(fmt.Sprintf("Cannot evaluate DERIVE attribute %q on %s: attribute not derived", attrName, step))
	}
	return wvUnknownVal()
}

func (ec *EntityWhereEval) withInstance(inst SIMPLE_ENTITY_INSTANCE) *EntityWhereEval {
	step, ok := instanceStepName(inst)
	if !ok {
		return &EntityWhereEval{ctx: ec.ctx, pop: ec.pop, evalUnknowns: ec.evalUnknowns}
	}
	return &EntityWhereEval{
		ctx:          ec.ctx,
		stepName:     step,
		self:         inst,
		params:       instanceParams(inst),
		pop:          ec.pop,
		evalUnknowns: ec.evalUnknowns,
	}
}
