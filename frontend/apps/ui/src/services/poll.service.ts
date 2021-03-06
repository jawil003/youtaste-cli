import { backend } from "../config/backend";
import { Poll } from "../types/poll.type";

export default class PollService {
  public async getAll() {
    return await (
      await backend.get<Poll[]>("/api/polls")
    ).data;
  }

  public async create(poll: Poll) {
    await backend.post("/api/polls/new", poll);
  }

  public async getTime() {
    return (await backend.get<{ time: number }>("/api/polls/timer"))?.data;
  }
}
