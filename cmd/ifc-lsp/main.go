// Copyright 2025 TypeFox GmbH
// This program and the accompanying materials are made available under the
// terms of the MIT License, which is available in the project root.

package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	lang "ifc"

	"typefox.dev/fastbelt/server"
	"typefox.dev/fastbelt/util/service"
)

// CreateServices creates a service container for the Ifc language to be used in the language server.
func CreateServices() *service.Container {
	sc := service.NewContainer()
	lang.SetupServices(sc)
	server.SetupDefaultServices(sc)
	lang.SetupIfcLanguageServer(sc)
	server.SetupStdioServices(sc)
	sc.Seal()
	return sc
}

func main() {
	if addr := os.Getenv("IFC_PPROF"); addr != "" {
		go func() {
			log.Printf("pprof listening on http://%s/debug/pprof/", addr)
			if err := http.ListenAndServe(addr, nil); err != nil {
				log.Printf("pprof server stopped: %v", err)
			}
		}()
	}
	ctx := context.Background()
	sc := CreateServices()

	if err := server.StartLanguageServer(ctx, sc); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
