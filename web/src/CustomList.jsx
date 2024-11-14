import { useState } from "react";

export default function CustomList() {
  const [userInput, setUserInput] = useState("");
  const [userList, setUserList] = useState([]);

  const createList = () => {
    const formattedList = userInput.split(", ");
    setUserList(formattedList);
};
  return (
    <div>
        <p>Create a custom list of items you would like to rank separated by commas</p>
      <input onInput={(e) => setUserInput(e.target.value)} type="text" />
      <button onClick={createList}>Create List</button>
    </div>
  );
}
