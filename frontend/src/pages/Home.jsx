import { useEffect, useState } from "react";
import { getPets } from "../api/pets";
import PetCard from "../components/PetCard";

export default function Home() {
  const [pets, setPets] = useState([]);

  useEffect(() => {
    getPets().then(setPets);
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h1>Available Pets 🐾</h1>

      <div style={{ display: "grid", gridTemplateColumns: "repeat(3,1fr)", gap: "20px" }}>
        {pets.map((p) => (
          <PetCard key={p.id} pet={p} />
        ))}
      </div>
    </div>
  );
}