import React from "react";
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

  return null;
};

Auth.defaultProps = { mode: "USER" };
