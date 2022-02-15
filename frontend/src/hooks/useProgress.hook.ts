import { useCallback, useEffect, useState } from "react";
import ProgressService from "../services/progress.service";
import { useUser } from "./user.hook";

export const useProgress = () => {
  const [progress, setProgress] = useState<string | null>(null);
  const [shouldRefetch, setShouldRefetch] = useState(false);
  const { data: user } = useUser();

  const refetch = useCallback(() => {
    setShouldRefetch((prev) => !prev);
  }, []);

  useEffect(() => {
    if (!user) return;
    (async () => {
      const progressService = new ProgressService();
      const { progress: localProgress } = await progressService.getProgress();

      if (localProgress === progress) return;

      console.debug({ progress: localProgress }, "useProgress: getProgress");
      setProgress(localProgress);
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
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [shouldRefetch, user]);

  return { progress, isFetched: !!progress, refetch };
};
