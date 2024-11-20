/* eslint-disable react/prop-types */
import { useContext, useEffect, useState } from "react";
import { WebSocketContext } from "./websocket";

export default function ListSelection() {
  const socket = useContext(WebSocketContext);

  const [userInput, setUserInput] = useState("");
  const [listOptions, setListOptions] = useState([]);

  // Wait for premade list options from backend
  useEffect(() => {
    if (socket) {
      socket.onmessage = (event) => {
        const message = JSON.parse(event.data);
        if (message.messageType == "List Options") {
          setListOptions(message.options);
        }
      };
    }
  }, [socket]);

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
