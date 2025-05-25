import { useState } from "react";
import Login from "./pages/Login";
import Students from "./pages/Students";


function App() {
  const [token, setToken] = useState(() => localStorage.getItem("token") || "");

  const logout = () => {
    localStorage.removeItem("token");
    setToken("");
  };
  
  return (
    <div>
      {!token ? (
        <Login setToken={setToken} />
      ) : (
        <>
          <Students token={token} />
          <button onClick={logout} style={{ margin: "20px" }}>
            Logout
          </button>
        </>
      )}
    </div>
  );
}

export default App;
