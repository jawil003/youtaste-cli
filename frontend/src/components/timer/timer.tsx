import * as dayjs from "dayjs";
import duration from "dayjs/plugin/duration";
import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Routes } from "../../enums/routes.enum";
import { useTime } from "../../hooks/useTime.hook";
import { useTimer } from "../../hooks/useTimer.hook";

dayjs.extend(duration);

export interface Props {
  mode?: "POLL" | "ORDER";
}

/**
 * An Timer React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Timer: React.FC<Props> = ({ mode }) => {
  const navigate = useNavigate();
  const { resTime: initialTime, isFetched } = useTime("POLL");
  const { start, time } = useTimer(() => {
    if (mode === "POLL") {
      navigate(Routes.ORDER_CONFIRM);
    } else if (mode === "ORDER") {
    }
  });

  useEffect(() => {
    if (isFetched) {
      start(initialTime);
    }
  }, [isFetched, start, initialTime]);

  return (
    <div className="absolute top-0 left-0 w-full flex items-center justify-center">
      <div className="rounded-b-lg p-2 shadow-lg text-white font-semibold bg-red-500">
        {dayjs.duration(time, "milliseconds").format("HH:mm:ss")}
      </div>
    </div>
  );
};

Timer.defaultProps = { mode: "POLL" };
