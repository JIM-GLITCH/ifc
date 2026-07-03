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
	sc *service.Container
}

func NewIfcLanguageServer(sc *service.Container) lsp.Server {
	base, ok := server.NewDefaultLanguageServer(sc).(*server.DefaultLanguageServer)
	if !ok {
		return server.NewDefaultLanguageServer(sc)
	}
	return &IfcLanguageServer{
		DefaultLanguageServer: base,
		sc:                    sc,
	}
}

func SetupIfcLanguageServer(sc *service.Container) {
	if !service.Has[*ValidationSettingsStore](sc) {
		service.Put(sc, NewValidationSettingsStore())
	}
	service.Override[lsp.Server](sc, NewIfcLanguageServer(sc))
}

func (s *IfcLanguageServer) Initialize(ctx context.Context, params *lsp.ParamInitialize) (*lsp.InitializeResult, error) {
	if params != nil && params.InitializationOptions != nil {
		ApplyValidationSettings(s.sc, params.InitializationOptions)
	}
	return s.DefaultLanguageServer.Initialize(ctx, params)
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
