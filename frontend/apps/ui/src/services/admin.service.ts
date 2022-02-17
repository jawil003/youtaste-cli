import { backend } from "../config/backend";
import { BackendConfig } from "../types/config.type";

export default class AdminService {
  public async setConfig(value: BackendConfig) {
    return (await backend.post("/api/admin/set", value))?.data;
  }

  public async isAdmin() {
    return (
      (await backend.get<{ isAdmin: boolean }>("/api/admin/isAdmin"))?.data
        ?.isAdmin ?? false
    );
  }

  public async next() {
    return await backend.put("/api/admin/next");
  }
}
