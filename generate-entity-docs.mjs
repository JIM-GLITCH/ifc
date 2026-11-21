import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const root = path.dirname(fileURLToPath(import.meta.url));
const expressPath = path.join(root, '..', 'EXPRESS', 't.exp');
const outputPath = path.join(root, 'entity_docs_gen.go');
const source = fs.readFileSync(expressPath, 'utf8');

const declarations = new Map();
const pattern = /^ENTITY\s+(\w+)([\s\S]*?)^END_ENTITY;/gm;
for (const match of source.matchAll(pattern)) {
  const name = match[1].toUpperCase();
  const declaration = match[0].trim();
  if (declaration.includes('`')) {
    throw new Error(`EXPRESS declaration for ${name} contains a backtick`);
  }
  declarations.set(name, declaration);
}

if (declarations.size === 0) {
  throw new Error(`no ENTITY declarations found in ${expressPath}`);
}

const names = [...declarations.keys()].sort();
const lines = [
  '// Code generated from ../EXPRESS/t.exp. DO NOT EDIT.',
  '',
  'package ifc',
  '',
  'var entityExpressDocs = map[string]string{',
];
for (const name of names) {
  lines.push(`\t${JSON.stringify(name)}: \`${declarations.get(name)}\`,`);
}
lines.push('}', '');

fs.writeFileSync(outputPath, lines.join('\n'), 'utf8');
console.log(`wrote ${outputPath} (${names.length} entity declarations)`);
