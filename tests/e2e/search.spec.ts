import { test, expect } from '@playwright/test';

test('Homepage load and search', async ({ page }) => {
  await page.goto('/');
  await expect(page).toHaveTitle(/Appenheimer/);

  const searchInput = page.locator('input[type="search"]');
  await searchInput.fill('figma');
  
  // Wait for results
  await expect(page.locator('.app-card')).toHaveCount(1);
});
