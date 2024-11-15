let socket;

function initWebSocket() {
  socket = new WebSocket("ws://localhost:8080/ws");

  socket.onopen = () => {
    console.log("Connected to WebSocket");
  };

 

  socket.onclose = () => {
    console.log("Disconnected from WebSocket");
  };

  socket.onerror = (error) => {
    console.error("WebSocket error:", error);
  };

  // setSocket(socket);

  return socket;
}

export default initWebSocket;
