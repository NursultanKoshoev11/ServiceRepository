package repository

import (
	"database/sql"
	"servicerepository/internal/models"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *models.User) error {
	err := r.db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.Password).Scan(&user.ID)
	return err
}

func (r *PostgresUserRepository) GetByID(id int64) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE id=$1", id)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email=$1", email)
	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
