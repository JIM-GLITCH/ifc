const { execFileSync } = require('node:child_process');
const fs = require('node:fs');
const path = require('node:path');

// Map an vsce target (<os>-<arch>) to Go build env + output binary name.
const GOOS = { win32: 'windows', linux: 'linux', darwin: 'darwin' };
const GOARCH = { x64: 'amd64', arm64: 'arm64' };

function detectTarget() {
	// Node uses "win32"/"linux"/"darwin" which match vsce target os names.
	return `${process.platform}-${process.arch}`;
}

const target = process.argv[2] || detectTarget();
const [os, arch] = target.split('-');
const goos = GOOS[os];
const goarch = GOARCH[arch];

if (!goos || !goarch) {
	console.error(`[build-server] Unsupported target: ${target}`);
	process.exit(1);
}

const outName = os === 'win32' ? 'server.exe' : 'server';
const outPath = path.join(__dirname, 'dist', outName);
fs.mkdirSync(path.dirname(outPath), { recursive: true });

// Remove stale binaries from other targets so each VSIX only carries its own.
for (const stale of ['server', 'server.exe']) {
	const p = path.join(__dirname, 'dist', stale);
	if (p !== outPath && fs.existsSync(p)) {
		fs.rmSync(p);
	}
}

const env = { ...process.env, GOOS: goos, GOARCH: goarch, CGO_ENABLED: '0' };
const args = ['build', '-o', outPath, '../cmd/ifc-lsp'];

console.log(`[build-server] ${target} -> ${outPath} (GOOS=${goos} GOARCH=${goarch})`);
execFileSync('go', args, { env, stdio: 'inherit' });
