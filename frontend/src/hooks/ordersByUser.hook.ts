import { useQuery } from "react-query";
import OrderService from "../services/order.service";

export const useOrdersByUser = () =>
  useQuery(["orders-by-user"], async () => {
    const orderService = new OrderService();

    return await (
      await orderService.getByUser()
    ).data;
  });
