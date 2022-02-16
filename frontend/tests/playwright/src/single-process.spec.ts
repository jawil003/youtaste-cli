import { test } from "@playwright/test";

test("Running trough Process as Admin", async ({ page }) => {
  // Go to http://localhost:3000/app
  await page.goto("http://localhost:3000/app");

  // Go to http://localhost:3000/app/new
  await page.goto("http://localhost:3000/app/new");

  // Click [placeholder="Maxine"]
  await page.locator('[placeholder="Maxine"]').click();

  // Fill [placeholder="Maxine"]
  await page.locator('[placeholder="Maxine"]').fill("Jannik");

  // Press Tab
  await page.locator('[placeholder="Maxine"]').press("Tab");

  // Fill [placeholder="Musterfrau"]
  await page.locator('[placeholder="Musterfrau"]').fill("Will");

  // Click button:has-text("Benutzer/in anlegen")
  await Promise.all([
    page.waitForNavigation(/*{ url: 'http://localhost:3000/app/admin/new' }*/),
    page.locator('button:has-text("Benutzer/in anlegen")').click(),
  ]);

  // Click input[name="youtastePhone"]
  await page.locator('input[name="youtastePhone"]').click();

  // Fill input[name="youtastePhone"]
  await page.locator('input[name="youtastePhone"]').fill("+4917624615787");

  // Press Tab
  await page.locator('input[name="youtastePhone"]').press("Tab");

  // Click input[name="youtastePassword"]
  await page.locator('input[name="youtastePassword"]').click();

  // Fill input[name="youtastePassword"]
  await page.locator('input[name="youtastePassword"]').fill("HZWUKUGP42C9xG");

  // Click input[name="lieferandoUsername"]
  await page.locator('input[name="lieferandoUsername"]').click();

  // Fill input[name="lieferandoUsername"]
  await page.locator('input[name="lieferandoUsername"]').fill("j.w98@gmx.de");

  // Click input[name="lieferandoPassword"]
  await page.locator('input[name="lieferandoPassword"]').click();

  // Fill input[name="lieferandoPassword"]
  await page.locator('input[name="lieferandoPassword"]').fill("9kg1a739");

  // Click text=Andere EinstellungenBestelldatum*Öffnungszeiten des Restaurants bevorzugen (fall >> div >> nth=3
  await page
    .locator(
      "text=Andere EinstellungenBestelldatum*Öffnungszeiten des Restaurants bevorzugen (fall >> div"
    )
    .nth(3)
    .click();

  // Click button:has-text("Submit")
  await page.locator('button:has-text("Submit")').click();
});
