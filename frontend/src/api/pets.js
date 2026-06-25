import { API_URL } from "./base";

export async function getPets() {
  const res = await fetch(`${API_URL}/pets`);
  return res.json();
}