import { useEffect, useState } from "react";
import PollService from "../services/poll.service";

export const useTime = (mode: "POLL" | "ORDER") => {
  const [resTime, setResTime] = useState<number | undefined>(undefined);

  useEffect(() => {
    (async () => {
      if (mode === "POLL") {
        const pollService = new PollService();

        setResTime((await pollService.getTime()).time);
      } else {
      }
    })();
  }, [mode]);

  return { resTime, isFetched: resTime !== undefined };
};
