// Copyright 2025 TypeFox GmbH
// This program and the accompanying materials are made available under the
// terms of the MIT License, which is available in the project root.

import * as vscode from 'vscode';
import type { LanguageClientOptions, ServerOptions } from 'vscode-languageclient/node.js';
import * as path from 'node:path';
import { DidChangeConfigurationNotification, LanguageClient, TransportKind } from 'vscode-languageclient/node.js';
import { buildInitializationOptions } from './validationLayers';
import { registerValidationStatusBar } from './statusBar';

let client: LanguageClient;

export async function activate(context: vscode.ExtensionContext): Promise<void> {
    registerValidationStatusBar(context);
    context.subscriptions.push(
        vscode.commands.registerCommand('ifc.openEntityReference', async (target: unknown) => {
            const payload = Array.isArray(target) ? target[0] : target;
            if (payload == null || typeof payload !== 'object') {
                return;
            }
            const { uri, range } = payload as {
                uri?: string;
                range?: {
                    start?: { line: number; character: number };
                    end?: { line: number; character: number };
                };
            };
            if (!uri) {
                return;
            }
            const document = await vscode.workspace.openTextDocument(vscode.Uri.parse(uri));
            const editor = await vscode.window.showTextDocument(document);
            if (range?.start && range?.end) {
                const selection = new vscode.Selection(
                    new vscode.Position(range.start.line, range.start.character),
                    new vscode.Position(range.end.line, range.end.character)
                );
                editor.selection = selection;
                editor.revealRange(selection);
            }
        })
    );
    client = await startLanguageClient(context);
    context.subscriptions.push(
        vscode.workspace.onDidChangeConfiguration((event) => {
            if (event.affectsConfiguration('ifc.validation')) {
                void pushValidationSettingsToServer(client);
            }
        })
    );
}

export function deactivate(): Thenable<void> | undefined {
    if (client) {
        return client.stop();
    }
    return undefined;
}

function pushValidationSettingsToServer(lc: LanguageClient): Thenable<void> {
    // Send the same shape as initializationOptions; the server also accepts the
    // vscode-languageclient wrapper { ifc: { validation: ... } }.
    return lc.sendNotification(DidChangeConfigurationNotification.type, {
        settings: buildInitializationOptions()
    });
}

async function startLanguageClient(context: vscode.ExtensionContext): Promise<LanguageClient> {
    const serverOptions: ServerOptions = {
        command: context.asAbsolutePath(path.join('dist', 'server' + (process.platform === 'win32' ? '.exe' : ''))),
        args: [],
        transport: TransportKind.stdio
    };

    const clientOptions: LanguageClientOptions = {
        documentSelector: [{ scheme: 'file', language: 'ifc' }],
        outputChannel: vscode.window.createOutputChannel('Ifc Language Server'),
        initializationOptions: buildInitializationOptions(),
        markdown: {
            isTrusted: { enabledCommands: ['ifc.openEntityReference'] },
            supportHtml: false
        },
        synchronize: {
            configurationSection: 'ifc'
        }
    };
    const lc = new LanguageClient(
        'ifc',
        'Ifc',
        serverOptions,
        clientOptions
    );
    await lc.start();
    return lc;
}
