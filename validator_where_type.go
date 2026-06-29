package ifc

import (
	"context"
	"fmt"
	"strings"

	core "typefox.dev/fastbelt"
)

func resolveEntityReference(ctx context.Context, pop *Population, ref ENTITY_REFERENCE) SIMPLE_ENTITY_INSTANCE {
	value := ref.Value()
	if value == nil {
		return nil
	}
	if target := value.Ref(ctx); target != nil {
		if simple, ok := target.(SIMPLE_ENTITY_INSTANCE); ok {
			return simple
		}
	}
	if pop == nil {
		pop = PopulationFromContext(ctx)
	}
	if pop != nil {
		if rec, ok := pop.Lookup(strings.TrimSpace(value.Text())); ok && rec.Instance != nil {
			return rec.Instance
		}
	}
	return nil
}

func parameterToWhereValueCtx(ctx context.Context, pop *Population, param PARAMETER) (WhereValue, bool) {
	if param == nil {
		return WhereValue{}, false
	}
	switch p := param.(type) {
	case INTEGER_LITERAL:
		return parseLiteralNumber(p.Value()), true
	case REAL_LITERAL:
		return parseLiteralNumber(p.Value()), true
	case STRING_LITERAL:
		return wvFromString(unquoteIFCString(p.Value())), true
	case ENUMERATION_LITERAL:
		return wvFromEnum(p.Value()), true
	case ENTITY_REFERENCE:
		if inst := resolveEntityReference(ctx, pop, p); inst != nil {
			return wvFromEntityInstance(inst), true
		}
		return wvUnknownVal(), false
	case DEFAULT_PARAMETER, OMITTED_PARAMETER:
		return wvUnknownVal(), true
	case LIST:
		items := make([]WhereValue, 0, len(p.Parameters()))
		for _, el := range p.Parameters() {
			v, ok := parameterToWhereValueCtx(ctx, pop, el)
			if !ok {
				return WhereValue{}, false
			}
			items = append(items, v)
		}
		return wvFromList(items), true
	case TYPED_PARAMETER:
		inner := p.Parameter()
		if inner == nil {
			return WhereValue{}, false
		}
		return parameterToWhereValueCtx(ctx, pop, inner)
	default:
		return WhereValue{}, false
	}
}

func whereTypeNameForAttr(attr SchemaAttr) string {
	if attr.TypeName == "" {
		return ""
	}
	return attr.TypeName
}

func parameterToWhereValue(param PARAMETER) (WhereValue, bool) {
	return parameterToWhereValueCtx(context.Background(), nil, param)
}

func unquoteIFCString(raw string) string {
	raw = strings.TrimSpace(raw)
	if len(raw) >= 2 && raw[0] == '\'' && raw[len(raw)-1] == '\'' {
		return strings.ReplaceAll(raw[1:len(raw)-1], "''", "'")
	}
	return raw
}

func shouldSkipTypeWhere(attr SchemaAttr, param PARAMETER) bool {
	info := parameterInfoFrom(param)
	if info.kind == paramKindOmitted && attr.Derived {
		return true
	}
	if info.kind == paramKindUnset && attr.Optional {
		return true
	}
	return false
}

func validateTypeWhereConstraints(
	attr SchemaAttr,
	param PARAMETER,
	diagnosticNode core.AstNode,
	token *core.Token,
	accept core.ValidationAcceptor,
) {
	if shouldSkipTypeWhere(attr, param) {
		return
	}
	typeName := whereTypeNameForAttr(attr)
	if typeName == "" {
		return
	}
	rules, ok := LookupTypeWhereRules(typeName)
	if !ok || len(rules) == 0 {
		return
	}
	value, convOK := parameterToWhereValue(param)
	if !convOK {
		return
	}
	ec := NewTypeWhereEval(value)
	for _, rule := range rules {
		if rule.Check == nil {
			continue
		}
		pass, detail := rule.Check(ec)
		if pass {
			continue
		}
		msg := detail
		if msg == "" {
			msg = fmt.Sprintf("Value violates WHERE constraint %q on type %q", rule.Label, typeName)
		} else {
			msg = fmt.Sprintf("Type %q constraint violated: %s", typeName, detail)
		}
		opts := append(diagnosticTokenOptions(token), core.WithCode("ifc-type-where"))
		accept(core.NewDiagnostic(core.SeverityError, msg, diagnosticNode, opts...))
	}
}

func validateTypeWhereForParameter(
	ctx context.Context,
	owner core.AstNode,
	attr SchemaAttr,
	param PARAMETER,
	accept core.ValidationAcceptor,
) {
	_ = ctx
	info := parameterInfoFrom(param)
	validateTypeWhereConstraints(attr, param, info.node, info.token, accept)
}
