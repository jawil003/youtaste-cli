import { Page } from "@playwright/test";

export class NewUserPage {
  constructor(private page: Page) {}

  public async navigate() {
    await this.page.goto("/app/new");
  }

  public async fillFirstname(firstname: string) {
    await this.page.locator('[placeholder="Maxine"]').click();

    // Fill input[name="username"]
    await this.page.locator('[placeholder="Maxine"]').fill(firstname);
  }

  public async fillLastname(lastname: string) {
    await this.page.locator('[placeholder="Musterfrau"]').fill("Will");
  }

  public async submit() {
    await Promise.all([
      this.page.waitForNavigation(/*{ url: 'http://localhost:3000/app/admin/new' }*/),
      this.page.locator('button:has-text("Benutzer/in anlegen")').click(),
    ]);
  }
}
