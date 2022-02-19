import { useQuery } from "react-query";
import ScrapperService from "../services/scrapper.service";

export const useRestaurantUrl = () =>
  useQuery(["restaurantUrl"], async () => {
    const scrapperService = new ScrapperService();

    return await scrapperService.getRestaurantUrl();
  });
