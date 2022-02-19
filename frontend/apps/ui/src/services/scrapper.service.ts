import { backend } from "../config/backend";

export default class ScrapperService {
  public async getRestaurantUrl() {
    return (
      await backend.get<{
        url: string;
        pending: boolean;
        provider: string;
      }>("/api/scrapper/url")
    )?.data;
  }
}
