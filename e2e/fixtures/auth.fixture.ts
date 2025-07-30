import { test as base, expect } from "@playwright/test";
import { takeAuthErrorScreenshot } from "../utils/screenshot";

export const test = base.extend<{
  authUser: { login: string; password: string };
}>({
  authUser: async ({ page }, use) => {
    const user = {
      login: process.env.TEST_USER!,
      password: process.env.TEST_PASSWORD!,
    };

    try {
        await page.goto("/en-US/", { waitUntil: "networkidle" });

        const loginButton = page.locator('a.login-link');
        await expect(loginButton).toBeVisible();
        await loginButton.click();

        const loginIngut = page.locator('[data-testid="input-field"]');
        await expect(loginIngut).toBeVisible();
        await loginIngut.fill(user.login);
        
        const continueButton = page.locator('//button[text()="Sign up or sign in"]');
        await expect(continueButton).toBeVisible();
        await continueButton.click();
        
        const passwordIngut = page.locator('[data-testid="input-field"]');
        await expect(passwordIngut).toBeVisible();
        await passwordIngut.fill(user.password);

        const signInButton = page.locator('//button[text()="Sign in"]');
        await expect(signInButton).toBeVisible();
        await signInButton.click();

        await use(user);
    } catch (error) {
        await takeAuthErrorScreenshot(page, "auth-error.png");
        throw new Error(`Auth fixture failed: ${error}`);
    }
  },
});
