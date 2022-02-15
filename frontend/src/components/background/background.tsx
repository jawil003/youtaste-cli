import React from "react";
import { ReactComponent as BurgerLogo } from "../../assets/Hamburger-cuate.svg";
import { Logo } from "../../assets/logo/logo";
import { Routes } from "../../enums/routes.enum";
import { useUser } from "../../hooks/user.hook";
import { ActiveOnRoutes } from "../activeOnRoutes";
import { Username } from "../username/username";
export interface Props {}

/**
 * An Background React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Background: React.FC<Props> = ({ children }) => {
  const { data: user } = useUser();
  return (
    <div className="relative w-screen h-screen overflow-hidden">
      <div className="absolute top-0 left-0 w-full h-full">
        <div className="relative w-full flex justify-between p-5 z-30">
          <Logo className="" />
          {user && (
            <Username>
              {user?.firstname} {user?.lastname}
            </Username>
          )}
        </div>
        <ActiveOnRoutes
          routes={[
            Routes.ADMIN_NEW,
            Routes.NEW,
            Routes.ORDER_CONFIRM,
            Routes.POLLS,
          ]}
        >
          <BurgerLogo />
        </ActiveOnRoutes>
      </div>

      <div className="absolute top-0 left-0 w-full h-full z-10">{children}</div>
    </div>
  );
};
