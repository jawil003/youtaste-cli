import { useEffect, useState } from "react";
import { Poll } from "../types/poll.type";

export const usePolls = () => {
  const [polls, setPolls] = useState<Poll[]>([]);

  useEffect(() => {}, []);

  const websocket = new WebSocket("ws://localhost:80/api/polls");

  websocket.onopen = () => {};

  websocket.onmessage = (event) => {
    setPolls(JSON.parse(event.data));
  };

  window.addEventListener("beforeunload", () => {
    websocket.close();
  });

  return polls;
};
