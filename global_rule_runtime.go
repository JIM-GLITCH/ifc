package ifc

import (
	"context"
	"sort"
	"strconv"
	"strings"

	core "typefox.dev/fastbelt"
)

type GlobalRuleEval struct {
	ctx        context.Context
	pop        *Population
	entityEval *EntityWhereEval
	sets       map[string][]SIMPLE_ENTITY_INSTANCE
}

func NewGlobalRuleEval(ctx context.Context, pop *Population) *GlobalRuleEval {
	if ctx == nil {
		ctx = context.Background()
	}
	return &GlobalRuleEval{
		ctx:        ctx,
		pop:        pop,
		entityEval: NewEntityWhereEval(ctx, "", nil, nil, pop),
		sets:       buildGlobalEntitySets(pop),
	}
}

func buildGlobalEntitySets(pop *Population) map[string][]SIMPLE_ENTITY_INSTANCE {
	sets := map[string][]SIMPLE_ENTITY_INSTANCE{}
	if pop == nil {
		return sets
	}
	for _, rec := range pop.byName {
		if rec == nil || rec.Instance == nil {
			continue
		}
		def, ok := LookupEntity(rec.StepName)
		if !ok {
			continue
		}
		visited := map[string]bool{}
		var walk func(EntityDef)
		walk = func(def EntityDef) {
			if visited[def.ExpressName] {
				return
			}
			visited[def.ExpressName] = true
			sets[def.ExpressName] = append(sets[def.ExpressName], rec.Instance)
			for _, superName := range def.SuperTypes {
				if superDef, ok := lookupEntityByExpressName(superName); ok {
					walk(superDef)
				}
			}
		}
		walk(def)
	}
	for name, insts := range sets {
		sort.SliceStable(insts, func(i, j int) bool {
			return compareInstanceNames(insts[i].Name(), insts[j].Name()) < 0
		})
		sets[name] = insts
	}
	return sets
}

func compareInstanceNames(a, b string) int {
	na := parseInstanceOrdinal(a)
	nb := parseInstanceOrdinal(b)
	if na != nb {
		return na - nb
	}
	return strings.Compare(a, b)
}

func parseInstanceOrdinal(name string) int {
	name = strings.TrimSpace(name)
	if len(name) > 1 && name[0] == '#' {
		if n, err := strconv.Atoi(name[1:]); err == nil {
			return n
		}
	}
	return 0
}

func (ec *GlobalRuleEval) EntitySet(expressName string) WhereValue {
	key := strings.TrimSpace(expressName)
	if key == "" {
		return wvFromEntityInstances(nil)
	}
	if insts, ok := ec.sets[key]; ok {
		return wvFromEntityInstances(insts)
	}
	for name, insts := range ec.sets {
		if strings.EqualFold(name, key) {
			return wvFromEntityInstances(insts)
		}
	}
	return wvFromEntityInstances(nil)
}

func (ec *GlobalRuleEval) CallSchema(name string, args []WhereValue) WhereValue {
	return ec.entityEval.CallSchema(name, args)
}

func (ec *GlobalRuleEval) QualIndex(v WhereValue, index int) WhereValue {
	return ec.entityEval.QualIndex(v, index)
}

func (ec *GlobalRuleEval) QualEntity(v WhereValue, entityName string) WhereValue {
	return ec.entityEval.QualEntity(v, entityName)
}

func (ec *GlobalRuleEval) QualAttr(v WhereValue, attrName string, index int) WhereValue {
	if inst := v.entityInstance(); inst != nil {
		sub := ec.entityEval.withInstance(inst)
		return sub.QualAttr(sub.Self(), attrName, index)
	}
	return ec.entityEval.QualAttr(v, attrName, index)
}

func CheckGlobalRules(ctx context.Context, pop *Population) (ok bool, detail string) {
	if len(GlobalRuleChecks) == 0 {
		return true, ""
	}
	ec := NewGlobalRuleEval(ctx, pop)
	for _, rule := range GlobalRuleChecks {
		if rule.Check == nil {
			continue
		}
		if pass, msg := rule.Check(ec); !pass {
			if msg == "" {
				return false, "RULE " + rule.ExpressName + " violated"
			}
			if rule.ExpressName != "" && !strings.Contains(msg, rule.ExpressName) {
				return false, rule.ExpressName + ": " + msg
			}
			return false, msg
		}
	}
	return true, ""
}

func validateGlobalRules(ctx context.Context, owner core.AstNode, accept core.ValidationAcceptor) {
	pop := PopulationFromContext(ctx)
	ok, detail := CheckGlobalRules(ctx, pop)
	if ok {
		return
	}
	msg := detail
	if msg == "" {
		msg = "Global RULE constraint violated"
	} else {
		msg = "Global RULE constraint violated: " + detail
	}
	accept(core.NewDiagnostic(core.SeverityError, msg, owner, core.WithCode("ifc-global-rule")))
}
