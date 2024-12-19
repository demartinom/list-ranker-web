/* eslint-disable react/prop-types */

import { createContext, useEffect, useState } from "react";

export const WebSocketContext = createContext(null);

export default function WebSocketProvider({ children }) {
  const [socket, setSocket] = useState(null);
  let [combatants, setCombatants] = useState([]);
  const [listOptions, setListOptions] = useState([]);
  const [endResult, setEndResult] = useState([]);
  
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      console.log("Connected to WebSocket");
    };

    ws.onclose = () => {
      console.log("Disconnected from WebSocket");
    };

    ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);

      switch (message.messageType) {
        case "List Options":
          setListOptions(message.options);
          break;
        case "Combatants":
          setCombatants(message.combatants);
          break;
        case "ResultsList":
          setEndResult(message.rankings)
          break
      }
    };

    setSocket(ws);

    return () => {
      ws.close();
    };
  }, []);

  return (
    <WebSocketContext.Provider value={{ socket, listOptions, combatants, endResult }}>
      {children}
    </WebSocketContext.Provider>
  );
}
