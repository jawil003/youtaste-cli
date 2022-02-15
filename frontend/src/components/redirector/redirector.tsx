import React, { useContext } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { Routes } from "../../enums/routes.enum";
import { useUser } from "../../hooks/user.hook";
import { ProgressProvider } from "../progress-provider/progress-provider";

const mapStateToRoute: Record<string, string> = {
  ADMIN_NEW: Routes.ADMIN_NEW,
  CHOOSE_RESTAURANT: Routes.POLLS,
  CHOOSE_MEALS: Routes.ORDER_CONFIRM,
  DONE: Routes.ON_THE_WAY,
};

export interface Props {}

/**
 * An Redirector React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Redirector: React.FC<Props> = () => {
  const { data: user, isFetched, isError } = useUser();
  const location = useLocation();

  const provider = useContext(ProgressProvider);

  if (user && location.pathname === mapStateToRoute[provider ?? ""])
    return null;
  if (isFetched)
    return (
      <Navigate
        to={
          user && !isError && provider ? mapStateToRoute[provider] : Routes.NEW
        }
      />
    );
  return null;
};
