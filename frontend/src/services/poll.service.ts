import { backend } from "../config/backend";
import { Poll } from "../types/poll.type";

export default class PollService {
  public async getAll() {
    return await (
      await backend.get<Poll[]>("/api/polls")
    ).data;
  }
}
