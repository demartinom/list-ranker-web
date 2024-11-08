import { useEffect, useState } from "react";

function WebSocketComponent() {
  const [socket, setSocket] = useState(null);
  const [, setMessages] = useState([]);
  const [input, setInput] = useState("test");

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      console.log("Connected to WebSocket");
    };

    ws.onmessage = (event) => {
      setMessages((prevMessages) => [...prevMessages, event.data]);
    };

    ws.onclose = () => {
      console.log("Disconnected from WebSocket");
    };

    ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    setSocket(ws);

    return () => ws.close();
  }, []);

  const sendMessage = () => {
    if (socket && input.trim()) {
      socket.send(input);
      setInput("");
    }
  };

  return <button onClick={sendMessage}>Click</button>;
}

export default WebSocketComponent;
