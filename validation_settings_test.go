package ifc

import "testing"

func TestParseValidationSettingsDefaults(t *testing.T) {
	settings := ParseValidationSettings(nil)
	if settings != DefaultValidationSettings() {
		t.Fatalf("expected defaults, got %+v", settings)
	}
}

func TestParseValidationSettingsMap(t *testing.T) {
	settings := ParseValidationSettings(map[string]any{
		"validation": map[string]any{
			"rules": false,
		},
	})
	if settings.Rules {
		t.Fatalf("expected rules disabled, got %+v", settings)
	}
	if !settings.Structure || !settings.Types {
		t.Fatalf("expected structure/types enabled, got %+v", settings)
	}
}

func TestParseValidationSettingsIfcWrapper(t *testing.T) {
	settings := ParseValidationSettings(map[string]any{
		"ifc": map[string]any{
			"validation": map[string]any{
				"rules": false,
				"types": false,
			},
		},
	})
	if settings.Rules || settings.Types {
		t.Fatalf("expected rules/types disabled from ifc wrapper, got %+v", settings)
	}
	if !settings.Structure {
		t.Fatalf("expected structure enabled, got %+v", settings)
	}
}

func TestParseValidationSettingsAllDisabled(t *testing.T) {
	settings := ParseValidationSettings(map[string]any{
		"validation": map[string]any{
			"structure": false,
			"types":     false,
			"rules":     false,
		},
	})
	if settings.Structure || settings.Types || settings.Rules {
		t.Fatalf("expected all tiers disabled, got %+v", settings)
	}
}

func TestParseValidationSettingsLegacyGranular(t *testing.T) {
	settings := ParseValidationSettings(map[string]any{
		"entityWhere": false,
	})
	if settings.Rules {
		t.Fatalf("expected legacy entityWhere=false to disable rules tier, got %+v", settings)
	}
}

func TestValidationSettingsEnabledCount(t *testing.T) {
	settings := DefaultValidationSettings()
	settings.Rules = false
	if settings.EnabledCount() != 2 {
		t.Fatalf("expected 2 enabled layers, got %d", settings.EnabledCount())
	}
}
