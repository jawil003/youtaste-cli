import { useQuery } from "react-query";
import { Queries } from "../enums/queries.enum";
import OrderService from "../services/order.service";

export const useOrdersByUser = () =>
  useQuery(Queries.ORDERS_BY_USER, async () => {
    const orderService = new OrderService();

    return await (
      await orderService.getByUser()
    ).data;
  });
