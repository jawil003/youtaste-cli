import { backend } from "../config/backend";

export default class UserService {
  public async create(firstname: string, lastname: string) {
    return await backend.post("/api/user/create", { firstname, lastname });
  }

  public async remove() {
    return await backend.delete("/api/user");
  }

  public async me() {
    return await backend.get<{ firstname: string; lastname: string }>(
      "/api/user/me"
    );
  }
}
