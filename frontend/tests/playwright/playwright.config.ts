import type { PlaywrightTestConfig } from "@playwright/test";
import { devices } from "@playwright/test";
import * as dotenv from "dotenv";
import * as path from "path";

const result = dotenv.config({ path: path.join(__dirname, "/.env") });

const config: PlaywrightTestConfig = {
  testDir: "./src",
  /* Maximum time one test can run for. */
  timeout: 30 * 1000,
  expect: {
    timeout: 5000,
  },

  forbidOnly: !!process.env.CI,

  retries: process.env.CI ? 2 : 0,

  workers: process.env.CI ? 1 : undefined,

  reporter: "html",

  webServer: {
    command: "cd ../../apps/ui && npm run start",
    port: 3000,
    timeout: 120 * 1000,
    reuseExistingServer: !process.env.CI,
  },

  use: {
    video: { mode: "on", size: { width: 1920, height: 1080 } },
    screenshot: "on",
    trace: "on",
    viewport: { width: 1920, height: 1080 },
  },

  /* Configure projects for major browsers */
  projects: [
    {
      name: "chromium",
      use: {
        ...devices["Desktop Chrome"],
      },
    },

    {
      name: "firefox",
      use: {
        ...devices["Desktop Firefox"],
      },
    },
  ],
};

export default config;
