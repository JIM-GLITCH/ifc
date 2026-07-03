package ifc

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	core "typefox.dev/fastbelt"
)

func lookupExpressTypeName(keyword string) (string, bool) {
	upper := strings.ToUpper(strings.TrimSpace(keyword))
	for name := range Types {
		if strings.ToUpper(name) == upper {
			return name, true
		}
	}
	return "", false
}

func aggregateElementTypeName(typeName string) string {
	upper := strings.ToUpper(typeName)
	idx := strings.Index(upper, " OF ")
	if idx < 0 {
		return ""
	}
	rest := strings.TrimSpace(typeName[idx+4:])
	if len(rest) >= 7 && strings.EqualFold(rest[:7], "UNIQUE ") {
		rest = strings.TrimSpace(rest[7:])
	}
	return rest
}

func isAggregateTypeExpression(typeName string) bool {
	upper := strings.ToUpper(strings.TrimSpace(typeName))
	return strings.HasPrefix(upper, "LIST ") ||
		strings.HasPrefix(upper, "SET ") ||
		strings.HasPrefix(upper, "BAG ") ||
		strings.HasPrefix(upper, "ARRAY ")
}

func schemaAttrForTypeName(typeName string) SchemaAttr {
	if typeName == "" {
		return SchemaAttr{Kind: ParamUnknown}
	}
	if isAggregateTypeExpression(typeName) {
		return SchemaAttr{Kind: ParamAggregate, TypeName: typeName}
	}
	if _, ok := LookupEntity(strings.ToUpper(typeName)); ok {
		return SchemaAttr{Kind: ParamEntity, TypeName: typeName}
	}
	if typeDef, ok := LookupType(typeName); ok {
		kind := typeDef.ValueKind
		if kind == ParamUnknown && typeDef.AliasOf != "" {
			kind = resolvedTypeValueKind(typeDef.AliasOf)
		}
		if kind == ParamUnknown {
			kind = resolvedTypeValueKind(typeName)
		}
		return SchemaAttr{Kind: kind, TypeName: typeName}
	}
	return SchemaAttr{Kind: resolvedTypeValueKind(typeName), TypeName: typeName}
}

func canonicalTypeName(typeName string) string {
	visited := map[string]bool{}
	for typeName != "" && !visited[typeName] {
		visited[typeName] = true
		typeDef, ok := LookupType(typeName)
		if !ok || typeDef.AliasOf == "" {
			break
		}
		typeName = typeDef.AliasOf
	}
	return typeName
}

func typesEquivalent(a, b string) bool {
	return strings.EqualFold(a, b) ||
		strings.EqualFold(canonicalTypeName(a), b) ||
		strings.EqualFold(a, canonicalTypeName(b)) ||
		strings.EqualFold(canonicalTypeName(a), canonicalTypeName(b))
}

func typeIsMemberOfSelect(keywordTypeName, selectTypeName string) bool {
	visited := map[string]bool{}
	return typeIsMemberOfSelectWalk(keywordTypeName, selectTypeName, visited)
}

func typeIsMemberOfSelectWalk(keywordTypeName, selectTypeName string, visited map[string]bool) bool {
	if selectTypeName == "" || visited[selectTypeName] {
		return false
	}
	visited[selectTypeName] = true

	typeDef, ok := LookupType(selectTypeName)
	if !ok || typeDef.ValueKind != ParamSelect {
		return false
	}
	for _, member := range typeDef.SelectMembers {
		if typesEquivalent(keywordTypeName, member) {
			return true
		}
		if memberDef, ok := LookupType(member); ok && memberDef.ValueKind == ParamSelect {
			if typeIsMemberOfSelectWalk(keywordTypeName, member, visited) {
				return true
			}
		}
	}
	return false
}

func keywordMatchesAttribute(keywordTypeName string, attr SchemaAttr) bool {
	if typesEquivalent(keywordTypeName, attr.TypeName) {
		return true
	}
	if attr.Kind == ParamSelect {
		if typeIsMemberOfSelect(keywordTypeName, attr.TypeName) {
			return true
		}
	}
	if attr.TypeName != "" && resolvedTypeValueKind(keywordTypeName) == resolvedTypeValueKind(attr.TypeName) {
		if attr.Kind != ParamSelect {
			return true
		}
	}
	if _, ok := LookupEntity(strings.ToUpper(keywordTypeName)); ok {
		if attr.Kind == ParamEntity {
			return entityIsOrSubtypes(strings.ToUpper(keywordTypeName), attr.TypeName)
		}
		return false
	}
	return false
}

func normalizeEnumValue(raw string) string {
	return strings.ToUpper(strings.Trim(raw, ". \t"))
}

func enumValuesForAttr(attr SchemaAttr) []string {
	if typeDef, ok := LookupType(attr.TypeName); ok && len(typeDef.EnumValues) > 0 {
		return typeDef.EnumValues
	}
	return nil
}

func validateEnumerationValue(
	attr SchemaAttr,
	literal ENUMERATION_LITERAL,
	accept core.ValidationAcceptor,
) {
	values := enumValuesForAttr(attr)
	if len(values) == 0 {
		return
	}
	actual := normalizeEnumValue(literal.Value())
	for _, expected := range values {
		if actual == normalizeEnumValue(expected) {
			return
		}
	}
	opts := append(diagnosticTokenOptions(literal.ValueToken()), core.WithCode("ifc-enum-value"))
	accept(core.NewDiagnostic(
		core.SeverityError,
		fmt.Sprintf("Invalid enumeration value %q for type %q (expected one of: %s)", literal.Value(), attr.TypeName, strings.Join(values, ", ")),
		literal,
		opts...,
	))
}

func validateTypedParameter(
	ctx context.Context,
	owner core.AstNode,
	attr SchemaAttr,
	typed TYPED_PARAMETER,
	position int,
	accept core.ValidationAcceptor,
) {
	keyword := typed.Keyword()
	expressType, ok := lookupExpressTypeName(keyword)
	if !ok {
		opts := append(diagnosticTokenOptions(typed.KeywordToken()), core.WithCode("ifc-typed-keyword"))
		accept(core.NewDiagnostic(
			core.SeverityError,
			fmt.Sprintf("Unknown typed parameter keyword %q", keyword),
			typed,
			opts...,
		))
		return
	}
	if !keywordMatchesAttribute(expressType, attr) {
		opts := append(diagnosticTokenOptions(typed.KeywordToken()), core.WithCode("ifc-typed-keyword"))
		accept(core.NewDiagnostic(
			core.SeverityError,
			fmt.Sprintf("Typed parameter %q is not valid for attribute type %q", keyword, typeDescription(attr)),
			typed,
			opts...,
		))
		return
	}
	inner := typed.Parameter()
	if inner == nil {
		opts := append(diagnosticTokenOptions(typed.KeywordToken()), core.WithCode("ifc-typed-keyword"))
		accept(core.NewDiagnostic(
			core.SeverityError,
			fmt.Sprintf("Typed parameter %q is missing a value", keyword),
			typed,
			opts...,
		))
		return
	}
	innerAttr := schemaAttrForTypeName(expressType)
	validateParameterValueCore(ctx, owner, innerAttr, inner, position, accept)
	validateTypeWhereConstraints(ctx, innerAttr, inner, inner, parameterInfoFrom(inner).token, accept)
}

func aggregateExpressionForAttr(attr SchemaAttr) string {
	upper := strings.ToUpper(attr.TypeName)
	if strings.Contains(upper, " OF ") {
		return attr.TypeName
	}
	if typeDef, ok := LookupType(attr.TypeName); ok && typeDef.AliasOf != "" {
		return typeDef.AliasOf
	}
	return attr.TypeName
}

type aggregateBounds struct {
	low  int // -1 when unbounded (?)
	high int // -1 when unbounded (?)
}

func parseAggregateBounds(typeName string) (aggregateBounds, bool) {
	upper := strings.ToUpper(strings.TrimSpace(typeName))
	prefixLen := 0
	for _, prefix := range []string{"LIST ", "SET ", "BAG ", "ARRAY "} {
		if strings.HasPrefix(upper, prefix) {
			prefixLen = len(prefix)
			break
		}
	}
	if prefixLen == 0 {
		return aggregateBounds{}, false
	}
	rest := strings.TrimSpace(typeName[prefixLen:])
	start := strings.Index(rest, "[")
	if start < 0 {
		return aggregateBounds{}, false
	}
	end := strings.Index(rest[start:], "]")
	if end < 0 {
		return aggregateBounds{}, false
	}
	inner := rest[start+1 : start+end]
	colon := strings.Index(inner, ":")
	if colon < 0 {
		return aggregateBounds{}, false
	}
	lowStr := strings.TrimSpace(inner[:colon])
	highStr := strings.TrimSpace(inner[colon+1:])

	bounds := aggregateBounds{low: -1, high: -1}
	if lowStr != "?" {
		low, err := strconv.Atoi(lowStr)
		if err != nil {
			return aggregateBounds{}, false
		}
		bounds.low = low
	}
	if highStr != "?" {
		high, err := strconv.Atoi(highStr)
		if err != nil {
			return aggregateBounds{}, false
		}
		bounds.high = high
	}
	return bounds, true
}

func aggregateBoundsForAttr(attr SchemaAttr) (aggregateBounds, bool) {
	expr := aggregateExpressionForAttr(attr)
	if !isAggregateTypeExpression(expr) {
		return aggregateBounds{}, false
	}
	return parseAggregateBounds(expr)
}

func validateAggregateBoundsCount(
	owner core.AstNode,
	attr SchemaAttr,
	count int,
	accept core.ValidationAcceptor,
) {
	bounds, ok := aggregateBoundsForAttr(attr)
	if !ok {
		return
	}
	typeDesc := typeDescription(attr)
	if bounds.low >= 0 && count < bounds.low {
		opts := []core.DiagnosticOption{core.WithCode("ifc-aggregate-bounds")}
		accept(core.NewDiagnostic(
			core.SeverityError,
			fmt.Sprintf("Aggregate for %q expects at least %d element(s), got %d", typeDesc, bounds.low, count),
			owner,
			opts...,
		))
	}
	if bounds.high >= 0 && count > bounds.high {
		opts := []core.DiagnosticOption{core.WithCode("ifc-aggregate-bounds")}
		accept(core.NewDiagnostic(
			core.SeverityError,
			fmt.Sprintf("Aggregate for %q expects at most %d element(s), got %d", typeDesc, bounds.high, count),
			owner,
			opts...,
		))
	}
}

func validateAggregateBounds(
	ctx context.Context,
	owner core.AstNode,
	attr SchemaAttr,
	list LIST,
	accept core.ValidationAcceptor,
) {
	if !validationEnabled(ctx, LayerAggregateBounds) {
		return
	}
	validateAggregateBoundsCount(owner, attr, len(list.Parameters()), accept)
}

func validateAggregateElements(
	ctx context.Context,
	owner core.AstNode,
	attr SchemaAttr,
	list LIST,
	accept core.ValidationAcceptor,
) {
	validateAggregateBounds(ctx, owner, attr, list, accept)
	elementType := aggregateElementTypeName(aggregateExpressionForAttr(attr))
	if elementType == "" {
		return
	}
	aggExpr := aggregateExpressionForAttr(attr)
	if aggregateHasUniqueElements(aggExpr) && validationEnabled(ctx, LayerUnique) {
		validateAggregateUniqueElements(ctx, owner, list, accept)
	}
	elementAttr := schemaAttrForTypeName(elementType)
	for _, element := range list.Parameters() {
		validateParameterValue(ctx, owner, "", "", elementAttr, element, -1, accept)
	}
}
