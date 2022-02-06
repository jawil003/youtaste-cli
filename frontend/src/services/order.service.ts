import { backend } from "../config/backend";
import { Order } from "../types/order.type";

export default class OrderService {
  public async createOrUpdate(orders: Order[]) {
    return await backend.post<{ status: string }>("/api/orders", {
      orders,
    });
  }

  public async getByUser() {
    return await backend.get<{ orders: Order[] }>(`/api/orders/user`);
  }

  public async getByUserAndName(name: string) {
    return await backend.get<{ order: Order }>(`/api/orders/user/${name}`);
  }

  public async getAll() {
    return await backend.get<{ orders: Order[] }>("/api/orders/all");
  }

  public async deleteOrder(name: string) {
    return await backend.delete(`/api/orders/user/${name}`);
  }
}
