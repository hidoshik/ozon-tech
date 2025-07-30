import path from "path";
import fs from "fs";
import { Page } from "@playwright/test";

export async function takeAuthErrorScreenshot(page: Page, errorName: string) {
  const screenshotsDir = path.join(__dirname, "../screenshots");
  fs.mkdirSync(screenshotsDir, { recursive: true });

  const screenshotPath = path.join(screenshotsDir, `${errorName}.png`);

  await page.screenshot({
    path: screenshotPath,
    fullPage: true,
  });
}
