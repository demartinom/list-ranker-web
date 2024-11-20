/* eslint-disable react/prop-types */

import { createContext, useEffect, useState } from "react";

export const WebSocketContext = createContext(null);

export default function WebSocketProvider({ children }) {
  const [socket, setSocket] = useState(null);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    setSocket(ws);
    socket.onopen = () => {
      console.log("Connected to WebSocket");
    };

    socket.onclose = () => {
      console.log("Disconnected from WebSocket");
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    return () => {
      ws.close();
    };
  }, []);

  return (
    <WebSocketContext.Provider value={socket}>
      {children}
    </WebSocketContext.Provider>
  );
}
