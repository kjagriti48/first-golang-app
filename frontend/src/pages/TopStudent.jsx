import { useEffect, useState } from "react";
import { getTopStudent } from "../Api";

function TopStudent({ token }) {
  const [student, setStudent] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    async function fetchTop() {
      try {
        const data = await getTopStudent(token);
        setStudent(data);
      } catch {
        setError("Unable to load top student");
      }
    }

    fetchTop();
  }, [token]);

  return (
    <div style={{ padding: 20 }}>
      <h3>Top Student</h3>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {student && (
        <div>
          <strong>{student.name}</strong> – Age: {student.age} – Status: {student.status}
        </div>
      )}
    </div>
  );
}

export default TopStudent;
