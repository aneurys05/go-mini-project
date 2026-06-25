import { Link } from "react-router-dom";
import styles from "../styles/navbar.module.css";

export default function Navbar() {
  return (
    <div className={styles.navbar}>
      <h2>🐾 PetPal</h2>

      <div>
        <Link to="/home">Home</Link>
        <Link to="/profile">Profile</Link>
        <Link to="/login">Logout</Link>
      </div>
    </div>
  );
}