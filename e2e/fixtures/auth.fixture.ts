import { test as base } from "@playwright/test";

export const test = base.extend<{
  authUser: { login: string; password: string };
}>({
  authUser: async ({ page }, use) => {
    const user = {
      login: process.env.TEST_USER!,
      password: process.env.TEST_PASSWORD!,
    };

    await page.goto("/en-US/");
    await page.click("a.login-link");

    await page.fill('[data-testid="input-field"]', user.login);
    await page.click('//button[text()="Sign up or sign in"]');

    await page.fill('[data-testid="input-field"]', user.password);
    await page.click('//button[text()="Sign in"]');

    await use(user);
  },
});
