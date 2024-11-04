import { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [message, setMessage] = useState(null);
  
  useEffect(() => {
    axios
      .get("http://localhost:8080/message")
      .then((response) => {
        setMessage(response.data.text);
      })
      .catch((error) => console.error("Error fetching data:", error));
  }, []);

  return (
    <>
      <h1>{message}</h1>
    </>
  );
}

export default App;
