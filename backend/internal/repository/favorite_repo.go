package repository

import "database/sql"

type FavoriteRepo struct {
	DB *sql.DB
}

func NewFavoriteRepo(db *sql.DB) *FavoriteRepo {
	return &FavoriteRepo{DB: db}
}

func (r *FavoriteRepo) Add(userID, petID int) error {
	_, err := r.DB.Exec(
		"INSERT INTO favorites(user_id,pet_id) VALUES($1,$2)",
		userID, petID,
	)
	return err
}

func (r *FavoriteRepo) GetUserFavorites(userID int) (*sql.Rows, error) {
	return r.DB.Query(`
		SELECT p.id,p.name,p.species,p.breed,p.age,p.image_url,p.description
		FROM pets p
		JOIN favorites f ON f.pet_id = p.id
		WHERE f.user_id=$1
	`, userID)
}
