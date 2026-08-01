import { execFileSync } from 'node:child_process';
import { mkdirSync } from 'node:fs';
import { dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const here = dirname(fileURLToPath(import.meta.url));

export const repoRoot = resolve(here, '..', '..');
export const apiBinary = process.env.TEST_API_BINARY || resolve(here, '.bin', 'menuplanner');

export default function globalSetup() {
  if (process.env.TEST_API_BINARY) return;

  mkdirSync(dirname(apiBinary), { recursive: true });
  execFileSync('go', ['build', '-buildvcs=false', '-o', apiBinary, '.'], {
    cwd: repoRoot,
    stdio: 'inherit',
  });
}
