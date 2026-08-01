import { test as base, expect } from '@playwright/test';
import { spawn } from 'node:child_process';
import { once } from 'node:events';
import { createServer } from 'node:net';
import { apiBinary } from './global-setup.js';

const FROZEN_NOW = new Date('2026-08-05T10:00:00Z');

async function freePort() {
  const server = createServer();
  server.listen(0, '127.0.0.1');
  await once(server, 'listening');
  const { port } = server.address();
  server.close();
  await once(server, 'close');
  return port;
}

async function waitForHealth(url, api, deadline = 10000) {
  const giveUp = Date.now() + deadline;
  for (;;) {
    if (api.exitCode !== null) throw new Error(`API exited with code ${api.exitCode}`);
    try {
      const response = await fetch(`${url}/health`);
      if (response.ok) return;
    } catch {
      // not listening yet
    }
    if (Date.now() > giveUp) throw new Error(`API did not become healthy at ${url}`);
    await new Promise((r) => { setTimeout(r, 50); });
  }
}

export const test = base.extend({
  apiURL: async ({}, use) => {
    const port = await freePort();
    const url = `http://127.0.0.1:${port}`;
    const api = spawn(apiBinary, [], {
      stdio: 'ignore',
      env: {
        ...process.env,
        MP_HOST: '127.0.0.1',
        MP_PORT: String(port),
        MP_JWT_SECRET: 'test-secret',
        MP_MYSQL_DSN: '',
        MP_CLIENT_ID: '',
        MP_CLIENT_SECRET: '',
      },
    });

    await waitForHealth(url, api);
    await use(url);

    api.kill('SIGKILL');
    await once(api, 'exit');
  },

  page: async ({ page, apiURL }, use) => {
    await page.route('**/config.js', (route) => route.fulfill({
      contentType: 'application/javascript',
      body: `window.config = {
        MP_VERSION: 'test',
        MP_CLIENT_ID: '',
        API_ADDRESS: '${apiURL}',
      };`,
    }));
    await page.clock.setFixedTime(FROZEN_NOW);
    await use(page);
  },
});

export function signInForm(page) {
  return page.locator('.tab-item:visible');
}

export async function login(page, email = 'demo@demo.com', password = 'demo') {
  await page.goto('/login');
  const form = signInForm(page);
  await form.locator('input[type=email]').fill(email);
  await form.locator('input[type=password]').fill(password);
  await page.getByRole('button', { name: 'Continue' }).click();
}

export async function loginAsDemo(page) {
  await login(page);
  await expect(page.locator('.navbar-brand')).toBeVisible();
}

const CELLS = '.week-grid > *';

export function daySlot(page, day, slot) {
  return page.locator(CELLS).nth(7 + slot * 7 + day);
}

export function recipeName(slot) {
  return slot.locator('.card-header-title');
}

export async function fillSlot(slot, name) {
  await slot.getByRole('button').click();
  await slot.locator('input[type=search]').fill(name);
  await slot.locator('.dropdown-item', { hasText: name }).first().click();
}

export async function confirmDialog(page) {
  await page.locator('.dialog').getByRole('button', { name: 'OK' }).click();
}

export { expect, FROZEN_NOW };
