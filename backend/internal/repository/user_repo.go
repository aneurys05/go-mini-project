package repository

import (
	"database/sql"
	"petpal/internal/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(name, email string) error {
	_, err := r.DB.Exec(
		"INSERT INTO users(name,email) VALUES($1,$2)",
		name, email,
	)
	return err
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	row := r.DB.QueryRow(
		"SELECT id,name,email FROM users WHERE email=$1",
		email,
	)

	u := &models.User{}
	err := row.Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}
