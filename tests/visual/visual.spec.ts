import { test, expect } from '@playwright/test';

test('Homepage visual regression', async ({ page }) => {
  await page.goto('/');
  
  // Takes a screenshot and compares against snapshot in tests/snapshots/
  await expect(page).toHaveScreenshot('homepage.png', { maxDiffPixels: 100 });
});
