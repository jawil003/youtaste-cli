import { Page } from "@playwright/test";

export class OrderPage {
  constructor(private page: Page) {}

  public async navigate() {
    await this.page.goto("/order");
  }

  public async waitForPage() {
    await this.page.waitForFunction(
      () => {
        return window.location.pathname === "/app/confirm";
      },
      undefined,
      { timeout: 300000 }
    );
  }

  public async createOrder(value: { name: string; variants?: string[] }) {
    // Click text=Bestellung hinzufügen
    await this.page.locator("text=Bestellung hinzufügen").click();
    // Click [placeholder="Pizza\ Cipola"]
    await this.page.locator('[placeholder="Pizza\\ Cipola"]').click();
    // Fill [placeholder="Pizza\ Cipola"]
    await this.page.locator('[placeholder="Pizza\\ Cipola"]').fill(value.name);
    // Press Tab
    await this.page.locator('[placeholder="Pizza\\ Cipola"]').press("Tab");

    for (const variant of value?.variants ?? []) {
      await this.page.locator('[placeholder="Groß"]').fill(variant);
      // Click text=Mahlzeit*VarianteBestellung hinzufügen >> button >> nth=0
      await this.page
        .locator("text=Mahlzeit*VarianteBestellung hinzufügen >> button")
        .first()
        .click();
    }
    // Fill [placeholder="Groß"]

    // Click #modal button:has-text("Bestellung hinzufügen")
    await this.page
      .locator('#modal button:has-text("Bestellung hinzufügen")')
      .click();
  }

  public async cancelOrderTime() {
    await Promise.all([
      this.page.waitForNavigation(/*{ url: 'http://localhost:3000/app/on-the-way' }*/),
      this.page.locator("text=Bestellzeit beenden").click(),
    ]);
  }

  public async waitForSpinner() {
    await this.page.waitForSelector('*[data-qa="loading-spinner"]');
  }

  public async waitForPendingEnd() {
    await this.page.waitForFunction(
      () => {
        const res = document.querySelectorAll('*[data-qa="loading-spinner"]');

        return res.length === 0;
      },
      undefined,
      { timeout: 300000 }
    );
  }
}
