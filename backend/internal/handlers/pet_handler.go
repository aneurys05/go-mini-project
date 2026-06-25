package handlers

import (
	"encoding/json"
	"net/http"
	"petpal/internal/repository"
)

type PetHandler struct {
	Repo *repository.PetRepo
}

func NewPetHandler(r *repository.PetRepo) *PetHandler {
	return &PetHandler{Repo: r}
}

func (h *PetHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pets, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "error fetching pets", 500)
		return
	}

	json.NewEncoder(w).Encode(pets)
}
