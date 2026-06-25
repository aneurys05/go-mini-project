package router

import (
	"database/sql"
	"net/http"

	"petpal/internal/handlers"
	"petpal/internal/repository"
)

func New(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	userRepo := repository.NewUserRepo(db)
	petRepo := repository.NewPetRepo(db)
	favRepo := repository.NewFavoriteRepo(db)

	userHandler := handlers.NewUserHandler(userRepo)
	petHandler := handlers.NewPetHandler(petRepo)
	favHandler := handlers.NewFavoriteHandler(favRepo)

	mux.HandleFunc("/api/users", userHandler.CreateUser)
	mux.HandleFunc("/api/pets", petHandler.GetAll)
	mux.HandleFunc("/api/favorites", favHandler.Add)

	return mux
}
