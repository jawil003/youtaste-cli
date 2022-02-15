import { backend } from "../config/backend";

export default class ProgressService {
  public async getProgress() {
    return await (
      await backend.get<{ progress: string }>("/api/progress")
    ).data;
  }

  public async goNext() {
    return await (
      await backend.put("/api/progress")
    ).data;
  }
}
