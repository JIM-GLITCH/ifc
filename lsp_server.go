package ifc

import (
	"context"

	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/server"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
	"typefox.dev/lsp"
)

// IfcLanguageServer extends the default fastbelt server with IFC validation settings.
type IfcLanguageServer struct {
	*server.DefaultLanguageServer
	sc            *service.Container
	signatureHelp *IfcSignatureHelpProvider
	inlayHint     *IfcInlayHintProvider
}

func NewIfcLanguageServer(sc *service.Container) lsp.Server {
	base, ok := server.NewDefaultLanguageServer(sc).(*server.DefaultLanguageServer)
	if !ok {
		return server.NewDefaultLanguageServer(sc)
	}
	return &IfcLanguageServer{
		DefaultLanguageServer: base,
		sc:                    sc,
		signatureHelp:         NewIfcSignatureHelpProvider(sc),
		inlayHint:             NewIfcInlayHintProvider(sc),
	}
}

func SetupIfcLanguageServer(sc *service.Container) {
	if !service.Has[*ValidationSettingsStore](sc) {
		service.Put(sc, NewValidationSettingsStore())
	}
	service.Override[lsp.Server](sc, NewIfcLanguageServer(sc))
	service.Override[server.HoverProvider](sc, NewIfcHoverProvider(sc))
	service.Override[server.DocumentSymbolProvider](sc, NewIfcDocumentSymbolProvider(sc))
	service.Override[server.CompletionProvider](sc, NewIfcCompletionProvider(sc))
}

func (s *IfcLanguageServer) Initialize(ctx context.Context, params *lsp.ParamInitialize) (*lsp.InitializeResult, error) {
	if params != nil && params.InitializationOptions != nil {
		ApplyValidationSettings(s.sc, params.InitializationOptions)
	}
	result, err := s.DefaultLanguageServer.Initialize(ctx, params)
	if err != nil {
		return nil, err
	}
	if result != nil {
		if result.Capabilities.CompletionProvider == nil {
			result.Capabilities.CompletionProvider = &lsp.CompletionOptions{}
		}
		result.Capabilities.CompletionProvider.ResolveProvider = true
		result.Capabilities.SignatureHelpProvider = &lsp.SignatureHelpOptions{
			TriggerCharacters:   []string{"(", ","},
			RetriggerCharacters: []string{","},
		}
		result.Capabilities.InlayHintProvider = &lsp.InlayHintRegistrationOptions{}
	}
	return result, nil
}

// ResolveCompletionItem attaches the entity documentation lazily so that
// completion responses stay small; see ifc_completion.go.
func (s *IfcLanguageServer) ResolveCompletionItem(ctx context.Context, params *lsp.CompletionItem) (*lsp.CompletionItem, error) {
	if params == nil {
		return nil, nil
	}
	if doc := ifcEntityCompletionDoc(params.Label); doc != "" {
		params.Documentation = &lsp.Or_CompletionItem_documentation{
			Value: lsp.MarkupContent{Kind: lsp.Markdown, Value: doc},
		}
	}
	return params, nil
}

func (s *IfcLanguageServer) SignatureHelp(ctx context.Context, params *lsp.SignatureHelpParams) (*lsp.SignatureHelp, error) {
	lock, err := service.Get[workspace.Lock](s.sc)
	if err != nil {
		return nil, err
	}
	var result *lsp.SignatureHelp
	var providerErr error
	if err := lock.Read(ctx, func(ctx context.Context) {
		result, providerErr = s.signatureHelp.HandleSignatureHelpRequest(ctx, params)
	}); err != nil {
		return nil, err
	}
	return result, providerErr
}

func (s *IfcLanguageServer) InlayHint(ctx context.Context, params *lsp.InlayHintParams) ([]lsp.InlayHint, error) {
	lock, err := service.Get[workspace.Lock](s.sc)
	if err != nil {
		return nil, err
	}
	var result []lsp.InlayHint
	var providerErr error
	if err := lock.Read(ctx, func(ctx context.Context) {
		result, providerErr = s.inlayHint.HandleInlayHintRequest(ctx, params)
	}); err != nil {
		return nil, err
	}
	return result, providerErr
}

func (s *IfcLanguageServer) DidChangeConfiguration(ctx context.Context, params *lsp.DidChangeConfigurationParams) error {
	if params != nil {
		ApplyValidationSettings(s.sc, params.Settings)
	}
	RebuildAllValidatedDocuments(s.sc)
	return nil
}

// RebuildAllValidatedDocuments re-runs validation for every open document after settings change.
func RebuildAllValidatedDocuments(sc *service.Container) {
	go func() {
		lock := service.MustGet[workspace.Lock](sc)
		builder := service.MustGet[workspace.Builder](sc)
		docManager := service.MustGet[workspace.DocumentManager](sc)
		lock.Write(context.Background(), func(ctx context.Context, downgrade func()) {
			keepState := core.DocStateParsed | core.DocStateExportedSymbols | core.DocStateImportedSymbols |
				core.DocStateLocalSymbols | core.DocStateLinked | core.DocStateReferences
			docs := make([]*core.Document, 0, 8)
			for doc := range docManager.All() {
				builder.Reset(doc, keepState)
				docs = append(docs, doc)
			}
			if len(docs) == 0 {
				return
			}
			_ = builder.Build(ctx, docs, downgrade)
		})
	}()
}
