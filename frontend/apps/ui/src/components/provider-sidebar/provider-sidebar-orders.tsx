import React from "react";
import { logger } from "../../config/logger";
import useInterval from "../../hooks/interval.hook";
import { useRestaurantUrl } from "../../hooks/restaurantUrl.hook";
import { ProviderSidebarBadge } from "./provider-sidebar-badge/provider-sidebar-badge";

export interface Props {}

/**
 * An ProviderSidebarOrders React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const ProviderSidebarOrders: React.FC<Props> = () => {
  const { data: restaurant, refetch } = useRestaurantUrl();

  useInterval(
    async () => {
      const res = await refetch();
      logger.debug(
        { old: restaurant, new: res },
        "ProviderSidebarOrders: Refetched url"
      );
    },
    restaurant?.pending ? 1000 : null
  );

  return (
    <div className="absolute top-0 right-0 h-full flex items-center justify-center flex-col gap-y-2 z-50">
      <ProviderSidebarBadge
        url={restaurant?.url ?? ""}
        pending={restaurant?.pending}
      />
    </div>
  );
};
