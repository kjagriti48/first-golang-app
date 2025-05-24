import { useState } from "react";
import { addStudent } from "../api";

function AddStudent({ token, onStudentAdded }) {
  const [form, setForm] = useState({
    name: "",
    age: "",
    marks: { math: "", science: "" },
  });
  const [error, setError] = useState("");

  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name in form.marks) {
      setForm({ ...form, marks: { ...form.marks, [name]: Number(value) } });
    } else {
      setForm({ ...form, [name]: value });
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    try {
      await addStudent(token, {
        name: form.name,
        age: Number(form.age),
        marks: form.marks,
      });
      onStudentAdded(); // refresh list
    } catch {
      setError("Failed to add student");
    }
  };

  return (
    <form onSubmit={handleSubmit} style={{ padding: 20 }}>
      <h3>Add Student</h3>
      <input name="name" placeholder="Name" value={form.name} onChange={handleChange} />
      <br /><br />
      <input name="age" placeholder="Age" value={form.age} onChange={handleChange} />
      <br /><br />
      <input name="math" placeholder="Math Marks" value={form.marks.math} onChange={handleChange} />
      <br /><br />
      <input name="science" placeholder="Science Marks" value={form.marks.science} onChange={handleChange} />
      <br /><br />
      <button type="submit">Add Student</button>
      {error && <p style={{ color: "red" }}>{error}</p>}
    </form>
  );
}

export default AddStudent;
