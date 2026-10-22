package ifc

import (
	"context"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

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
		sets:       pop.entitySets(),
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
		for _, expressName := range stepNameExpressClosure(rec.StepName) {
			sets[expressName] = append(sets[expressName], rec.Instance)
		}
	}
	for name, insts := range sets {
		sort.SliceStable(insts, func(i, j int) bool {
			return compareInstanceNames(insts[i].Name(), insts[j].Name()) < 0
		})
		sets[name] = insts
	}
	return sets
}

// stepNameExpressClosure returns the EXPRESS names of the entity and all of
// its supertypes, deduplicated, self first. Derived from the immutable
// Entities table and cached per step name.
func stepNameExpressClosure(stepName string) []string {
	key := strings.ToUpper(strings.TrimSpace(stepName))
	if cached, ok := expressClosureCache.Load(key); ok {
		names, _ := cached.([]string)
		return names
	}
	names := computeStepNameExpressClosure(key)
	expressClosureCache.Store(key, names)
	return names
}

func computeStepNameExpressClosure(stepName string) []string {
	def, ok := LookupEntity(stepName)
	if !ok {
		return nil
	}
	var names []string
	seen := map[string]bool{}
	var walk func(EntityDef)
	walk = func(def EntityDef) {
		if seen[def.ExpressName] {
			return
		}
		seen[def.ExpressName] = true
		names = append(names, def.ExpressName)
		for _, superName := range def.SuperTypes {
			if superDef, ok := lookupEntityByExpressName(superName); ok {
				walk(superDef)
			}
		}
	}
	walk(def)
	return names
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

func (ec *GlobalRuleEval) WvOrLazy(a, b func() WhereValue) WhereValue {
	if ec == nil || ec.entityEval == nil {
		return wvOr(a(), b())
	}
	return ec.entityEval.WvOrLazy(a, b)
}

func (ec *GlobalRuleEval) WvAndLazy(a, b func() WhereValue) WhereValue {
	if ec == nil || ec.entityEval == nil {
		return wvAnd(a(), b())
	}
	return ec.entityEval.WvAndLazy(a, b)
}

// globalRuleOutcome is one rule's result, collected for order-preserving merge.
type globalRuleOutcome struct {
	pass     bool
	detail   string
	unknowns []string
}

// CheckGlobalRules evaluates every global RULE. Rules are independent, so they
// run concurrently with per-rule eval caches; results are merged in rule order
// so the reported failure and eval-unknowns match the sequential semantics
// (stop at the first rule with eval-unknowns, keep the first failure detail).
func CheckGlobalRules(ctx context.Context, pop *Population) (ok bool, detail string, evalUnknowns []string) {
	if len(GlobalRuleChecks) == 0 {
		return true, "", nil
	}
	outcomes := make([]globalRuleOutcome, len(GlobalRuleChecks))
	workers := runtime.GOMAXPROCS(0)
	if workers > len(GlobalRuleChecks) {
		workers = len(GlobalRuleChecks)
	}
	if workers <= 1 {
		for i, rule := range GlobalRuleChecks {
			outcomes[i] = runGlobalRule(ctx, pop, rule)
		}
	} else {
		var wg sync.WaitGroup
		for w := 0; w < workers; w++ {
			wg.Add(1)
			go func(w int) {
				defer wg.Done()
				for i := w; i < len(GlobalRuleChecks); i += workers {
					outcomes[i] = runGlobalRule(ctx, pop, GlobalRuleChecks[i])
				}
			}(w)
		}
		wg.Wait()
	}

	for i, outcome := range outcomes {
		rule := GlobalRuleChecks[i]
		if rule.Check == nil {
			continue
		}
		evalUnknowns = append(evalUnknowns, outcome.unknowns...)
		if len(outcome.unknowns) > 0 {
			break
		}
		if !outcome.pass && detail == "" {
			if outcome.detail == "" {
				detail = "RULE " + rule.ExpressName + " violated"
			} else if rule.ExpressName != "" && !strings.Contains(outcome.detail, rule.ExpressName) {
				detail = rule.ExpressName + ": " + outcome.detail
			} else {
				detail = outcome.detail
			}
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

func runGlobalRule(ctx context.Context, pop *Population, rule GlobalRuleSpec) globalRuleOutcome {
	if rule.Check == nil {
		return globalRuleOutcome{pass: true}
	}
	// Per-rule eval cache: rules run concurrently and the shared cache map is
	// not safe for concurrent writes.
	ruleCtx := context.WithValue(ctx, entityEvalCacheKey{}, newEntityEvalCache())
	ec := NewGlobalRuleEval(ruleCtx, pop)
	ec.entityEval.whereKind = "global"
	ec.entityEval.ruleLabel = rule.Label
	ec.entityEval.stepName = rule.ExpressName
	pass, msg := rule.Check(ec)
	unknowns := ec.entityEval.takeEvalUnknowns()
	return globalRuleOutcome{pass: pass, detail: msg, unknowns: unknowns}
}

func validateGlobalRules(ctx context.Context, owner core.AstNode, accept core.ValidationAcceptor) {
	pop := PopulationFromContext(ctx)
	ok, detail, _ := CheckGlobalRules(ctx, pop)
	// evalUnknowns from global rule evaluation are side effects while reading
	// entity attributes; they are already reported per-entity during AST validation.
	if ok {
		return
	}
	if detail == "" {
		return
	}
	anchorNode, anchorToken := validationAnchorNode(owner)
	msg := "Global RULE constraint violated: " + detail
	opts := []core.DiagnosticOption{core.WithCode("ifc-global-rule")}
	if anchorToken != nil {
		opts = append(opts, core.WithToken(anchorToken))
	}
	accept(core.NewDiagnostic(core.SeverityError, msg, anchorNode, opts...))
}
