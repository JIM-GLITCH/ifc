// Copyright 2025 TypeFox GmbH
// This program and the accompanying materials are made available under the
// terms of the MIT License, which is available in the project root.

import * as vscode from 'vscode';

export interface ValidationLayerSpec {
    id: string;
    label: string;
    description: string;
    configKey: string;
}

export const VALIDATION_LAYERS: ValidationLayerSpec[] = [
    {
        id: 'structure',
        label: 'Structure',
        description: 'STEP encoding, attribute types, aggregate bounds, entity keywords',
        configKey: 'validation.structure'
    },
    {
        id: 'types',
        label: 'Types',
        description: 'Type WHERE rules and DERIVE evaluation for * attributes',
        configKey: 'validation.types'
    },
    {
        id: 'rules',
        label: 'Rules',
        description: 'Entity WHERE, Global RULE, and UNIQUE constraints',
        configKey: 'validation.rules'
    }
];

export function getValidationConfig(): vscode.WorkspaceConfiguration {
    return vscode.workspace.getConfiguration('ifc');
}

export function readValidationSettings(): Record<string, boolean> {
    const config = getValidationConfig();
    const settings: Record<string, boolean> = {};
    for (const layer of VALIDATION_LAYERS) {
        settings[layer.id] = config.get<boolean>(layer.configKey, true);
    }
    return settings;
}

export function countEnabledLayers(settings: Record<string, boolean>): number {
    return VALIDATION_LAYERS.filter((layer) => settings[layer.id] !== false).length;
}

export async function setLayerEnabled(layer: ValidationLayerSpec, enabled: boolean): Promise<void> {
    const target = vscode.workspace.workspaceFolders?.length
        ? vscode.ConfigurationTarget.Workspace
        : vscode.ConfigurationTarget.Global;
    await getValidationConfig().update(layer.configKey, enabled, target);
}

export async function setAllLayersEnabled(enabled: boolean): Promise<void> {
    const target = vscode.workspace.workspaceFolders?.length
        ? vscode.ConfigurationTarget.Workspace
        : vscode.ConfigurationTarget.Global;
    const config = getValidationConfig();
    for (const layer of VALIDATION_LAYERS) {
        await config.update(layer.configKey, enabled, target);
    }
}

export function buildInitializationOptions(): { validation: Record<string, boolean> } {
    return { validation: readValidationSettings() };
}
