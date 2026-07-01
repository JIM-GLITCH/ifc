package ifc

import (
	"context"
	"fmt"
	"sort"
	"strings"

	core "typefox.dev/fastbelt"
)

type UniqueViolation struct {
	Rule      UniqueRuleSpec
	Instances []SIMPLE_ENTITY_INSTANCE
	Detail    string
}

func uniqueKeyPart(v WhereValue) string {
	if v.kind == wvKindUnknown {
		return "*"
	}
	switch v.kind {
	case wvKindString:
		return "S:" + strings.ToLower(v.s)
	case wvKindEnum:
		return "N:" + strings.ToLower(normalizeEnumMember(v.s))
	case wvKindInteger:
		return fmt.Sprintf("I:%d", v.i)
	case wvKindReal:
		return fmt.Sprintf("R:%g", v.r)
	case wvKindLogical:
		return fmt.Sprintf("L:%d", v.log)
	case wvKindEntityInst:
		if v.entity != nil {
			return "REF:" + strings.ToUpper(strings.TrimSpace(v.entity.Name()))
		}
		return "*"
	case wvKindList:
		parts := make([]string, len(v.list))
		for i, item := range v.list {
			parts[i] = uniqueKeyPart(item)
		}
		return "[" + strings.Join(parts, ",") + "]"
	case wvKindRecord:
		names := make([]string, 0, len(v.recordAttrs))
		for name := range v.recordAttrs {
			names = append(names, name)
		}
		sort.Strings(names)
		parts := make([]string, len(names))
		for i, name := range names {
			parts[i] = name + "=" + uniqueKeyPart(v.recordAttrs[name])
		}
		return "{" + strings.Join(parts, ",") + "}"
	default:
		return "*"
	}
}

func instanceUniqueKey(ctx context.Context, pop *Population, inst SIMPLE_ENTITY_INSTANCE, attrNames []string) (string, bool) {
	step, ok := instanceStepName(inst)
	if !ok {
		return "", false
	}
	ec := NewEntityWhereEval(ctx, step, inst, instanceParams(inst), pop)
	parts := make([]string, len(attrNames))
	for i, name := range attrNames {
		parts[i] = uniqueKeyPart(ec.Attr(name))
	}
	return strings.Join(parts, "\x1e"), true
}

func parameterUniqueKey(ctx context.Context, pop *Population, param PARAMETER) (string, bool) {
	v, ok := parameterToWhereValueCtx(ctx, pop, param)
	if !ok {
		return "", false
	}
	return uniqueKeyPart(v), true
}

func whereValueEntityInstances(v WhereValue) []SIMPLE_ENTITY_INSTANCE {
	if v.kind == wvKindEntityList {
		return append([]SIMPLE_ENTITY_INSTANCE(nil), v.entities...)
	}
	return nil
}

func formatInstanceNames(insts []SIMPLE_ENTITY_INSTANCE) string {
	names := make([]string, 0, len(insts))
	for _, inst := range insts {
		if inst == nil {
			continue
		}
		names = append(names, inst.Name())
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

func CheckUniqueRules(ctx context.Context, pop *Population) []UniqueViolation {
	if len(UniqueRuleChecks) == 0 || pop == nil {
		return nil
	}
	eval := NewGlobalRuleEval(ctx, pop)
	var violations []UniqueViolation
	for _, rule := range UniqueRuleChecks {
		instances := whereValueEntityInstances(eval.EntitySet(rule.OwnerExpressName))
		if len(instances) < 2 {
			continue
		}
		byKey := map[string][]SIMPLE_ENTITY_INSTANCE{}
		for _, inst := range instances {
			key, ok := instanceUniqueKey(ctx, pop, inst, rule.AttrNames)
			if !ok {
				continue
			}
			byKey[key] = append(byKey[key], inst)
		}
		for _, group := range byKey {
			if len(group) < 2 {
				continue
			}
			sort.SliceStable(group, func(i, j int) bool {
				return compareInstanceNames(group[i].Name(), group[j].Name()) < 0
			})
			label := rule.Label
			if label == "" {
				label = strings.Join(rule.AttrNames, ", ")
			}
			detail := fmt.Sprintf(
				"UNIQUE %s on %s violated for %s: duplicate among %s",
				label,
				rule.OwnerExpressName,
				strings.Join(rule.AttrNames, ", "),
				formatInstanceNames(group),
			)
			violations = append(violations, UniqueViolation{
				Rule:      rule,
				Instances: append([]SIMPLE_ENTITY_INSTANCE(nil), group...),
				Detail:    detail,
			})
		}
	}
	return violations
}

func validateUniqueRules(ctx context.Context, accept core.ValidationAcceptor) {
	pop := PopulationFromContext(ctx)
	for _, violation := range CheckUniqueRules(ctx, pop) {
		for _, inst := range violation.Instances {
			if inst == nil {
				continue
			}
			opts := []core.DiagnosticOption{core.WithCode("ifc-unique")}
			if token := inst.NameToken(); token != nil {
				opts = append(opts, core.WithToken(token))
			}
			accept(core.NewDiagnostic(core.SeverityError, violation.Detail, inst, opts...))
		}
	}
}

func aggregateHasUniqueElements(typeName string) bool {
	upper := strings.ToUpper(strings.TrimSpace(typeName))
	idx := strings.Index(upper, " OF ")
	if idx < 0 {
		return false
	}
	rest := strings.TrimSpace(typeName[idx+4:])
	return len(rest) >= 7 && strings.EqualFold(rest[:7], "UNIQUE ")
}

func validateAggregateUniqueElements(
	ctx context.Context,
	owner core.AstNode,
	list LIST,
	accept core.ValidationAcceptor,
) {
	pop := PopulationFromContext(ctx)
	seen := map[string]PARAMETER{}
	for _, element := range list.Parameters() {
		key, ok := parameterUniqueKey(ctx, pop, element)
		if !ok {
			continue
		}
		if prev, dup := seen[key]; dup {
			opts := []core.DiagnosticOption{core.WithCode("ifc-unique-aggregate")}
			if info := parameterInfoFrom(element); info.token != nil {
				opts = append(opts, core.WithToken(info.token))
			}
			accept(core.NewDiagnostic(
				core.SeverityError,
				fmt.Sprintf("Duplicate element in UNIQUE aggregate (same value as earlier element %q)", parameterSummary(prev)),
				diagnosticNodeForParameter(owner, element),
				opts...,
			))
			continue
		}
		seen[key] = element
	}
}

func parameterSummary(param PARAMETER) string {
	if param == nil {
		return ""
	}
	switch p := param.(type) {
	case ENTITY_REFERENCE:
		return strings.TrimSpace(p.Value().Text())
	case STRING_LITERAL:
		return p.Value()
	default:
		return fmt.Sprintf("%T", param)
	}
}

func diagnosticNodeForParameter(owner core.AstNode, param PARAMETER) core.AstNode {
	if info := parameterInfoFrom(param); info.node != nil {
		return info.node
	}
	return owner
}
