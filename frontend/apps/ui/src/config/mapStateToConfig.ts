import { Routes } from "../enums/routes.enum";

export const mapStateToRoute: Record<string, string> = {
  ADMIN_NEW: Routes.ADMIN_NEW,
  CHOOSE_RESTAURANT: Routes.POLLS,
  GET_URL_AND_OPENING_TIMES: Routes.WAIT_FOR_SCRAPPING_URL_AND_OPENING_TIMES,
  CHOOSE_MEALS: Routes.ORDER_CONFIRM,
  ORDER: Routes.ORDER_IN_PROGRESS,
  DONE: Routes.ON_THE_WAY,
};
