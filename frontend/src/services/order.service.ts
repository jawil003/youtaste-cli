import axios from "axios";
import { Order } from "../types/order.type";

export default class OrderService {
  public async createOrUpdate(orders: Order[], user: string) {
    return await axios.post<{ status: string }>("/api/orders", {
      orders,
      user,
    });
  }

  public async getByUser(user: string) {
    return await axios.get<{ orders: Order[] }>(`/api/orders/${user}`);
  }

  public async getAll() {
    return await axios.get<{ orders: Order[] }>("/api/orders/all");
  }
}
