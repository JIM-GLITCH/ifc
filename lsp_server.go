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
	if !sc.Has[*ValidationSettingsStore]() {
		sc.Put(NewValidationSettingsStore())
	}
	sc.Override[server.FoldingRangeProvider](&noopFoldingProvider{})
	sc.Override[lsp.Server](NewIfcLanguageServer(sc))
	sc.Override[server.HoverProvider](NewIfcHoverProvider(sc))
	sc.Override[server.DocumentSymbolProvider](NewIfcDocumentSymbolProvider(sc))
	sc.Override[server.CompletionProvider](NewIfcCompletionProvider(sc))
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
		params.Documentation = &lsp.CompletionItemDocumentation{
			MarkupContent: &lsp.MarkupContent{Kind: lsp.Markdown, Value: doc},
		}
	}
	return params, nil
}

func (s *IfcLanguageServer) SignatureHelp(ctx context.Context, params *lsp.SignatureHelpParams) (*lsp.SignatureHelp, error) {
	lock, err := s.sc.Get[workspace.Lock]()
	if err != nil {
		return nil, err
	}
	var result *lsp.SignatureHelp
	var providerErr error
	// Signature help only reads the token stream: wait for the lexed state of
	// this document instead of the whole workspace build.
	if err := lock.ReadAt(ctx, core.DocStateLexed, []core.URI{core.ParseURI(string(params.TextDocument.URI))}, func(ctx context.Context) {
		result, providerErr = s.signatureHelp.HandleSignatureHelpRequest(ctx, params)
	}); err != nil {
		return nil, err
	}
	return result, providerErr
}

func (s *IfcLanguageServer) InlayHint(ctx context.Context, params *lsp.InlayHintParams) ([]lsp.InlayHint, error) {
	lock, err := s.sc.Get[workspace.Lock]()
	if err != nil {
		return nil, err
	}
	var result []lsp.InlayHint
	var providerErr error
	if err := lock.ReadAt(ctx, core.DocStateLexed, []core.URI{core.ParseURI(string(params.TextDocument.URI))}, func(ctx context.Context) {
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
		lock := sc.MustGet[workspace.Lock]()
		builder := sc.MustGet[workspace.Builder]()
		docManager := sc.MustGet[workspace.DocumentManager]()
		lock.Write(context.Background(), func(ctx context.Context) {
			// Keep every build state up to References so only validation re-runs.
			// Reset cascades: the kept mask must be a prefix of the state order,
			// otherwise fields are cleared while their state bits survive.
			keepState := core.DocStateComplete.Without(core.DocStateValidated)
			docs := make([]*core.Document, 0, 8)
			for doc := range docManager.All() {
				builder.Reset(doc, keepState)
				docs = append(docs, doc)
			}
			if len(docs) == 0 {
				return
			}
			_ = builder.Build(ctx, docs)
		})
	}()
}
