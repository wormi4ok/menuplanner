import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const dirname = path.dirname(fileURLToPath(import.meta.url));

const { version } = JSON.parse(
  fs.readFileSync(path.join(dirname, '..', 'package.json'), 'utf8'),
);

const config = {
  MP_VERSION: process.env.MP_VERSION || version,
  MP_CLIENT_ID: process.env.MP_CLIENT_ID || '',
  API_ADDRESS: process.env.API_ADDRESS || 'http://localhost:8081',
};

const dist = path.join(dirname, '..', 'dist');

fs.writeFileSync(path.join(dist, 'config.js'), `window.config = ${JSON.stringify(config, null, 2)};\n`);
fs.writeFileSync(path.join(dist, 'version.txt'), `${config.MP_VERSION}\n`);
