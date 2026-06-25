import { API_URL } from "./base";

export async function addFavorite(user_id, pet_id) {
  return fetch(`${API_URL}/favorites`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ user_id, pet_id }),
  });
}