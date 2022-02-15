import React from "react";
import { Navigate } from "react-router-dom";
import { useIsAdmin } from "../../hooks/isAdmin.hook";
import { useUser } from "../../hooks/user.hook";
import { useRedirector } from "../../hooks/useRedirector";

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

  const path = useRedirector();

  if (isFetching || isFetchingAdmin) return null;

  switch (mode) {
    case "ADMIN": {
      if (isAdmin) {
        return <>{children}</>;
      }
      break;
    }
    case "USER": {
      if (user) {
        return <>{children}</>;
      }
      break;
    }
    default: {
      if (!user) {
        return <>{children}</>;
      }
      break;
    }
  }

  return <Navigate to={path} />;
};

Auth.defaultProps = { mode: "USER" };
