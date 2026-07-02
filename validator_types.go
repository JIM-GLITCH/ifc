package ifc

import (
	"context"
	"fmt"
	"strings"

	core "typefox.dev/fastbelt"
)

const (
	paramKindUnset   ParamKind = -1
	paramKindOmitted ParamKind = -2
)

type parameterInfo struct {
	kind  ParamKind
	node  core.AstNode
	token *core.Token
}

func validateParameterCount(def EntityDef, paramCount int) (ok bool, message string) {
	attrs := def.Attrs
	maxCount := len(attrs)
	if paramCount > maxCount {
		return false, fmt.Sprintf("Entity %q accepts at most %d parameter(s), got %d", def.ExpressName, maxCount, paramCount)
	}
	for i := 0; i < maxCount; i++ {
		if !attrs[i].Optional && i >= paramCount {
			return false, fmt.Sprintf("Entity %q is missing required parameter at position %d (got %d parameter(s))", def.ExpressName, i+1, paramCount)
		}
	}
	return true, ""
}

func parameterInfoFrom(param PARAMETER) parameterInfo {
	if param == nil {
		return parameterInfo{kind: ParamUnknown}
	}
	switch p := param.(type) {
	case INTEGER_LITERAL:
		return parameterInfo{kind: ParamInteger, node: p, token: p.ValueToken()}
	case REAL_LITERAL:
		return parameterInfo{kind: ParamReal, node: p, token: p.ValueToken()}
	case STRING_LITERAL:
		return parameterInfo{kind: ParamString, node: p, token: p.ValueToken()}
	case ENUMERATION_LITERAL:
		return parameterInfo{kind: ParamEnum, node: p, token: p.ValueToken()}
	case BINARY_LITERAL:
		return parameterInfo{kind: ParamBinary, node: p, token: p.ValueToken()}
	case ENTITY_REFERENCE:
		return parameterInfo{kind: ParamEntity, node: p}
	case LIST:
		return parameterInfo{kind: ParamAggregate, node: p, token: listToken(p)}
	case DEFAULT_PARAMETER:
		return parameterInfo{kind: paramKindUnset, node: p, token: p.ValueToken()}
	case OMITTED_PARAMETER:
		return parameterInfo{kind: paramKindOmitted, node: p, token: p.ValueToken()}
	case TYPED_PARAMETER:
		return parameterInfoFrom(p.Parameter())
	default:
		return parameterInfo{kind: ParamUnknown, node: p}
	}
}

func listToken(list LIST) *core.Token {
	return nil
}

func builtinTypeKind(typeName string) ParamKind {
	switch strings.ToUpper(strings.TrimSpace(typeName)) {
	case "INTEGER":
		return ParamInteger
	case "REAL", "NUMBER":
		return ParamReal
	case "STRING":
		return ParamString
	case "BOOLEAN":
		return ParamBoolean
	case "LOGICAL":
		return ParamLogical
	case "BINARY":
		return ParamBinary
	default:
		return ParamUnknown
	}
}

func resolvedTypeValueKind(typeName string) ParamKind {
	if typeName == "" {
		return ParamUnknown
	}
	visited := map[string]bool{}
	for typeName != "" && !visited[typeName] {
		visited[typeName] = true
		if kind := builtinTypeKind(typeName); kind != ParamUnknown {
			return kind
		}
		typeDef, ok := LookupType(typeName)
		if !ok {
			return ParamUnknown
		}
		if typeDef.ValueKind != ParamUnknown {
			if typeDef.AliasOf == "" || typeDef.ValueKind != ParamEntity {
				return typeDef.ValueKind
			}
		}
		if typeDef.AliasOf == "" {
			return typeDef.ValueKind
		}
		typeName = typeDef.AliasOf
	}
	return ParamUnknown
}

func allowedParameterKinds(attr SchemaAttr) []ParamKind {
	switch attr.Kind {
	case ParamEntity:
		return []ParamKind{ParamEntity}
	case ParamEnum:
		return []ParamKind{ParamEnum}
	case ParamAggregate:
		return []ParamKind{ParamAggregate, ParamEntity}
	case ParamSelect:
		return allowedSelectKinds(attr.TypeName)
	case ParamReal, ParamNumber:
		return []ParamKind{ParamReal, ParamInteger}
	case ParamInteger:
		return []ParamKind{ParamInteger}
	case ParamString:
		return []ParamKind{ParamString}
	case ParamBinary:
		return []ParamKind{ParamBinary}
	case ParamBoolean:
		return []ParamKind{ParamEnum, ParamInteger}
	case ParamLogical:
		return []ParamKind{ParamEnum, ParamInteger}
	default:
		if attr.TypeName != "" {
			switch resolvedTypeValueKind(attr.TypeName) {
			case ParamString:
				return []ParamKind{ParamString}
			case ParamReal, ParamNumber:
				return []ParamKind{ParamReal, ParamInteger}
			case ParamInteger:
				return []ParamKind{ParamInteger}
			case ParamEnum:
				return []ParamKind{ParamEnum}
			case ParamSelect:
				return allowedSelectKinds(attr.TypeName)
			}
		}
		return []ParamKind{attr.Kind}
	}
}

func allowedSelectKinds(typeName string) []ParamKind {
	typeDef, ok := LookupType(typeName)
	if !ok {
		return []ParamKind{ParamEntity, ParamAggregate, ParamString, ParamReal, ParamInteger, ParamEnum}
	}
	kinds := make([]ParamKind, 0, 8)
	seen := map[ParamKind]bool{}
	add := func(kind ParamKind) {
		if kind == ParamUnknown || seen[kind] {
			return
		}
		seen[kind] = true
		kinds = append(kinds, kind)
	}
	for _, member := range typeDef.SelectMembers {
		if _, ok := LookupEntity(strings.ToUpper(member)); ok {
			add(ParamEntity)
			continue
		}
		switch resolvedTypeValueKind(member) {
		case ParamString:
			add(ParamString)
		case ParamReal, ParamNumber:
			add(ParamReal)
			add(ParamInteger)
		case ParamInteger:
			add(ParamInteger)
		case ParamEnum:
			add(ParamEnum)
		case ParamAggregate:
			add(ParamAggregate)
		case ParamEntity:
			add(ParamEntity)
		default:
			add(ParamAggregate)
		}
	}
	if len(kinds) == 0 {
		return []ParamKind{ParamEntity, ParamAggregate, ParamString, ParamReal, ParamInteger, ParamEnum}
	}
	return kinds
}

func parameterKindAllowed(attr SchemaAttr, actual ParamKind) bool {
	if actual == paramKindOmitted {
		return attr.Derived
	}
	if actual == paramKindUnset {
		return attr.Optional
	}
	for _, allowed := range allowedParameterKinds(attr) {
		if allowed == actual {
			return true
		}
	}
	return false
}

func entityExpressName(stepName string) string {
	if def, ok := LookupEntity(stepName); ok {
		return def.ExpressName
	}
	return stepName
}

func entityIsOrSubtypes(actualStepName, expectedExpressName string) bool {
	if strings.EqualFold(actualStepName, expectedExpressName) {
		return true
	}
	actualDef, ok := LookupEntity(actualStepName)
	if !ok {
		return false
	}
	if strings.EqualFold(actualDef.ExpressName, expectedExpressName) {
		return true
	}
	visited := map[string]bool{}
	var walk func(EntityDef) bool
	walk = func(def EntityDef) bool {
		if visited[def.ExpressName] {
			return false
		}
		visited[def.ExpressName] = true
		for _, superName := range def.SuperTypes {
			if strings.EqualFold(superName, expectedExpressName) {
				return true
			}
			if superDef, ok := LookupEntity(strings.ToUpper(superName)); ok {
				if walk(superDef) {
					return true
				}
			}
		}
		return false
	}
	return walk(actualDef)
}

func entityMatchesAttribute(actualStepName string, attr SchemaAttr) bool {
	switch attr.Kind {
	case ParamEntity:
		return entityIsOrSubtypes(actualStepName, attr.TypeName)
	case ParamSelect:
		typeDef, ok := LookupType(attr.TypeName)
		if !ok {
			return true
		}
		for _, member := range typeDef.SelectMembers {
			if _, ok := LookupEntity(strings.ToUpper(member)); ok {
				if entityIsOrSubtypes(actualStepName, member) {
					return true
				}
				continue
			}
		}
		return false
	default:
		return true
	}
}

func referencedEntityStepName(ctx context.Context, ref *core.Reference[SIMPLE_ENTITY_INSTANCE]) (string, bool) {
	if ref == nil {
		return "", false
	}
	target := ref.Ref(ctx)
	if target == nil {
		return "", false
	}
	switch inst := target.(type) {
	case SIMPLE_ENTITY_INSTANCE:
		record := inst.Record()
		if record == nil {
			return "", false
		}
		return record.Keyword(), true
	default:
		return "", false
	}
}

func validateEntityReference(
	ctx context.Context,
	owner core.AstNode,
	attr SchemaAttr,
	param ENTITY_REFERENCE,
	accept core.ValidationAcceptor,
) {
	ref := param.Value()
	if ref == nil {
		return
	}
	pop := PopulationFromContext(ctx)
	inst := resolveEntityReference(ctx, pop, param)
	if inst == nil {
		// Unresolved # references are reported by the fastbelt linker.
		return
	}
	actualStepName, ok := instanceStepName(inst)
	if !ok {
		return
	}
	if entityMatchesAttribute(actualStepName, attr) {
		return
	}
	expected := attr.TypeName
	if attr.Kind == ParamSelect {
		expected = fmt.Sprintf("one of %s", selectMemberSummary(attr.TypeName))
	}
	opts := []core.DiagnosticOption{core.WithCode("ifc-entity-type"), core.WithReference(ref)}
	accept(core.NewDiagnostic(
		core.SeverityError,
		fmt.Sprintf("Expected %s for attribute type %q, got %q", expected, attr.TypeName, entityExpressName(actualStepName)),
		owner,
		opts...,
	))
}

func selectMemberSummary(typeName string) string {
	typeDef, ok := LookupType(typeName)
	if !ok || len(typeDef.SelectMembers) == 0 {
		return typeName
	}
	return strings.Join(typeDef.SelectMembers, ", ")
}

func diagnosticTokenOptions(token *core.Token) []core.DiagnosticOption {
	if token == nil {
		return nil
	}
	return []core.DiagnosticOption{core.WithToken(token)}
}

func validateParameterValue(
	ctx context.Context,
	owner core.AstNode,
	stepName string,
	attrName string,
	attr SchemaAttr,
	param PARAMETER,
	position int,
	accept core.ValidationAcceptor,
) {
	if typed, ok := param.(TYPED_PARAMETER); ok {
		validateTypedParameter(ctx, owner, attr, typed, position, accept)
		return
	}

	validateParameterValueCore(ctx, owner, attr, param, position, accept)
	validateTypeWhereForParameter(ctx, owner, stepName, attrName, attr, param, accept)
}

func validateParameterValueCore(
	ctx context.Context,
	owner core.AstNode,
	attr SchemaAttr,
	param PARAMETER,
	position int,
	accept core.ValidationAcceptor,
) {
	info := parameterInfoFrom(param)
	diagnosticNode := info.node
	if diagnosticNode == nil {
		diagnosticNode = owner
	}
	if info.kind == paramKindOmitted {
		if attr.Derived {
			return
		}
		opts := append(diagnosticTokenOptions(info.token), core.WithCode("ifc-param-derived"))
		accept(core.NewDiagnostic(
			core.SeverityError,
			fmt.Sprintf("Parameter at position %d must be omitted with * only for derived attributes", position+1),
			diagnosticNode,
			opts...,
		))
		return
	}
	if info.kind == paramKindUnset {
		if !attr.Optional {
			opts := append(diagnosticTokenOptions(info.token), core.WithCode("ifc-param-required"))
			accept(core.NewDiagnostic(
				core.SeverityError,
				fmt.Sprintf("Parameter at position %d requires a value", position+1),
				diagnosticNode,
				opts...,
			))
		}
		return
	}
	if !parameterKindAllowed(attr, info.kind) {
		opts := append(diagnosticTokenOptions(info.token), core.WithCode("ifc-param-type"))
		msg := fmt.Sprintf("Parameter at position %d has incompatible value type for %q", position+1, typeDescription(attr))
		if position < 0 {
			msg = fmt.Sprintf("Aggregate element has incompatible value type for %q", typeDescription(attr))
		}
		accept(core.NewDiagnostic(
			core.SeverityError,
			msg,
			diagnosticNode,
			opts...,
		))
		return
	}
	if enumLit, ok := param.(ENUMERATION_LITERAL); ok {
		validateEnumerationValue(attr, enumLit, accept)
	}
	if info.kind == ParamEntity {
		if entityRef, ok := param.(ENTITY_REFERENCE); ok {
			validateEntityReference(ctx, owner, attr, entityRef, accept)
		}
	}
	if list, ok := param.(LIST); ok {
		validateAggregateElements(ctx, owner, attr, list, accept)
	}
}

func typeDescription(attr SchemaAttr) string {
	if attr.TypeName != "" {
		return attr.TypeName
	}
	return attr.Kind.String()
}

func validateParameterTypes(
	ctx context.Context,
	owner core.AstNode,
	def EntityDef,
	params []PARAMETER,
	accept core.ValidationAcceptor,
) {
	stepName := ""
	if inst, ok := owner.(SIMPLE_ENTITY_INSTANCE); ok {
		stepName, _ = instanceStepName(inst)
	}
	for i, param := range params {
		if i >= len(def.Attrs) {
			break
		}
		validateParameterValue(ctx, owner, stepName, entityAttrNameAt(stepName, i), def.Attrs[i], param, i, accept)
	}
}

func (k ParamKind) String() string {
	switch k {
	case ParamInteger:
		return "INTEGER"
	case ParamReal:
		return "REAL"
	case ParamString:
		return "STRING"
	case ParamBinary:
		return "BINARY"
	case ParamBoolean:
		return "BOOLEAN"
	case ParamLogical:
		return "LOGICAL"
	case ParamNumber:
		return "NUMBER"
	case ParamEntity:
		return "ENTITY"
	case ParamSelect:
		return "SELECT"
	case ParamEnum:
		return "ENUM"
	case ParamAggregate:
		return "AGGREGATE"
	default:
		return "UNKNOWN"
	}
}
