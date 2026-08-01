import {
  test, expect, loginAsDemo, daySlot, recipeName, confirmDialog,
} from './fixtures.js';

const modal = (page) => page.locator('.modal-card');

function numberField(page, label) {
  return modal(page).locator('.field.has-numberinput', { hasText: label }).locator('input');
}

function row(page, name) {
  return page.locator('tbody tr', { hasText: name });
}

async function fillRecipeForm(page, values) {
  const form = modal(page);
  await form.locator('input[type=text]').fill(values.name);
  for (const course of values.courses) {
    await form.locator('label.checkbox', { hasText: course }).click();
  }
  await numberField(page, 'Calories').fill(String(values.calories));
  await numberField(page, 'Protein').fill(String(values.protein));
  await numberField(page, 'Fat').fill(String(values.fat));
  await numberField(page, 'Carbs').fill(String(values.carbs));
}

test.beforeEach(async ({ page }) => {
  await loginAsDemo(page);
  await page.getByRole('link', { name: 'Recipes' }).click();
});

test('the seeded recipes are listed', async ({ page }) => {
  await expect(page.locator('tbody tr')).toHaveCount(7);
  await expect(row(page, 'Spicy Poblano Bolognese')).toContainText('main');
  await expect(row(page, 'Lemon Ricotta Pancakes')).toContainText('breakfast');
  await expect(row(page, 'Lemon Ricotta Pancakes')).toContainText('pudding');
});

test('a recipe is added through the navbar modal', async ({ page }) => {
  await page.getByRole('button', { name: 'Add Recipe' }).click();
  await expect(modal(page)).toContainText('Add new');

  await fillRecipeForm(page, {
    name: 'Test Pancakes',
    courses: ['breakfast'],
    calories: 400,
    protein: 12,
    fat: 8,
    carbs: 60,
  });
  await page.getByRole('button', { name: 'Create' }).click();

  await expect(modal(page)).toHaveCount(0);
  await expect(page.locator('tbody tr')).toHaveCount(8);
  await expect(row(page, 'Test Pancakes')).toContainText('breakfast');

  await page.reload();
  await expect(row(page, 'Test Pancakes')).toContainText('400');
});

test('a recipe with no nutrition never reaches the API', async ({ page }) => {
  let posted = false;
  page.on('request', (r) => {
    if (r.method() === 'POST' && r.url().endsWith('/recipe')) posted = true;
  });

  await page.getByRole('button', { name: 'Add Recipe' }).click();
  await modal(page).locator('input[type=text]').fill('Missing Nutrition');
  await page.getByRole('button', { name: 'Create' }).click();

  await expect(modal(page)).toBeVisible();
  await expect(page.locator('tbody tr')).toHaveCount(7);
  expect(posted).toBe(false);
});

test('an existing recipe is edited', async ({ page }) => {
  await row(page, 'Italian Noodle Soup').locator('button').first().click();
  await expect(modal(page)).toContainText('Update recipe');
  await expect(modal(page).locator('input[type=text]')).toHaveValue('Italian Noodle Soup');

  await modal(page).locator('input[type=text]').fill('Italian Noodle Soup v2');
  await numberField(page, 'Calories').fill('321');
  await page.getByRole('button', { name: 'Update' }).click();

  await expect(modal(page)).toHaveCount(0);
  await expect(page.locator('tbody tr')).toHaveCount(7);
  await expect(row(page, 'Italian Noodle Soup v2')).toContainText('321');

  await page.reload();
  await expect(row(page, 'Italian Noodle Soup v2')).toContainText('321');
});

test('editing a second recipe does not show the first one', async ({ page }) => {
  await row(page, 'Italian Noodle Soup').locator('button').first().click();
  await expect(modal(page).locator('input[type=text]')).toHaveValue('Italian Noodle Soup');
  await page.getByRole('button', { name: 'Close' }).click();
  await expect(modal(page)).toHaveCount(0);

  await row(page, 'Chocolate Fudgy Brownies').locator('button').first().click();
  await expect(modal(page).locator('input[type=text]')).toHaveValue('Chocolate Fudgy Brownies');
});

test('an unused recipe is deleted after confirmation', async ({ page }) => {
  await row(page, 'Lemon Ricotta Pancakes').locator('button').last().click();
  await expect(page.locator('.dialog')).toContainText('Remove Lemon Ricotta Pancakes?');
  await confirmDialog(page);

  await expect(page.locator('tbody tr')).toHaveCount(6);
  await expect(row(page, 'Lemon Ricotta Pancakes')).toHaveCount(0);

  await page.reload();
  await expect(page.locator('tbody tr')).toHaveCount(6);
});

test('cancelling the delete dialog keeps the recipe', async ({ page }) => {
  await row(page, 'Lemon Ricotta Pancakes').locator('button').last().click();
  await page.locator('.dialog').getByRole('button', { name: 'Cancel' }).click();

  await expect(page.locator('tbody tr')).toHaveCount(7);
});

test('deleting a recipe used this week warns and empties its slots', async ({ page }) => {
  await row(page, 'Spicy Poblano Bolognese').locator('button').last().click();
  await expect(page.locator('.dialog')).toContainText('is used in the current week');
  await confirmDialog(page);

  await expect(page.locator('tbody tr')).toHaveCount(6);

  await page.getByRole('link', { name: 'Week' }).click();
  await expect(page.getByText('Spicy Poblano Bolognese')).toHaveCount(0);
  await expect(daySlot(page, 0, 1).getByRole('button')).toBeVisible();

  await page.reload();
  await expect(recipeName(daySlot(page, 0, 2))).toHaveText('Moroccan-Style Chickpea & Tomato Stew');
  await expect(daySlot(page, 0, 1).getByRole('button')).toBeVisible();
});
