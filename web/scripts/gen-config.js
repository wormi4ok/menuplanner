import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const dirname = path.dirname(fileURLToPath(import.meta.url));

const config = {
  MP_VERSION: process.env.MP_VERSION
    || (process.env.CF_PAGES_COMMIT_SHA || 'dev').slice(0, 7),
  MP_CLIENT_ID: process.env.MP_CLIENT_ID || '',
  API_ADDRESS: process.env.API_ADDRESS || 'http://localhost:8081',
};

fs.writeFileSync(
  path.join(dirname, '..', 'dist', 'config.js'),
  `window.config = ${JSON.stringify(config, null, 2)};\n`,
);
