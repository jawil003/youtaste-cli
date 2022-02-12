import React from "react";
import { Navigate } from "react-router-dom";
import { Routes } from "../../enums/routes.enum";
import { useIsAdmin } from "../../hooks/isAdmin.hook";
import { useUser } from "../../hooks/user.hook";

export interface Props {
  mode?: "USER" | "NO_USER" | "ADMIN";
}

/**
 * An Auth React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Auth: React.FC<Props> = ({ mode, children }) => {
  const { data: user, isFetching } = useUser();
  const { data: isAdmin, isFetching: isFetchingAdmin } = useIsAdmin();

  if (isFetching || isFetchingAdmin) return null;

  switch (mode) {
    case "ADMIN": {
      if (isAdmin) {
        return <>{children}</>;
      } else {
        if (user) {
          return <Navigate to={Routes.ORDER_CONFIRM} />;
        } else {
          return <Navigate to={Routes.NEW} />;
        }
      }
    }
    case "USER": {
      if (user) {
        return <>{children}</>;
      } else {
        return <Navigate to={Routes.NEW} />;
      }
    }
    default: {
      if (!user) {
        return <>{children}</>;
      } else {
        return <Navigate to={Routes.ORDER_CONFIRM} />;
      }
    }
  }
};

Auth.defaultProps = { mode: "USER" };
