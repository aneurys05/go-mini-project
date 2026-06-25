import { Link } from "react-router-dom";
import styles from "../styles/petcard.module.css";

export default function PetCard({ pet }) {
  return (
    <div className={styles.card}>
      <h3>🐶 {pet.name}</h3>
      <p>{pet.breed}</p>
      <p>{pet.age} years old</p>

      <Link to={`/pet/${pet.id}`}>View Details</Link>
    </div>
  );
}