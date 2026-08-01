import { test, expect, loginAsDemo } from './fixtures.js';

async function swipe(page, distance) {
  await page.locator('.b-tabs.recipes').evaluate((el, dx) => {
    const box = el.getBoundingClientRect();
    const y = box.top + 30;
    const startX = box.left + box.width / 2;
    const at = (x) => new Touch({
      identifier: 1, target: el, clientX: x, clientY: y,
    });
    const fire = (type, x, moving) => el.dispatchEvent(new TouchEvent(type, {
      bubbles: true,
      cancelable: true,
      touches: moving ? [at(x)] : [],
      changedTouches: [at(x)],
    }));

    fire('touchstart', startX, true);
    fire('touchmove', startX + dx / 2, true);
    fire('touchmove', startX + dx, true);
    fire('touchend', startX + dx, false);
  }, distance);
}

const activeDay = (page) => page.locator('.b-tabs.recipes li.is-active');

test.beforeEach(async ({ page }) => {
  await loginAsDemo(page);
});

test('the slider replaces the grid on a phone', async ({ page }) => {
  await expect(page.locator('.b-tabs.recipes')).toBeVisible();
  await expect(page.locator('.week-grid')).toHaveCount(0);
});

test('the frozen clock opens the week on Wednesday', async ({ page }) => {
  await expect(activeDay(page)).toHaveText('We');
});

test('swiping left moves to the next day and toasts its name', async ({ page }) => {
  await swipe(page, -200);

  await expect(activeDay(page)).toHaveText('Th');
  await expect(page.locator('.toast')).toHaveText('Thursday');
});

test('swiping right moves to the previous day and toasts its name', async ({ page }) => {
  await swipe(page, 200);

  await expect(activeDay(page)).toHaveText('Tu');
  await expect(page.locator('.toast')).toHaveText('Tuesday');
});

test('a swipe shorter than the tolerance does nothing', async ({ page }) => {
  await swipe(page, -40);

  await expect(activeDay(page)).toHaveText('We');
  await expect(page.locator('.toast')).toHaveCount(0);
});

test('the week boundary stops the slider and suppresses the toast', async ({ page }) => {
  for (let i = 0; i < 3; i += 1) {
    await swipe(page, 200);
  }
  await expect(activeDay(page)).toHaveText('Mo');

  await swipe(page, 200);
  await expect(activeDay(page)).toHaveText('Mo');
  await expect(page.locator('.toast')).toHaveCount(0);
});

test('the visible day shows its own recipes', async ({ page }) => {
  await expect(page.locator('.tab-item:visible .card-header-title')).toHaveText([
    'Italian Noodle Soup',
    'Chicken Gyro Couscous Bowls',
  ]);

  await swipe(page, 200);
  await expect(page.locator('.tab-item:visible .card-header-title')).toHaveText([
    'Apple Pecan Breakfast Oatmeal Bake',
    'Moroccan-Style Chickpea & Tomato Stew',
    'Chicken Gyro Couscous Bowls',
  ]);
});
