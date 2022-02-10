import React, { useEffect } from "react";

export interface Props {}

/**
 * An Poll React Component.
 * @author Jannik Will
 * @version 0.1
 */
export const Poll: React.FC<Props> = () => {
  useEffect(() => {
    const websocket = new WebSocket("ws://localhost:80/api/polls");

    websocket.onopen = () => {
      console.log("connected");
    };

    websocket.onmessage = (event) => {
      console.log(event.data);
    };

    window.addEventListener("beforeunload", () => {
      websocket.close();
    });
  });

  return <div></div>;
};
