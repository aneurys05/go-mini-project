import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { getPets } from "../api/pets";

export default function PetDetails() {
  const { id } = useParams();
  const [pet, setPet] = useState(null);

  useEffect(() => {
    getPets().then((pets) => {
      setPet(pets.find((p) => p.id == id));
    });
  }, [id]);

  if (!pet) return <p>Loading...</p>;

  return (
    <div style={{ padding: 40 }}>
      <h1>{pet.name}</h1>
      <p>{pet.breed}</p>
      <p>{pet.description}</p>
    </div>
  );
}