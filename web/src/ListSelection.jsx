/* eslint-disable react/prop-types */
import { useEffect, useState } from "react";

export default function ListSelection({ socket }) {
  const [userInput, setUserInput] = useState("");
  const [userList, setUserList] = useState([]);
  const [listOptions, setListOptions] = useState([]);

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

  const createList = () => {
    const formattedList = userInput.split(", ");
    setUserList(formattedList);
  };

  const sendChoice = (choice) => {
    if (socket) {
      let message = JSON.stringify({ messageType: "Choice", data: choice });
      socket.send(message);
    }
  };
  const premadeOptions = listOptions.map((item, index) => (
    <button onClick={() => sendChoice(item)} key={index}>
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
      <button onClick={createList}>Create List</button>
    </div>
  );
}
