import {
  test, expect, login, loginAsDemo, signInForm,
} from './fixtures.js';

test('unauthenticated visit to the week redirects to login', async ({ page }) => {
  await page.goto('/');
  await expect(page).toHaveURL(/\/login$/);
  await expect(page.getByRole('button', { name: 'Continue' })).toBeVisible();
});

test('the demo hint fills in the sign-in form', async ({ page }) => {
  await page.goto('/login');
  await page.locator('.notification a', { hasText: 'demo' }).click();
  await expect(signInForm(page).locator('input[type=email]')).toHaveValue('demo@demo.com');

  await page.getByRole('button', { name: 'Continue' }).click();
  await expect(page).toHaveURL(/:\d+\/$/);
});

test('email login lands on the current week', async ({ page }) => {
  await loginAsDemo(page);
  await expect(page).toHaveURL(/:\d+\/$/);
  await expect(page.getByText('Demo User')).toBeVisible();
  await expect(page.getByText('Spicy Poblano Bolognese').first()).toBeVisible();
});

test('a wrong password is reported and keeps you on login', async ({ page }) => {
  await login(page, 'demo@demo.com', 'nope');
  await expect(page.getByRole('alert')).toContainText('Invalid email or password');
  await expect(page).toHaveURL(/\/login$/);
});

test('signing up gets its own session', async ({ page }) => {
  await page.goto('/login');
  await page.getByRole('tab', { name: 'Sign Up' }).click();

  const form = page.locator('.tab-item').nth(1);
  await form.locator('input[type=email]').fill('newcomer@example.com');
  await form.locator('input[type=password]').first().fill('hunter2');
  await form.locator('input[type=password]').last().fill('hunter2');
  await page.getByRole('button', { name: 'Register' }).click();

  await expect(page).toHaveURL(/:\d+\/$/);
});

test('logout ends the session and re-protects the app', async ({ page }) => {
  await loginAsDemo(page);
  await page.locator('.navbar-link').click();
  await page.getByText('Logout').click();

  await expect(page).toHaveURL(/\/login$/);
  await expect(page.getByText('Spicy Poblano Bolognese')).toHaveCount(0);
  expect(await page.evaluate(() => localStorage.getItem('token'))).toBeFalsy();
  expect(await page.evaluate(() => localStorage.getItem('refresh_token'))).toBeFalsy();

  await page.goto('/');
  await expect(page).toHaveURL(/\/login$/);
});

test('a hard reload keeps the session', async ({ page }) => {
  await loginAsDemo(page);
  await page.reload();

  await expect(page).toHaveURL(/:\d+\/$/);
  await expect(page.getByText('Demo User')).toBeVisible();
  await expect(page.getByText('Spicy Poblano Bolognese').first()).toBeVisible();
});

test('a refresh token the API rejects lands on login instead of stalling', async ({ page }) => {
  await loginAsDemo(page);
  await page.evaluate(() => {
    localStorage.setItem('token_valid_until', '0');
    localStorage.setItem('refresh_token', 'not-a-token');
  });

  await page.reload();
  await expect(page).toHaveURL(/\/login$/);
  await expect(page.getByRole('button', { name: 'Continue' })).toBeVisible();
});

test('an expired access token is refreshed on reload', async ({ page }) => {
  await loginAsDemo(page);
  await page.evaluate(() => localStorage.setItem('token_valid_until', '0'));

  const refreshed = page.waitForRequest((r) => r.url().endsWith('/token/refresh'));
  await page.reload();
  await refreshed;

  await expect(page.getByText('Demo User')).toBeVisible();
  await expect(page.getByText('Spicy Poblano Bolognese').first()).toBeVisible();
});
