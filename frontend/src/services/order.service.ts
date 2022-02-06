import { backend } from "../config/backend";
import { Order } from "../types/order.type";

export default class OrderService {
  public async createOrUpdate(orders: Order[]) {
    return await backend.post<{ status: string }>("/api/orders", {
      orders,
    });
  }

  public async getByUser(user: string) {
    return await backend.get<{ orders: Order[] }>(`/api/orders/${user}`);
  }

  public async getAll() {
    return await backend.get<{ orders: Order[] }>("/api/orders/all");
  }
}
