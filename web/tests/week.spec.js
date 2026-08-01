import {
  test, expect, loginAsDemo, daySlot, recipeName, fillSlot, confirmDialog,
} from './fixtures.js';

test.beforeEach(async ({ page }) => {
  await loginAsDemo(page);
});

test('the seeded week renders in the grid', async ({ page }) => {
  await expect(page.locator('.week-grid > *')).toHaveCount(35);
  await expect(recipeName(daySlot(page, 0, 0))).toHaveText('Apple Pecan Breakfast Oatmeal Bake');
  await expect(recipeName(daySlot(page, 0, 1))).toHaveText('Spicy Poblano Bolognese');
  await expect(recipeName(daySlot(page, 2, 1))).toHaveText('Italian Noodle Soup');
  await expect(daySlot(page, 6, 0).getByRole('button')).toBeVisible();
});

test('an empty slot is filled through the autocomplete', async ({ page }) => {
  const slot = daySlot(page, 3, 1);
  await fillSlot(slot, 'Chicken Gyro Couscous Bowls');

  await expect(recipeName(slot)).toHaveText('Chicken Gyro Couscous Bowls');

  await page.reload();
  await expect(recipeName(daySlot(page, 3, 1))).toHaveText('Chicken Gyro Couscous Bowls');
});

test('the autocomplete only offers recipes for the course', async ({ page }) => {
  const breakfast = daySlot(page, 3, 0);
  await breakfast.getByRole('button').click();
  await expect(breakfast.locator('.dropdown-item')).toHaveText([
    'Lemon Ricotta Pancakes',
    'Apple Pecan Breakfast Oatmeal Bake',
  ]);

  const main = daySlot(page, 3, 1);
  await main.getByRole('button').click();
  await expect(main.locator('.dropdown-item')).toHaveCount(4);
  await expect(main.locator('.dropdown-item', { hasText: 'Lemon Ricotta Pancakes' })).toHaveCount(0);
});

test('a filled slot is emptied from the card footer', async ({ page }) => {
  const slot = daySlot(page, 0, 1);
  await expect(recipeName(slot)).toHaveText('Spicy Poblano Bolognese');

  await slot.locator('.card-footer-item').first().click();
  await expect(slot.getByRole('button')).toBeVisible();

  await page.reload();
  await expect(daySlot(page, 0, 1).getByRole('button')).toBeVisible();
});

test('fill gaps completes the week', async ({ page }) => {
  const fillGaps = page.getByRole('button', { name: 'Fill gaps' });
  await expect(fillGaps).toBeVisible();
  await fillGaps.click();

  await expect(page.locator('.week-grid .card-header-title')).toHaveCount(21);
  await expect(page.getByRole('button', { name: 'Clear week' })).toBeVisible();

  await page.reload();
  await expect(page.locator('.week-grid .card-header-title')).toHaveCount(21);
});

test('clear week empties every slot after confirmation', async ({ page }) => {
  await page.getByRole('button', { name: 'Fill gaps' }).click();
  await expect(page.locator('.week-grid .card-header-title')).toHaveCount(21);

  await page.getByRole('button', { name: 'Clear week' }).click();
  await expect(page.locator('.dialog')).toContainText('Remove all recipes chosen for the week?');
  await confirmDialog(page);

  await expect(page.locator('.week-grid .card-header-title')).toHaveCount(0);

  await page.reload();
  await expect(page.locator('.week-grid .card-header-title')).toHaveCount(0);
});

test('cancelling the clear-week dialog changes nothing', async ({ page }) => {
  await page.getByRole('button', { name: 'Fill gaps' }).click();
  await expect(page.locator('.week-grid .card-header-title')).toHaveCount(21);

  await page.getByRole('button', { name: 'Clear week' }).click();
  await page.locator('.dialog').getByRole('button', { name: 'Cancel' }).click();

  await expect(page.locator('.week-grid .card-header-title')).toHaveCount(21);
});
