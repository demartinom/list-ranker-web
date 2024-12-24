import { WebSocketContext } from "./websocket";
import { useContext } from "react";

export default function Results() {
  const { endResult } = useContext(WebSocketContext);

  let ranking = endResult.map((item) => <li key={item}>{item}</li>);
  return endResult.length > 0 && <ol>{ranking}</ol>;
}
