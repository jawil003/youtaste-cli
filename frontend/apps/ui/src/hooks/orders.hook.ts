import { useQuery } from "react-query";
import { Queries } from "../enums/queries.enum";
import OrderService from "../services/order.service";

export const useOrders = () =>
  useQuery(Queries.ORDERS, async () => {
    const orderService = new OrderService();

    return await (
      await orderService.getAll()
    ).data;
  });
