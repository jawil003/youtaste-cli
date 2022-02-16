import { Routes } from "../enums/routes.enum";

export const mapStateToRoute: Record<string, string> = {
  ADMIN_NEW: Routes.ADMIN_NEW,
  CHOOSE_RESTAURANT: Routes.POLLS,
  CHOOSE_MEALS: Routes.ORDER_CONFIRM,
  DONE: Routes.ON_THE_WAY,
};
