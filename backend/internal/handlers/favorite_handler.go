package handlers

import (
	"encoding/json"
	"net/http"
	"petpal/internal/repository"
)

type FavoriteHandler struct {
	Repo *repository.FavoriteRepo
}

func NewFavoriteHandler(r *repository.FavoriteRepo) *FavoriteHandler {
	return &FavoriteHandler{Repo: r}
}

func (h *FavoriteHandler) Add(w http.ResponseWriter, r *http.Request) {
	var body struct {
		UserID int `json:"user_id"`
		PetID  int `json:"pet_id"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	h.Repo.Add(body.UserID, body.PetID)
	w.WriteHeader(http.StatusCreated)
}
