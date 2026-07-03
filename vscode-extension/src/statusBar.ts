// Copyright 2025 TypeFox GmbH
// This program and the accompanying materials are made available under the
// terms of the MIT License, which is available in the project root.

import * as vscode from 'vscode';
import {
    VALIDATION_LAYERS,
    countEnabledLayers,
    readValidationSettings,
    setAllLayersEnabled,
    setLayerEnabled,
    type ValidationLayerSpec
} from './validationLayers';

const STATUS_BAR_ID = 'ifc.validationLayers';

export class ValidationStatusBar {
    private readonly item: vscode.StatusBarItem;
    private readonly disposables: vscode.Disposable[] = [];

    constructor() {
        this.item = vscode.window.createStatusBarItem(vscode.StatusBarAlignment.Right, 90);
        this.item.command = 'ifc.showValidationLayers';
        this.disposables.push(this.item);
        this.disposables.push(
            vscode.window.onDidChangeActiveTextEditor(() => this.refresh()),
            vscode.workspace.onDidChangeConfiguration((event) => {
                if (event.affectsConfiguration('ifc.validation')) {
                    this.refresh();
                }
            })
        );
        this.refresh();
    }

    dispose(): void {
        this.item.dispose();
        for (const disposable of this.disposables) {
            disposable.dispose();
        }
    }

    refresh(): void {
        const editor = vscode.window.activeTextEditor;
        if (!editor || editor.document.languageId !== 'ifc') {
            this.item.hide();
            return;
        }
        const settings = readValidationSettings();
        const enabled = countEnabledLayers(settings);
        const total = VALIDATION_LAYERS.length;
        this.item.text = `$(list-selection) IFC ${enabled}/${total}`;
        this.item.tooltip = `IFC validation layers (${enabled}/${total} enabled)\nClick to toggle layers`;
        this.item.show();
    }
}

function layerQuickPickItem(layer: ValidationLayerSpec, enabled: boolean): vscode.QuickPickItem & { layer?: ValidationLayerSpec; action?: 'all-on' | 'all-off' } {
    return {
        label: `${enabled ? '$(check)' : '$(circle-large-outline)'} ${layer.label}`,
        description: layer.description,
        detail: enabled ? 'Enabled' : 'Disabled',
        layer,
        picked: enabled
    };
}

export async function showValidationLayersQuickPick(): Promise<void> {
    while (true) {
        const settings = readValidationSettings();
        const enabled = countEnabledLayers(settings);
        const items: Array<vscode.QuickPickItem & { layer?: ValidationLayerSpec; action?: 'all-on' | 'all-off' }> = [
            { label: '$(check-all) Enable all layers', action: 'all-on' },
            { label: '$(circle-slash) Disable all layers', action: 'all-off' },
            { label: '', kind: vscode.QuickPickItemKind.Separator },
            ...VALIDATION_LAYERS.map((layer) => layerQuickPickItem(layer, settings[layer.id] !== false))
        ];
        const pick = await vscode.window.showQuickPick(items, {
            title: 'IFC validation layers',
            placeHolder: `${enabled}/${VALIDATION_LAYERS.length} layers enabled — choose a layer to toggle (Esc to close)`
        });
        if (!pick) {
            return;
        }
        if (pick.action === 'all-on') {
            await setAllLayersEnabled(true);
            continue;
        }
        if (pick.action === 'all-off') {
            await setAllLayersEnabled(false);
            continue;
        }
        if (pick.layer) {
            await setLayerEnabled(pick.layer, settings[pick.layer.id] === false);
        }
    }
}

export function registerValidationStatusBar(context: vscode.ExtensionContext): ValidationStatusBar {
    const statusBar = new ValidationStatusBar();
    context.subscriptions.push(statusBar);
    context.subscriptions.push(
        vscode.commands.registerCommand('ifc.showValidationLayers', () => showValidationLayersQuickPick())
    );
    return statusBar;
}
