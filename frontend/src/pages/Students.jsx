import { useEffect, useState } from "react";
import { getStudents } from "../Api";
import AddStudent from "./AddStudent";
import TopStudent from "./TopStudent";


function Students({ token }) {
  const [students, setStudents] = useState([]);
  const [error, setError] = useState("");

  const fetchStudents = async () => {
    try {
      const data = await getStudents(token);
      setStudents(data);
    } catch {
      setError("Failed to load students");
    }
  };

  useEffect(() => {
    fetchStudents();
  }, [token]);

  return (
    <div style={{ padding: 20 }}>
      <h2>Student List</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}

      <AddStudent token={token} onStudentAdded={fetchStudents} />
      <TopStudent token={token} />
      <ul>
        {students.map((s) => (
          <li key={s.name}>
            <strong>{s.name}</strong> – Age: {s.age} – Status: {s.status}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Students;