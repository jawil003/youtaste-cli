import { useContext, useEffect, useState } from "react";
import { ProgressProvider } from "../components/progress-provider/progress-provider";
import { mapStateToRoute } from "../config/mapStateToConfig";
import { Routes } from "../enums/routes.enum";
import { useUser } from "./user.hook";

export const useRedirector = () => {
  const progress = useContext(ProgressProvider);
  const { data: user, isFetched } = useUser();

  const [route, setRoute] = useState("");

  useEffect(() => {
    if (isFetched) {
      if (user && progress) setRoute(mapStateToRoute[progress]);
      else setRoute(Routes.NEW);
    }
  }, [progress, isFetched, user]);

  return route;
};
