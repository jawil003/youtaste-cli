import React from "react";
import { ReactComponent as BurgerLogo } from "../../assets/Hamburger-cuate.svg";
import { Logo } from "../../assets/logo/logo";
export interface Props {}

/**
 * An Background React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Background: React.FC<Props> = ({ children }) => {
  return (
    <div className="relative w-screen h-screen overflow-hidden">
      <div className="absolute top-0 left-0 w-full h-full z-0">
        <Logo className="absolute top-4 left-6" />
        <BurgerLogo />
      </div>

      <div className="absolute top-0 left-0 w-full h-full z-10">{children}</div>
    </div>
  );
};
