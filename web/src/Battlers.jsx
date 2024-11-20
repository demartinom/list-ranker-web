/* eslint-disable react/prop-types */
import { useEffect, useState } from "react";
import { WebSocketContext } from "./websocket";

export default function Battlers() {
  const socket = WebSocketContext;
  let [combatants, setCombatants] = useState([]);

  useEffect(() => {
    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      console.log(message);
      if (message.messageType == "Combatants") {
        setCombatants(message.options);
      }
    };
  }, [socket]);

  console.log(combatants);

  return <div>Battlers</div>;
}
