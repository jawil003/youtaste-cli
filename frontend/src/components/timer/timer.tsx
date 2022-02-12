import * as dayjs from "dayjs";
import duration from "dayjs/plugin/duration";
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Routes } from "../../enums/routes.enum";
import PollService from "../../services/poll.service";

dayjs.extend(duration);

export interface Props {}

/**
 * An Timer React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Timer: React.FC<Props> = () => {
  const [initialTimeUpdated, setInitialTimeUpdated] = useState(false);
  const [time, setTime] = React.useState(0);
  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      const pollService = new PollService();

      const initialtime = await pollService.getTime();
      setTime(initialtime?.time);
      setInitialTimeUpdated(true);
    })();
  }, []);

  useEffect(() => {
    let intervalId: number;

    const countDown = () => {
      if (time <= 0) {
        window.clearInterval(intervalId);
        navigate(Routes.NEW_ORDER);
        return;
      }

      setTime((time) => time - 1000);
      return () => window.clearInterval(intervalId);
    };
    intervalId = window.setInterval(countDown, 1000);
    return () => window.clearInterval(intervalId);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [initialTimeUpdated]);

  return (
    <div className="absolute top-0 left-0 w-full flex items-center justify-center">
      <div className="rounded-b-lg p-2 shadow-lg text-white font-semibold bg-red-500">
        {dayjs.duration(time).format("HH:mm:ss")}
      </div>
    </div>
  );
};
