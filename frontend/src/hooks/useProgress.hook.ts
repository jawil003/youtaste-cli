import { useEffect, useState } from "react";
import ProgressService from "../services/progress.service";

export const useProgress = () => {
  const [progress, setProgress] = useState<string | null>(null);

  useEffect(() => {
    (async () => {
      const progressService = new ProgressService();

      setProgress((await progressService.getProgress())?.progress);
    })();

    const handleMessage = (event: MessageEvent<string>) => {
      setProgress(event.data);
    };
    const websocket = new WebSocket(
      `ws://${process.env.REACT_APP_BASE_URL?.replace(
        "http://",
        ""
      )}/api/progress/ws`
    );

    websocket.onmessage = handleMessage;

    window.addEventListener("beforeunload", () => {
      websocket.close();
    });
    return () => {
      setProgress(null);
      websocket.onclose = null;
      websocket.onmessage = null;
      websocket.onopen = null;
      websocket.close();
    };
  }, []);

  return { progress, isFetched: !!progress };
};
