import { test } from "@playwright/test";
import AdminNewPage from "./models/admin-new";
import { NewUserPage } from "./models/new-user";
import { OrderPage } from "./models/order";
import { PollPage } from "./models/poll";

test("Running trough Process as Admin", async ({ page }) => {
  const userPage = new NewUserPage(page);

  await userPage.navigate();

  await userPage.fillFirstname("Jannik");
  await userPage.fillLastname("Will");
  await userPage.submit();

  const adminNewPage = new AdminNewPage(page);

  await adminNewPage.fillYoutastePhone(process.env.YOUTASTE_PHONE);
  await adminNewPage.fillYoutastePassword(process.env.YOUTASTE_PASSWORD);
  await adminNewPage.fillLieferandoUsername(process.env.LIEFERANDO_USERNAME);
  await adminNewPage.fillLieferandoPassword(process.env.LIEFERANDO_PASSWORD);
  await adminNewPage.toggleOpeningTimes();
  await adminNewPage.submit();

  const pollPage = new PollPage(page);

  await pollPage.addRestaurant("American Food Factory", true);
  await pollPage.cancelPoll();

  const orderPage = new OrderPage(page);

  await orderPage.createOrder({ name: "Maxi Burger" });
  await orderPage.cancelOrderTime();
});
