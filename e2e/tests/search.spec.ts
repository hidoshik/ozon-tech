import { test, expect } from "@playwright/test";

test('should navigate to first article when pressing \'enter\'', async ({ page }) => {
    await page.goto('/en-US/');

    const searchInput = page.locator('#top-nav-search-input');
    await expect(searchInput).toBeEditable();
    await searchInput.fill('Web');

    const firstElement = page.locator('div.search-results div:first-child a');
    await expect(firstElement).toBeVisible();

    const articleName = await firstElement.locator('b').innerText();
    const articleTitle = articleName + ' | MDN';

    await Promise.all([
        page.waitForLoadState('load'),
        searchInput.press('Enter')
    ])

    await expect(page, 'Correct page title should display for the article').toHaveTitle(articleTitle);
})