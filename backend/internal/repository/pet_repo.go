package repository

import (
	"database/sql"
	"petpal/internal/models"
)

type PetRepo struct {
	DB *sql.DB
}

func NewPetRepo(db *sql.DB) *PetRepo {
	return &PetRepo{DB: db}
}

func (r *PetRepo) GetAll() ([]models.Pet, error) {
	rows, err := r.DB.Query("SELECT id,name,species,breed,age,image_url,description FROM pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []models.Pet

	for rows.Next() {
		var p models.Pet
		rows.Scan(&p.ID, &p.Name, &p.Species, &p.Breed, &p.Age, &p.ImageURL, &p.Description)
		pets = append(pets, p)
	}

	return pets, nil
}
