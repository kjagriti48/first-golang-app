import { useState } from "react";
import Login from "./pages/Login";
import Students from "./pages/Students";

function App() {
  const [token, setToken] = useState("");

  return (
    <div>
      {!token ? (
        <Login setToken={setToken} />
      ) : (
        <Students token={token} />
      )}
    </div>
  );
}

export default App;
