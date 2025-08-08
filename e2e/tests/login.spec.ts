import { expect } from "@playwright/test";
import { test } from "../fixtures/auth.fixture";

test("should login to existing profile", async ({ page, authUser }) => {
  const profileButton = page.locator("#my-mdn-plus-button");
  await expect(profileButton).toBeVisible();
  await profileButton.click();

  const loginHeading = page.locator(
    `//*[contains(@class, "submenu-item-heading") and text()='${authUser.login}']`
  );
  await expect(loginHeading, 'User login is visible when logged in').toBeVisible();
});
