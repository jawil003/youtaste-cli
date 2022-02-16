import { useQuery } from "react-query";
import { Queries } from "../enums/queries.enum";
import UserService from "../services/user.service";

export const useUser = () =>
  useQuery(Queries.USER, async () => {
    const userService = new UserService();

    return (await userService.me())?.data;
  });
