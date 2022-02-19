import React, { useEffect } from "react";
import { logger } from "../../config/logger";
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

  useEffect(() => {
    if (restaurant?.pending) {
      window.setTimeout(async () => {
        const res = await refetch();
        logger.debug(
          { old: restaurant, new: res },
          "ProviderSidebarOrders: Refetched url"
        );
      }, 1000);
    }
  }, [restaurant?.pending, refetch, restaurant]);

  return (
    <div className="absolute top-0 right-0 h-full flex items-center justify-center flex-col gap-y-2 z-50">
      <ProviderSidebarBadge url={restaurant?.url ?? ""} />
    </div>
  );
};
