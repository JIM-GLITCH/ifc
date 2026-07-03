package ifc

import (
	"context"
	"sync"

	"typefox.dev/fastbelt/util/service"
)

// ValidationLayer identifies a sub-check mapped to one of three UI tiers.
type ValidationLayer string

const (
	LayerStepEncoding    ValidationLayer = "stepEncoding"
	LayerAttributeTypes  ValidationLayer = "attributeTypes"
	LayerTypeWhere       ValidationLayer = "typeWhere"
	LayerDerived         ValidationLayer = "derived"
	LayerEntityWhere     ValidationLayer = "entityWhere"
	LayerGlobalRule      ValidationLayer = "globalRule"
	LayerUnique          ValidationLayer = "unique"
	LayerAggregateBounds ValidationLayer = "aggregateBounds"
)

// ValidationSettings toggles the three validation tiers exposed in the extension UI.
// Parse and linker diagnostics are always reported regardless of these flags.
type ValidationSettings struct {
	Structure bool `json:"structure"`
	Types     bool `json:"types"`
	Rules     bool `json:"rules"`
}

func DefaultValidationSettings() ValidationSettings {
	return ValidationSettings{
		Structure: true,
		Types:     true,
		Rules:     true,
	}
}

func (s ValidationSettings) EnabledCount() int {
	n := 0
	if s.Structure {
		n++
	}
	if s.Types {
		n++
	}
	if s.Rules {
		n++
	}
	return n
}

func (s ValidationSettings) LayerEnabled(layer ValidationLayer) bool {
	switch layer {
	case LayerStepEncoding, LayerAttributeTypes, LayerAggregateBounds:
		return s.Structure
	case LayerTypeWhere, LayerDerived:
		return s.Types
	case LayerEntityWhere, LayerGlobalRule, LayerUnique:
		return s.Rules
	default:
		return true
	}
}

// ValidationSettingsStore holds the active settings for the language server process.
type ValidationSettingsStore struct {
	mu   sync.RWMutex
	data ValidationSettings
}

func NewValidationSettingsStore() *ValidationSettingsStore {
	return &ValidationSettingsStore{data: DefaultValidationSettings()}
}

func (store *ValidationSettingsStore) Get() ValidationSettings {
	if store == nil {
		return DefaultValidationSettings()
	}
	store.mu.RLock()
	defer store.mu.RUnlock()
	return store.data
}

func (store *ValidationSettingsStore) Set(settings ValidationSettings) {
	if store == nil {
		return
	}
	store.mu.Lock()
	store.data = settings
	store.mu.Unlock()
}

func ValidationSettingsFromService(sc *service.Container) ValidationSettings {
	if sc == nil {
		return DefaultValidationSettings()
	}
	store, err := service.Get[*ValidationSettingsStore](sc)
	if err != nil || store == nil {
		return DefaultValidationSettings()
	}
	return store.Get()
}

type validationSettingsContextKey struct{}

func WithValidationSettings(ctx context.Context, settings ValidationSettings) context.Context {
	return context.WithValue(ctx, validationSettingsContextKey{}, settings)
}

func ValidationSettingsFromContext(ctx context.Context) ValidationSettings {
	if ctx == nil {
		return DefaultValidationSettings()
	}
	settings, ok := ctx.Value(validationSettingsContextKey{}).(ValidationSettings)
	if !ok {
		return DefaultValidationSettings()
	}
	return settings
}

func validationEnabled(ctx context.Context, layer ValidationLayer) bool {
	return ValidationSettingsFromContext(ctx).LayerEnabled(layer)
}
