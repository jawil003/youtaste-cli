import { useCallback, useEffect, useState } from "react";
import ProgressService from "../services/progress.service";

export const useProgress = () => {
  const [progress, setProgress] = useState<string | null>(null);
  const [shouldRefetch, setShouldRefetch] = useState(false);

  const refetch = useCallback(() => {
    setShouldRefetch((prev) => !prev);
  }, []);

  useEffect(() => {
    (async () => {
      const progressService = new ProgressService();
      const { progress } = await progressService.getProgress();
      console.debug({ progress }, "useProgress: getProgress");
      setProgress(progress);
    })();

    const handleMessage = (event: MessageEvent<string>) => {
      console.debug({ progress: event.data }, "useProgress: updated via ws");
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
  }, [shouldRefetch]);

  return { progress, isFetched: !!progress, refetch };
};
