import { test, expect, loginAsDemo } from './fixtures.js';

test('login page', async ({ page }) => {
  await page.goto('/login');
  await expect(page.getByRole('button', { name: 'Continue' })).toBeVisible();

  await expect(page).toHaveScreenshot('login.png', { fullPage: true });
});

test('current week', async ({ page }) => {
  await loginAsDemo(page);
  await expect(page.locator('.card-header-title:visible').first()).toBeVisible();

  await expect(page).toHaveScreenshot('week.png', { fullPage: true });
});

test('empty week', async ({ page }) => {
  await loginAsDemo(page);
  await expect(page.locator('.card-header-title:visible').first()).toBeVisible();

  await page.evaluate(async () => {
    const token = localStorage.getItem('token');
    const { API_ADDRESS } = window.config;
    for (let day = 0; day < 7; day += 1) {
      for (let slot = 0; slot < 3; slot += 1) {
        await fetch(`${API_ADDRESS}/week/day/${day}/slot/${slot}`, {
          method: 'DELETE',
          headers: { Authorization: `Bearer ${token}` },
        });
      }
    }
  });
  await page.reload();
  await expect(page.locator('.card-header-title')).toHaveCount(0);

  await expect(page).toHaveScreenshot('week-empty.png', { fullPage: true });
});

test('recipe list', async ({ page }) => {
  await loginAsDemo(page);
  await page.goto('/recipes');
  await expect(page.locator('tbody tr')).toHaveCount(7);

  await expect(page).toHaveScreenshot('recipes.png', { fullPage: true });
});

test('add recipe form', async ({ page }) => {
  await loginAsDemo(page);
  await page.goto('/recipes');
  await expect(page.locator('tbody tr')).toHaveCount(7);

  await page.locator('tbody tr', { hasText: 'Italian Noodle Soup' }).locator('button').first()
    .click();
  await expect(page.locator('.modal-card')).toContainText('Update recipe');

  await expect(page.locator('.modal-card')).toHaveScreenshot('recipe-form.png');
});
