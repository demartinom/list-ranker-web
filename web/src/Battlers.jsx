import { useContext } from "react";
import { WebSocketContext } from "./websocket";

export default function Battlers() {
  const { combatants } = useContext(WebSocketContext);

  return (
    <div>
      <h1>Battlers</h1>
    </div>
  );
}
