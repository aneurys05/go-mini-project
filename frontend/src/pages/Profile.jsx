import { useContext } from "react";
import { UserContext } from "../context/UserContext";

export default function Profile() {
  const { user } = useContext(UserContext);

  return (
    <div style={{ padding: 40 }}>
      <h1>Profile 🐾</h1>

      {user ? (
        <>
          <p>Name: {user.name}</p>
          <p>Email: {user.email}</p>
        </>
      ) : (
        <p>No user logged in</p>
      )}
    </div>
  );
}