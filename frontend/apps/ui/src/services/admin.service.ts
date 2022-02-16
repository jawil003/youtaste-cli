import { backend } from "../config/backend";

export default class AdminService {
  public async setTimeout(value: { pollMs: number; orderMs: number }) {
    return (
      await backend.post("/api/admin/timer", {
        pollTime: value.pollMs,
        orderTime: value.orderMs,
      })
    )?.data;
  }

  public async setLieferandoLogin(username: string, password: string) {
    return (
      await backend.post("/api/admin/lieferando", {
        username,
        password,
      })
    )?.data;
  }

  public async setYouTasteLogin(phone: string, password: string) {
    return (
      await backend.post("/api/admin/youtaste", {
        phone,
        password,
      })
    )?.data;
  }

  public async isAdmin() {
    return (
      (await backend.get<{ isAdmin: boolean }>("/api/admin/isAdmin"))?.data
        ?.isAdmin ?? false
    );
  }
}
