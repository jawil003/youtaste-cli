import { useQuery } from "react-query";
import OrderService from "../services/order.service";

export const useOrders = () =>
  useQuery(["orders"], async () => {
    const orderService = new OrderService();

    return await (
      await orderService.getAll()
    ).data;
  });
