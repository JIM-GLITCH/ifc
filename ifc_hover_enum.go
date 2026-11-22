package ifc

import (
	"fmt"
	"strings"

	core "typefox.dev/fastbelt"
	"typefox.dev/lsp"
)

func enumHover(first, second *core.Token) *lsp.Hover {
	for _, token := range []*core.Token{first, second} {
		if hover := enumHoverToken(token); hover != nil {
			return hover
		}
	}
	return nil
}

func enumHoverToken(token *core.Token) *lsp.Hover {
	if token == nil || token.Type != Token_ENUMERATION {
		return nil
	}
	lit, ok := token.Element.(ENUMERATION_LITERAL)
	if !ok || lit == nil {
		return nil
	}

	defs, attrName := enumContextForLiteral(lit)
	if len(defs) == 0 {
		return nil
	}

	value := normalizeEnumValue(lit.Value())
	chosen := defs[0]
	for _, def := range defs {
		if enumTypeHasValue(def, value) {
			chosen = def
			break
		}
	}
	decl, ok := enumTypeExpressDocs[strings.ToUpper(chosen.Name)]
	if !ok {
		return nil
	}

	var b strings.Builder
	if attrName != "" {
		fmt.Fprintf(&b, "**%s** : `%s`\n\n", attrName, chosen.Name)
	}
	b.WriteString("```express\n" + decl + "\n```")

	return &lsp.Hover{
		Contents: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: b.String(),
		},
		Range: token.Segment().Range.LspRange(),
	}
}

// enumContextForLiteral resolves the enumeration literal to the enum types
// accepted by the enclosing entity attribute, plus the attribute name.
func enumContextForLiteral(lit ENUMERATION_LITERAL) ([]TypeDef, string) {
	var record SIMPLE_RECORD
	var topParam core.AstNode
	for node := core.AstNode(lit); node != nil; node = node.Container() {
		if r, ok := node.(SIMPLE_RECORD); ok {
			record = r
			break
		}
		if _, ok := node.(PARAMETER); ok {
			topParam = node
		}
	}
	if record == nil || topParam == nil {
		return nil, ""
	}

	paramList := record.Parameters()
	if paramList == nil {
		return nil, ""
	}
	params := paramList.Parameters()
	index := -1
	for i, param := range params {
		if core.AstNode(param) == topParam {
			index = i
			break
		}
	}
	if index < 0 {
		return nil, ""
	}

	stepName := strings.ToUpper(strings.TrimSpace(record.Keyword()))
	def, ok := LookupEntity(stepName)
	if !ok || index >= len(def.Attrs) {
		return nil, ""
	}
	attr := def.Attrs[index]
	return enumTypesForAttr(attr), entityAttrNameAt(stepName, index)
}

// enumTypesForAttr collects the enum type definitions acceptable for the
// given attribute, resolving aggregates, selects and aliases.
func enumTypesForAttr(attr SchemaAttr) []TypeDef {
	var out []TypeDef
	seen := map[string]bool{}
	var add func(name string)
	add = func(name string) {
		name = strings.TrimSpace(name)
		if name == "" || seen[strings.ToUpper(name)] {
			return
		}
		seen[strings.ToUpper(name)] = true
		typeDef, ok := LookupType(canonicalTypeName(name))
		if !ok {
			typeDef, ok = LookupType(name)
		}
		if !ok {
			return
		}
		switch typeDef.ValueKind {
		case ParamEnum:
			out = append(out, typeDef)
		case ParamSelect:
			for _, member := range typeDef.SelectMembers {
				add(member)
			}
		}
	}

	switch attr.Kind {
	case ParamAggregate:
		add(aggregateElementTypeName(attr.TypeName))
	default:
		add(attr.TypeName)
	}
	return out
}

func enumTypeHasValue(def TypeDef, value string) bool {
	for _, expected := range def.EnumValues {
		if normalizeEnumValue(expected) == value {
			return true
		}
	}
	return false
}
