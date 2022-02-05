import { useQuery } from "react-query";
import OrderService from "../services/order.service";

export const useOrdersByUser = (user: string) =>
  useQuery(["orders", user], async () => {
    const orderService = new OrderService();

    return await (
      await orderService.getByUser(user)
    ).data;
  });
