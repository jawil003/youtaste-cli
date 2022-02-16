import { useEffect, useState } from "react";
import OrderService from "../services/order.service";

export const useTime = (mode: "POLL" | "ORDER") => {
  const [resTime, setResTime] = useState<number | undefined>(undefined);

  useEffect(() => {
    (async () => {
      const orderService = new OrderService();

      setResTime((await orderService.getTime()).time);
    })();
  }, [mode]);

  return { resTime, isFetched: resTime !== undefined };
};
