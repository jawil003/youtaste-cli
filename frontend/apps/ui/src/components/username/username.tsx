import React from "react";
import { LogoutIcon } from "@heroicons/react/outline";
import { useQueryClient } from "react-query";
import { Queries } from "../../enums/queries.enum";
import { Routes } from "../../enums/routes.enum";
import UserService from "../../services/user.service";
export interface Props {}

/**
 * An Username React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Username: React.FC<Props> = ({ children }) => {
  const queryClient = useQueryClient();

  return (
    <div className="flex gap-x-2">
      <button
        onClick={async () => {
          const userService = new UserService();
          await userService.remove();
          await queryClient.invalidateQueries(Queries.USER);
          window.location.pathname = "/app" + Routes.NEW;
        }}
      >
        <LogoutIcon className="text-gray-500" width={"1.25rem"} />
      </button>
      <span className="flex items-center text-gray-500 text-md">
        {children}
      </span>
    </div>
  );
};
