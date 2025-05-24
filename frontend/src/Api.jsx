const BASE_URL = "https://first-golang-app-production.up.railway.app";

export async function loginUser(credentials) {
  const res = await fetch(`${BASE_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(credentials),
  });

  if (!res.ok) throw new Error("Invalid login");

  return res.json();
}

export async function getStudents(token) {
    const res = await fetch(`${BASE_URL}/students`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  
    if (!res.ok) throw new Error("Unauthorized");
    return res.json();
  }

  export async function addStudent(token, student) {
    const res = await fetch(`${BASE_URL}/students`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(student),
    });
  
    if (!res.ok) throw new Error("Failed to add student");
  }

  export async function getTopStudent(token) {
    const res = await fetch(`${BASE_URL}/top-student`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  
    if (!res.ok) throw new Error("Failed to fetch top student");
    return res.json();
  }
  
  
  