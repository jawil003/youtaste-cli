import { useEffect, useState } from "react";
import PollService from "../services/poll.service";
import { PollWithoutCount } from "../types/poll-without-count.type";

export const usePolls = () => {
  const [polls, setPolls] = useState<{ [x: string]: number }>({});

  useEffect(() => {
    (async () => {
      const pollService = new PollService();

      const polls = await pollService.getAll();

      setPolls((old) => {
        for (const poll of polls) old[poll.restaurantName] = poll.count ?? 1;
        return { ...old };
      });
    })();

    const handleMessage = (event: MessageEvent<any>) => {
      setPolls((old) => {
        const res = JSON.parse(event.data) as PollWithoutCount;

        const oldCount = old[res.restaurantName] ?? 0;

        old[res.restaurantName] = oldCount + 1;

        return { ...old };
      });
    };
    const websocket = new WebSocket(
      `ws://${process.env.REACT_APP_BASE_URL?.replace(
        "http://",
        ""
      )}/api/polls/ws`
    );

    websocket.onmessage = handleMessage;

    window.addEventListener("beforeunload", () => {
      websocket.close();
    });
    return () => {
      setPolls({});
      websocket.onclose = null;
      websocket.onmessage = null;
      websocket.onopen = null;
      websocket.close();
    };
  }, []);

  return polls;
};
