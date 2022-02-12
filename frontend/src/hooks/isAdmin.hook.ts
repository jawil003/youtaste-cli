import { useQuery } from "react-query";
import AdminService from "../services/admin.service";

export const useIsAdmin = () =>
  useQuery(["is-admin"], async () => {
    const adminService = new AdminService();

    return await adminService.isAdmin();
  });
