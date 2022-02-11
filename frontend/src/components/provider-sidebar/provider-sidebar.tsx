import React from "react";
import youtasteImg from "../../assets/youtaste-white-logo.png";
import { ReactComponent as LieferandoLogo } from "../../assets/lieferandoat-small.svg";
export interface Props {}

/**
 * An ProviderSidebar React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const ProviderSidebar: React.FC<Props> = () => {
  return (
    <div className="absolute top-0 right-0 h-full flex items-center justify-center flex-col gap-y-2 z-50">
      <a
        href="https://www.lieferando.de/"
        target="_blank"
        rel="noopener noreferrer"
        className="rounded-l-lg bg-orange-300 p-2 w-full flex items-center justify-center"
      >
        <LieferandoLogo width={40} />
      </a>
      <a
        href="https://youtaste.com/"
        target="_blank"
        rel="noopener noreferrer"
        className="rounded-l-lg bg-red-300 p-2 w-full aspect-square flex items-center justify-center"
      >
        <img src={youtasteImg} width={20} alt="Youtaste Logo" />
      </a>
    </div>
  );
};
