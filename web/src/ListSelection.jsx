import { useContext, useState } from "react";
import { WebSocketContext } from "./websocket";

export default function ListSelection() {
  const { socket, listOptions } = useContext(WebSocketContext);

  const [userInput, setUserInput] = useState("");

  // Create array with items in user created list
  const createList = () => {
    const formattedList = userInput.split(", ");
    sendChoice("Custom List", formattedList);
  };

  // Function to send list (custom or premade) to backend
  const sendChoice = (choiceType, choice) => {
    if (socket) {
      let message = JSON.stringify({ messageType: choiceType, data: choice });
      socket.send(message);
    }
  };

  // If user selects one of the premade lists, it sends that selection to the backend
  const premadeOptions = listOptions.map((item, index) => (
    <button onClick={() => sendChoice("Premade List", item)} key={index}>
      {item}
    </button>
  ));

  return (
    <div>
      <div>
        <h1>Choose a template list</h1>
        <ul>{premadeOptions}</ul>
      </div>
      <p>
        Create a custom list of items you would like to rank separated by commas
      </p>
      <input onInput={(e) => setUserInput(e.target.value)} type="text" />
      <button onClick={() => createList("Custom List", userInput)}>
        Create List
      </button>
    </div>
  );
}
