import { useQuery } from "react-query";
import { Queries } from "../enums/queries.enum";
import AdminService from "../services/admin.service";

export const useIsAdmin = () =>
  useQuery(Queries.ISADMIN, async () => {
    const adminService = new AdminService();

    return await adminService.isAdmin();
  });
