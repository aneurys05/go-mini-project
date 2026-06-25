import { Link } from "react-router-dom";

export default function Landing() {
  return (
    <div style={styles.wrapper}>
      <h1 style={styles.title}>🐾 PetPal</h1>

      <p style={styles.subtitle}>
        Find your new best friend. Adopt, love, repeat.
      </p>

      <div style={styles.buttons}>
        <Link to="/login" style={styles.btnPrimary}>
          Login
        </Link>

        <Link to="/signup" style={styles.btnSecondary}>
          Sign Up
        </Link>
      </div>
    </div>
  );
}

const styles = {
  wrapper: {
    height: "100vh",
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center",
    background: "#fff8f0",
    textAlign: "center",
  },
  title: {
    fontSize: "48px",
    color: "#8b5e3c",
    marginBottom: "10px",
  },
  subtitle: {
    fontSize: "18px",
    color: "#555",
    marginBottom: "30px",
  },
  buttons: {
    display: "flex",
    gap: "15px",
  },
  btnPrimary: {
    padding: "10px 20px",
    background: "#8b5e3c",
    color: "white",
    borderRadius: "10px",
    textDecoration: "none",
  },
  btnSecondary: {
    padding: "10px 20px",
    background: "#A8D5BA",
    color: "#333",
    borderRadius: "10px",
    textDecoration: "none",
  },
};