import React from "react";
import { useCookies } from "react-cookie";
import { Navigate } from "react-router-dom";
import { Routes } from "../../enums/routes.enum";
import { useUser } from "../../hooks/user.hook";

export interface Props {
  mode?: "USER" | "NO_USER";
}

/**
 * An Auth React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Auth: React.FC<Props> = ({ mode, children }) => {
  const { data: user } = useUser();
  switch (mode) {
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
