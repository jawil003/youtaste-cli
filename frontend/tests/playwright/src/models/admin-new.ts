import { Page } from "@playwright/test";

export default class AdminNewPage {
  constructor(private page: Page) {}

  public async navigate() {
    await this.page.goto("/admin/new");
  }

  public async fillYoutastePhone(phone: string) {
    await this.page.locator('input[name="youtastePhone"]').click();

    // Fill input[name="youtastePhone"]
    await this.page.locator('input[name="youtastePhone"]').fill(phone);
  }

  public async fillYoutastePassword(password: string) {
    // Click input[name="youtastePassword"]
    await this.page.locator('input[name="youtastePassword"]').click();

    // Fill input[name="youtastePassword"]
    await this.page.locator('input[name="youtastePassword"]').fill(password);
  }

  public async fillLieferandoUsername(username: string) {
    // Click input[name="lieferandoUsername"]
    await this.page.locator('input[name="lieferandoUsername"]').click();

    // Fill input[name="lieferandoUsername"]
    await this.page.locator('input[name="lieferandoUsername"]').fill(username);
  }

  public async fillLieferandoPassword(password: string) {
    // Click input[name="lieferandoPassword"]
    await this.page.locator('input[name="lieferandoPassword"]').click();

    // Fill input[name="lieferandoPassword"]
    await this.page.locator('input[name="lieferandoPassword"]').fill(password);
  }

  public async toggleOpeningTimes() {
    await this.page
      .locator(
        "text=Andere EinstellungenBestelldatum*Öffnungszeiten des Restaurants bevorzugen (fall >> div"
      )
      .nth(3)
      .click();
  }

  public async submit() {
    await Promise.all([
      this.page.waitForNavigation(/*{ url: 'http://localhost:3000/app/poll' }*/),
      this.page.locator('button:has-text("Submit")').click(),
    ]);
  }
}
