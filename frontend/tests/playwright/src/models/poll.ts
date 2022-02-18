import { Page } from "@playwright/test";

export class PollPage {
  constructor(private page: Page) {}

  public async navigate() {
    await this.page.goto("/poll");
  }

  public async addRestaurant(name: string) {
    // Click text=Ein weiteres Restaurant hinzufügen
    await this.page.locator("text=Ein weiteres Restaurant hinzufügen").click();
    // Click [placeholder="Restaurant\ am\ Markt"]
    await this.page.locator('[placeholder="Restaurant\\ am\\ Markt"]').click();
    // Fill [placeholder="Restaurant\ am\ Markt"]
    await this.page
      .locator('[placeholder="Restaurant\\ am\\ Markt"]')
      .fill(name);

    // Click #modal button:has-text("Restaurant hinzufügen")
    await this.page
      .locator('#modal button:has-text("Restaurant hinzufügen")')
      .click();
  }

  public async cancelPoll() {
    await Promise.all([
      this.page.waitForNavigation(/*{ url: 'http://localhost:3000/app/confirm' }*/),
      this.page.locator("text=Umfrage beenden").click(),
    ]);
  }
}
