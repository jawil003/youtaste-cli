import { useContext, useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { ProgressProvider } from "../components/progress-provider/progress-provider";
import { logger } from "../config/logger";
import { mapStateToRoute } from "../config/mapStateToConfig";
import { Routes } from "../enums/routes.enum";
import { useUser } from "./user.hook";

export const useRedirector = () => {
  const progress = useContext(ProgressProvider);
  const { data: user, isFetched } = useUser();
  const { pathname } = useLocation();

  const [route, setRoute] = useState<string | null>(null);

  useEffect(() => {
    if (isFetched) {
      if (user && progress) {
        const res = mapStateToRoute[progress];

        if (res === pathname) return;

        logger.debug({ progress: res }, "useRedirector: progress updated");
        setRoute(res);
      } else {
        if (Routes.NEW === pathname) return;
        setRoute(Routes.NEW);
        logger.debug(
          { progress: Routes.NEW },
          "useRedirector: progress updated"
        );
      }
    }
  }, [progress, isFetched, user, pathname]);

  return route;
};
