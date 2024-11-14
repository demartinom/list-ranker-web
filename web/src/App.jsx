import { useEffect, useState } from "react";
import ListSelection from "./ListSelection";
import initWebSocket from "./websocket";

function App() {
  const [socket, setSocket] = useState(null);
  useEffect(() => {
    setSocket(initWebSocket());

    return () => {
      if (socket) {
        socket.close();
      }
    };
  }, []);
  return (
    <>
      <ListSelection socket={socket} />
    </>
  );
}

export default App;
