import { createRoot } from "react-dom/client";
import App from "./App.jsx";
import WebSocketProvider from "./websocket.jsx";

createRoot(document.getElementById("root")).render(
  <WebSocketProvider>
    <App />
  </WebSocketProvider>
);
