import React from "react";
import { useLocation } from "react-router-dom";

export interface Props {
  routes: string[];
}

/**
 * An ActiveOnRoutes React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const ActiveOnRoutes: React.FC<Props> = ({ children, routes }) => {
  const { pathname } = useLocation();

  const active = routes.find((route) => pathname.startsWith(route));

  if (active) return <>{children}</>;
  return null;
};
