import { useContext } from "react";
import { WebSocketContext } from "./websocket";

export default function Battlers() {
  const { socket, combatants } = useContext(WebSocketContext);

  return (
    <div>
      <h1>Battlers</h1>
      {combatants.length > 0 ? (
        <ul>
          {combatants.map((item, index) => (
            <button key={index} onClick={()=>sendResult(socket, item)}>
              {item.Name}
            </button>
          ))}
        </ul>
      ) : (
        <p>No combatants available yet.</p>
      )}
    </div>
  );
}

function sendResult(socket, choice) {
  if (socket) {
    let result = JSON.stringify({ messageType: "Result", winner: choice });
    console.log(result)
    socket.send(result);
  }
}
