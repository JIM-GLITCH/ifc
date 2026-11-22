package ifc

import (
	"context"
	"fmt"
	"strings"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
	"typefox.dev/lsp"
)

// IfcSignatureHelpProvider implements signature help for IFC entity
// argument lists. While the cursor is inside the parenthesized parameter
// list of an entity instance (e.g. `#1=IFCAXIS2PLACEMENT3D(|)`), it shows
// the entity's attribute signature and highlights the active argument.
type IfcSignatureHelpProvider struct {
	sc *service.Container
}

func NewIfcSignatureHelpProvider(sc *service.Container) *IfcSignatureHelpProvider {
	return &IfcSignatureHelpProvider{sc: sc}
}

func (p *IfcSignatureHelpProvider) HandleSignatureHelpRequest(ctx context.Context, params *lsp.SignatureHelpParams) (*lsp.SignatureHelp, error) {
	documentManager := service.MustGet[workspace.DocumentManager](p.sc)
	doc := documentManager.Get(core.ParseURI(string(params.TextDocument.URI)))
	if doc == nil {
		return nil, nil
	}

	offset := doc.TextDoc.OffsetAt(params.Position)
	call, ok := enclosingEntityCall(doc.Tokens, offset)
	if !ok {
		return nil, nil
	}

	sigParams, ok := entitySignatureParameters(call.keyword)
	if !ok {
		return nil, nil
	}
	count := len(sigParams)

	var label strings.Builder
	label.WriteString(call.keyword)
	label.WriteByte('(')
	parameters := make([]lsp.ParameterInformation, 0, count)
	for i, param := range sigParams {
		if i > 0 {
			label.WriteString(", ")
		}
		paramLabel := param.Name + ": " + entityParameterTypeDisplay(param)
		label.WriteString(paramLabel)
		parameters = append(parameters, lsp.ParameterInformation{
			Label:         paramLabel,
			Documentation: signatureParameterDoc(param),
		})
	}
	label.WriteByte(')')

	signature := lsp.SignatureInformation{
		Label:      label.String(),
		Parameters: parameters,
	}
	if typeDoc := entityParameterTypeMarkdown(sigParams); typeDoc != "" {
		signature.Documentation = &lsp.Or_SignatureInformation_documentation{
			Value: lsp.MarkupContent{
				Kind:  lsp.Markdown,
				Value: typeDoc,
			},
		}
	}

	activeParameter := uint32(call.argIndex)
	if count > 0 && call.argIndex >= count {
		activeParameter = uint32(count - 1)
	}
	return &lsp.SignatureHelp{
		Signatures:      []lsp.SignatureInformation{signature},
		ActiveSignature: 0,
		ActiveParameter: &activeParameter,
	}, nil
}

// entityCallSite describes the entity argument list enclosing a cursor
// offset: the STEP entity keyword and the index of the argument being edited.
type entityCallSite struct {
	keyword  string
	argIndex int
}

// enclosingEntityCall locates the innermost entity instance argument list
// containing offset and counts the top-level commas before offset to derive
// the active argument index. It works purely on the token stream so that it
// also succeeds while the argument list is incomplete or unparseable.
func enclosingEntityCall(tokens core.TokenSlice, offset int) (entityCallSite, bool) {
	// Index of the last token starting before the cursor.
	idx := -1
	lo, hi := 0, len(tokens)
	for lo < hi {
		mid := (lo + hi) / 2
		if int(tokens[mid].TextSegment.Indices.Start) < offset {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	idx = lo - 1
	if idx < 0 {
		return entityCallSite{}, false
	}

	// Scan backwards for the innermost unmatched '('. A '(' that is not
	// preceded by an entity keyword belongs to a nested list or typed
	// parameter; skip it and keep looking for the enclosing record.
	paren := -1
	depth := 0
	for i := idx; i >= 0; i-- {
		switch tokens[i].Image {
		case ")":
			depth++
		case "(":
			if depth == 0 {
				if i > 0 && tokens[i-1].Type == Token_KEYWORD {
					keyword := strings.ToUpper(strings.TrimPrefix(tokens[i-1].Image, "!"))
					if _, ok := LookupEntity(keyword); ok {
						paren = i
					}
				}
				if paren >= 0 {
					goto found
				}
				// Nested '(' (list or typed parameter): treat it as an open
				// paren of the enclosing argument and continue scanning.
			} else {
				depth--
			}
		}
	}
	return entityCallSite{}, false

found:
	// Count top-level commas between the record's '(' and the cursor.
	argIndex := 0
	depth = 0
	for j := paren + 1; j < len(tokens); j++ {
		t := &tokens[j]
		if int(t.TextSegment.Indices.Start) >= offset {
			break
		}
		switch t.Image {
		case "(":
			depth++
		case ")":
			if depth == 0 {
				// The cursor is past the end of the argument list.
				return entityCallSite{}, false
			}
			depth--
		case ",":
			if depth == 0 {
				argIndex++
			}
		}
	}

	return entityCallSite{
		keyword:  strings.ToUpper(strings.TrimPrefix(tokens[paren-1].Image, "!")),
		argIndex: argIndex,
	}, true
}

type entitySignatureParameter struct {
	Name    string
	Attr    SchemaAttr
	HasAttr bool
}

func entitySignatureParameters(stepName string) ([]entitySignatureParameter, bool) {
	def, ok := LookupEntity(stepName)
	if !ok {
		return nil, false
	}
	names := entityAttrNames[stepName]
	count := len(def.Attrs)
	if count == 0 {
		count = def.ParamCount
	}
	if len(names) > count {
		count = len(names)
	}

	params := make([]entitySignatureParameter, 0, count)
	for i := 0; i < count; i++ {
		param := entitySignatureParameter{}
		if i < len(names) {
			param.Name = names[i]
		}
		if i < len(def.Attrs) {
			param.Attr = def.Attrs[i]
			param.HasAttr = true
		}
		if param.Name == "" {
			if param.HasAttr && param.Attr.TypeName != "" {
				param.Name = param.Attr.TypeName
			} else {
				param.Name = fmt.Sprintf("param%d", i+1)
			}
		}
		params = append(params, param)
	}
	return params, true
}

func entityParameterTypeDisplay(param entitySignatureParameter) string {
	if !param.HasAttr {
		return "?"
	}
	if param.Attr.TypeName != "" {
		return param.Attr.TypeName
	}
	if kind := param.Attr.Kind.String(); kind != "UNKNOWN" {
		return kind
	}
	return "?"
}

func entityParameterTypeDisplayWithModifiers(param entitySignatureParameter) string {
	var b strings.Builder
	if param.HasAttr {
		if param.Attr.Optional {
			b.WriteString("OPTIONAL ")
		}
		if param.Attr.Derived {
			b.WriteString("DERIVED ")
		}
	}
	b.WriteString(entityParameterTypeDisplay(param))
	return b.String()
}

func entitySignatureLabel(stepName string, params []entitySignatureParameter, typed bool) string {
	var b strings.Builder
	b.WriteString(stepName)
	b.WriteByte('(')
	for i, param := range params {
		if i > 0 {
			b.WriteString(", ")
		}
		if typed {
			fmt.Fprintf(&b, "%s: %s", param.Name, entityParameterTypeDisplay(param))
		} else {
			b.WriteString(param.Name)
		}
	}
	b.WriteByte(')')
	return b.String()
}

func entityParameterTypeMarkdown(params []entitySignatureParameter) string {
	if len(params) == 0 {
		return ""
	}
	lines := make([]string, 0, len(params))
	for _, param := range params {
		lines = append(lines, fmt.Sprintf("- `%s`: `%s`", param.Name, entityParameterTypeDisplayWithModifiers(param)))
	}
	return strings.Join(lines, "\n")
}

// signatureParameterDoc renders the per-argument documentation shown in the
// signature help popup: the expected type, optionality and whether the
// attribute is a derived attribute that must be passed as '*'.
func signatureParameterDoc(param entitySignatureParameter) string {
	var b strings.Builder
	if typ := entityParameterTypeDisplay(param); typ != "?" {
		fmt.Fprintf(&b, "Type: %s\n", typ)
	}
	attr := param.Attr
	switch {
	case param.HasAttr && attr.Derived:
		b.WriteString("Derived attribute; pass * to omit it.")
	case param.HasAttr && attr.Optional:
		b.WriteString("Optional; use $ for an unset value.")
	default:
		b.WriteString("Required.")
	}
	return b.String()
}
