package ifc

import (
	"encoding/json"

	"typefox.dev/fastbelt/util/service"
)

func ParseValidationSettings(raw any) ValidationSettings {
	if raw == nil {
		return DefaultValidationSettings()
	}
	flat := flattenValidationSection(raw)
	if flat == nil {
		return DefaultValidationSettings()
	}
	return settingsFromValidationMap(flat)
}

func flattenValidationSection(raw any) map[string]any {
	switch value := raw.(type) {
	case map[string]any:
		return unwrapValidationMap(value)
	default:
		return unwrapValidationMap(extractValidationMap(raw))
	}
}

func unwrapValidationMap(raw map[string]any) map[string]any {
	if raw == nil {
		return nil
	}
	if nested, ok := raw["validation"].(map[string]any); ok {
		return nested
	}
	if ifcSection, ok := raw["ifc"].(map[string]any); ok {
		if nested, ok := ifcSection["validation"].(map[string]any); ok {
			return nested
		}
		return ifcSection
	}
	return raw
}

func extractValidationMap(raw any) map[string]any {
	data, err := json.Marshal(raw)
	if err != nil {
		return nil
	}
	var root map[string]any
	if err := json.Unmarshal(data, &root); err != nil {
		return nil
	}
	return root
}

func settingsFromValidationMap(raw map[string]any) ValidationSettings {
	settings := DefaultValidationSettings()
	if raw == nil {
		return settings
	}
	if v, ok := raw["structure"].(bool); ok {
		settings.Structure = v
	}
	if v, ok := raw["types"].(bool); ok {
		settings.Types = v
	}
	if v, ok := raw["rules"].(bool); ok {
		settings.Rules = v
	}
	applyLegacyGranularKeys(raw, &settings)
	return settings
}

func applyLegacyGranularKeys(raw map[string]any, settings *ValidationSettings) {
	if settings == nil {
		return
	}
	if _, has := raw["structure"]; has {
		return
	}
	if _, has := raw["types"]; has {
		return
	}
	if _, has := raw["rules"]; has {
		return
	}
	if !hasAnyLegacyValidationKey(raw) {
		return
	}
	settings.Structure = legacyLayerEnabled(raw,
		"stepEncoding", "attributeTypes", "aggregateBounds")
	settings.Types = legacyLayerEnabled(raw, "typeWhere", "derived")
	settings.Rules = legacyLayerEnabled(raw, "entityWhere", "globalRule", "unique")
}

func hasAnyLegacyValidationKey(raw map[string]any) bool {
	for _, key := range []string{
		"stepEncoding", "attributeTypes", "typeWhere", "derived",
		"entityWhere", "globalRule", "unique", "aggregateBounds",
	} {
		if _, ok := raw[key]; ok {
			return true
		}
	}
	return false
}

func legacyLayerEnabled(raw map[string]any, keys ...string) bool {
	for _, key := range keys {
		if v, ok := raw[key].(bool); ok && !v {
			return false
		}
	}
	return true
}

func ApplyValidationSettings(sc *service.Container, raw any) ValidationSettings {
	settings := ParseValidationSettings(raw)
	if store, err := service.Get[*ValidationSettingsStore](sc); err == nil && store != nil {
		store.Set(settings)
	}
	return settings
}
