import { useState } from "react";

export default function Signup() {
  const [name, setName] = useState("");

  return (
    <div style={{ padding: 40 }}>
      <h1>Sign Up</h1>

      <input placeholder="name" onChange={(e) => setName(e.target.value)} />

      <input placeholder="email" />

      <button>Create Account</button>
    </div>
  );
}