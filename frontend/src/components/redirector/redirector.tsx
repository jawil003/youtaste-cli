import React from "react";
import { Navigate, useLocation } from "react-router-dom";

import { useRedirector } from "../../hooks/useRedirector";
export interface Props {}

/**
 * An Redirector React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Redirector: React.FC<Props> = () => {
  const { pathname: currentPathname } = useLocation();

  const pathname = useRedirector();

  if (pathname && pathname !== currentPathname)
    return <Navigate to={pathname} />;

  return null;
};
