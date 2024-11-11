import { useEffect, useState } from "react";

function WebSocketComponent() {
  const [socket, setSocket] = useState(null);
  const [listOptions, setListOptions] = useState([]);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      console.log("Connected to WebSocket");
    };

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      if (message.messageType == "List Options") {
        setListOptions(message.options);
      }
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

  // const sendMessage = () => {
  //   if (socket && input.trim()) {
  //     socket.send(input);
  //     setInput("");
  //   }
  // };
  const premadeOptions = listOptions.map((item, index) => (
    <button key={index}>{item}</button>
  ));

  return (
    <div>
      <h1>Choose a template list</h1>
      <ul>{premadeOptions}</ul>
    </div>
  );
}

export default WebSocketComponent;
