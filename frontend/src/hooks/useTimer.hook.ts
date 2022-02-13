import { useCallback, useEffect, useState } from "react";

let intervalId: number | undefined;

export const useTimer = (onStopCallback?: () => void) => {
  const [time, setTime] = useState(0);

  useEffect(() => {
    if (time < 1000 && intervalId) {
      window.clearInterval(intervalId);
      intervalId = undefined;
      onStopCallback?.();
    }
  }, [time, onStopCallback]);

  const startTimer = useCallback((ms?: number) => {
    if (ms) {
      setTime(ms);
    }

    intervalId = window.setInterval(() => {
      setTime((time) => {
        if (time > 0) {
          return time - 1000;
        } else {
          return time;
        }
      });
    }, 1000);
  }, []);

  const stopTimer = useCallback(() => {
    if (intervalId) {
      window.clearInterval(intervalId);
    }
  }, []);

  return { start: startTimer, isStarted: !!intervalId, stop: stopTimer, time };
};
