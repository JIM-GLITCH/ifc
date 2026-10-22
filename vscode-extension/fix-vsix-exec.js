const yauzl = require('yauzl');
const yazl = require('yazl');
const fs = require('node:fs');

// VSIXs built on Windows carry no Unix exec bit, so the bundled language
// server binary fails with EACCES when extracted on Linux/macOS. This rewrites
// the VSIX (a zip) and forces the executable bit on the server binary.
const vsixPath = process.argv[2];
if (!vsixPath) {
	console.error('usage: node fix-vsix-exec.js <file.vsix>');
	process.exit(1);
}

const TARGET = 'extension/dist/server';
const EXEC_MODE = 0o100755;

function openReadStream(zipfile, entry) {
	return new Promise((resolve, reject) => {
		zipfile.openReadStream(entry, (err, rs) => (err ? reject(err) : resolve(rs)));
	});
}

async function run() {
	const zipfile = await new Promise((resolve, reject) => {
		yauzl.open(vsixPath, { lazyEntries: true }, (err, z) => (err ? reject(err) : resolve(z)));
	});
	const tmp = vsixPath + '.tmp';
	const out = fs.createWriteStream(tmp);
	const zip = new yazl.ZipFile();
	zip.outputStream.pipe(out);

	await new Promise((resolve, reject) => {
		let active = 0;
		const finalize = () => {
			if (active === 0) {
				zip.end();
				out.on('close', resolve);
			} else {
				setTimeout(finalize, 25);
			}
		};
		out.on('error', reject);
		zipfile.on('entry', (entry) => {
			try {
				if (/\/$/.test(entry.fileName)) {
					zip.addEmptyDirectory(entry.fileName, { mode: entry.mode || 0o040755 });
					zipfile.readEntry();
					return;
				}
				active++;
				openReadStream(zipfile, entry)
					.then((rs) => {
						const chunks = [];
						rs.on('data', (c) => chunks.push(c));
						rs.on('end', () => {
							const buf = Buffer.concat(chunks);
							const mode = entry.fileName === TARGET ? EXEC_MODE : entry.mode || 0o100644;
							zip.addBuffer(buf, entry.fileName, { mode, mtime: entry.mtime, compress: true });
							active--;
							zipfile.readEntry();
						});
						rs.on('error', reject);
					})
					.catch(reject);
			} catch (e) {
				reject(e);
			}
		});
		zipfile.on('end', finalize);
		zipfile.readEntry();
	});

	fs.renameSync(tmp, vsixPath);
	console.log(`[fix-vsix-exec] ensured +x on "${TARGET}" in ${vsixPath}`);
}

run().catch((e) => {
	console.error(e);
	process.exit(1);
});
