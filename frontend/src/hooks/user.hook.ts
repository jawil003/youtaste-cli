import { useQuery } from "react-query";
import UserService from "../services/user.service";

export const useUser = () =>
  useQuery("user", async () => {
    const userService = new UserService();

    return (await userService.me())?.data;
  });
