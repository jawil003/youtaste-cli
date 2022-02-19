import React from "react";
import { ReactComponent as LieferandoLogo } from "../../../assets/lieferandoat-small.svg";
import youtasteImg from "../../../assets/youtaste-white-logo.png";
import { Spinner } from "../../spinner/spinner";

export interface Props {
  url: string;
  pending?: boolean;
}

/**
 * An ProviderSidebarBadge React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const ProviderSidebarBadge: React.FC<Props> = ({ url, pending }) => {
  if (url.includes("lieferando")) {
    return (
      <a
        href={url}
        target="_blank"
        rel="noopener noreferrer"
        className="rounded-l-lg bg-orange-300 p-2 w-full flex items-center justify-center"
      >
        {pending && <Spinner />}
        <LieferandoLogo width={40} />
      </a>
    );
  } else {
    return (
      <a
        href={url}
        target="_blank"
        rel="noopener noreferrer"
        className="rounded-l-lg bg-red-300 p-2 w-full aspect-square flex items-center justify-center"
      >
        <img src={youtasteImg} width={20} alt="Youtaste Logo" />
      </a>
    );
  }
};

ProviderSidebarBadge.defaultProps = { pending: false };
