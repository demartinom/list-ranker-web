import { useContext } from "react";
import { WebSocketContext } from "./websocket";

export default function Battlers() {
  const { combatants } = useContext(WebSocketContext);

  return (
    <div>
      <h1>Battlers</h1>
      {combatants.length > 0 ? (
        <ul>
          {combatants.map((item, index) => (
            <li key={index}>{item.Name}</li>
          ))}
        </ul>
      ) : (
        <p>No combatants available yet.</p>
      )}
    </div>
  );
}
