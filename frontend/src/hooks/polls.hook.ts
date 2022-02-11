import { useEffect, useState } from "react";
import PollService from "../services/poll.service";
import { Poll } from "../types/poll.type";

export const usePolls = () => {
  const [polls, setPolls] = useState<{ [x: string]: number }>({});

  useEffect(() => {
    (async () => {
      const pollService = new PollService();

      const polls = await pollService.getAll();

      setPolls((old) => {
        for (const poll of polls) old[poll.restaurantName] = poll.count ?? 1;
        console.log(old);
        return { ...old };
      });
    })();
    const websocket = new WebSocket("ws://localhost:80/api/polls/ws");

    websocket.onopen = () => {};

    websocket.onmessage = (event) => {
      setPolls(JSON.parse(event.data));
    };

    window.addEventListener("beforeunload", () => {
      websocket.close();
    });
    return () => {
      websocket.close();
    };
  }, []);

  return polls;
};
