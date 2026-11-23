// Copyright 2025 TypeFox GmbH
// This program and the accompanying materials are made available under the
// terms of the MIT License, which is available in the project root.

package ifc

import (
	core "typefox.dev/fastbelt"
	"typefox.dev/fastbelt/linking"
	"typefox.dev/fastbelt/textdoc"
	"typefox.dev/fastbelt/util/service"
	"typefox.dev/fastbelt/workspace"
)

// SetupServices sets up the base services for the Ifc language.
func SetupServices(sc *service.Container) {
	sc.Put[core.LanguageSelector](
		workspace.NewDefaultLanguageSelector(
			sc,
			workspace.NewDocumentMatcherWithPatterns("ifc", "**/*.ifc"),
		),
	)
	textdoc.SetupDefaultServices(sc)
	linking.SetupDefaultServices(sc)
	workspace.SetupDefaultServices(sc)
	sc.Put(NewValidationSettingsStore())
	sc.Override[workspace.DocumentValidator](NewIfcDocumentValidator(sc))
	// Registered before SetupGeneratedServices so the generated default
	// (linear-scan SeqScope) is skipped; see ifc_scope.go.
	sc.Put[IfcScopeProvider](ifcScopeProvider{})
	SetupGeneratedServices(sc)
}

// CreateServices creates a service container for the Ifc language to be used in the CLI and tests.
func CreateServices() *service.Container {
	sc := service.NewContainer()
	SetupServices(sc)
	sc.Seal()
	return sc
}
