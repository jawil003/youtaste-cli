import React from "react";
import { ProviderSidebarBadge } from "./provider-sidebar-badge/provider-sidebar-badge";
export interface Props {}

/**
 * An ProviderSidebar React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const ProviderSidebar: React.FC<Props> = () => {
  return (
    <div className="absolute top-0 right-0 h-full flex items-center justify-center flex-col gap-y-2 z-50">
      <ProviderSidebarBadge url="https://www.lieferando.com/" />
      <ProviderSidebarBadge url="https://www.youtaste.com/" />
    </div>
  );
};
