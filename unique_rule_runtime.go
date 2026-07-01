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

type UniqueEvalFailure struct {
	Instance SIMPLE_ENTITY_INSTANCE
	Message  string
}

func uniqueKeyPart(v WhereValue) (string, bool) {
	if v.kind == wvKindUnknown {
		return "", false
	}
	switch v.kind {
	case wvKindString:
		return "S:" + strings.ToLower(v.s), true
	case wvKindEnum:
		return "N:" + strings.ToLower(normalizeEnumMember(v.s)), true
	case wvKindInteger:
		return fmt.Sprintf("I:%d", v.i), true
	case wvKindReal:
		return fmt.Sprintf("R:%g", v.r), true
	case wvKindLogical:
		return fmt.Sprintf("L:%d", v.log), true
	case wvKindEntityInst:
		if v.entity != nil {
			return "REF:" + strings.ToUpper(strings.TrimSpace(v.entity.Name())), true
		}
		return "", false
	case wvKindList:
		parts := make([]string, len(v.list))
		for i, item := range v.list {
			part, ok := uniqueKeyPart(item)
			if !ok {
				return "", false
			}
			parts[i] = part
		}
		return "[" + strings.Join(parts, ",") + "]", true
	case wvKindRecord:
		names := make([]string, 0, len(v.recordAttrs))
		for name := range v.recordAttrs {
			names = append(names, name)
		}
		sort.Strings(names)
		parts := make([]string, len(names))
		for i, name := range names {
			part, ok := uniqueKeyPart(v.recordAttrs[name])
			if !ok {
				return "", false
			}
			parts[i] = name + "=" + part
		}
		return "{" + strings.Join(parts, ",") + "}", true
	default:
		return "", false
	}
}

func instanceUniqueKey(ec *EntityWhereEval, attrNames []string) (string, bool) {
	parts := make([]string, len(attrNames))
	for i, name := range attrNames {
		value := ec.Attr(name)
		part, ok := uniqueKeyPart(value)
		if !ok {
			return "", false
		}
		parts[i] = part
	}
	return strings.Join(parts, "\x1e"), true
}

func parameterUniqueKey(ctx context.Context, pop *Population, param PARAMETER) (string, bool) {
	v, ok := parameterToWhereValueCtx(ctx, pop, param)
	if !ok {
		return "", false
	}
	part, ok := uniqueKeyPart(v)
	return part, ok
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

func CheckUniqueRules(ctx context.Context, pop *Population) ([]UniqueViolation, []UniqueEvalFailure) {
	if len(UniqueRuleChecks) == 0 || pop == nil {
		return nil, nil
	}
	eval := NewGlobalRuleEval(ctx, pop)
	var violations []UniqueViolation
	var failures []UniqueEvalFailure
	for _, rule := range UniqueRuleChecks {
		instances := whereValueEntityInstances(eval.EntitySet(rule.OwnerExpressName))
		if len(instances) < 2 {
			continue
		}
		byKey := map[string][]SIMPLE_ENTITY_INSTANCE{}
		for _, inst := range instances {
			step, ok := instanceStepName(inst)
			if !ok {
				continue
			}
			ec := NewEntityWhereEval(ctx, step, inst, instanceParams(inst), pop)
			key, keyOK := instanceUniqueKey(ec, rule.AttrNames)
			notes := ec.takeEvalUnknowns()
			if len(notes) > 0 {
				for _, msg := range notes {
					failures = append(failures, UniqueEvalFailure{Instance: inst, Message: msg})
				}
				continue
			}
			if !keyOK {
				failures = append(failures, UniqueEvalFailure{
					Instance: inst,
					Message:  fmt.Sprintf("Cannot evaluate UNIQUE key on %s", rule.OwnerExpressName),
				})
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
	return violations, failures
}

func validateUniqueRules(ctx context.Context, accept core.ValidationAcceptor) {
	pop := PopulationFromContext(ctx)
	violations, failures := CheckUniqueRules(ctx, pop)
	for _, failure := range failures {
		if failure.Instance == nil {
			continue
		}
		acceptEvalUnknown(accept, failure.Instance, failure.Instance.NameToken(), failure.Message)
	}
	for _, violation := range violations {
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
			info := parameterInfoFrom(element)
			acceptEvalUnknown(accept, diagnosticNodeForParameter(owner, element), info.token, parameterEvalFailureReason(ctx, pop, element))
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
