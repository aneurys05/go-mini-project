package handlers

import (
	"encoding/json"
	"net/http"
	"petpal/internal/repository"
)

type UserHandler struct {
	Repo *repository.UserRepo
}

func NewUserHandler(r *repository.UserRepo) *UserHandler {
	return &UserHandler{Repo: r}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	err := h.Repo.Create(body.Name, body.Email)
	if err != nil {
		http.Error(w, "error creating user", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
