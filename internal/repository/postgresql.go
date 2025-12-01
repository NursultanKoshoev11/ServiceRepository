package repository

import (
	"database/sql"
	"log"
	"servicerepository/internal/models"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) DeleteUserByID(user_id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", user_id)
	log.Println(err)
	return err
}
func (r *PostgresUserRepository) DeleteUserByEmail(email string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE email = $1", email)
	log.Println(err)
	return err
}
func (r *PostgresUserRepository) DeleteProfileByUserID(user_id int64) error {
	_, err := r.db.Exec("DELETE FROM profiles WHERE user_id = $1", user_id)
	log.Println(err)
	return err
}

func (r *PostgresUserRepository) CreateUser(user *models.User) error {
	err := r.db.QueryRow("INSERT INTO users (email, password, role_id) VALUES ($1, $2, $3) RETURNING id", user.Email, user.Password, user.RoleID).Scan(&user.ID)
	log.Println(err)
	return err
}

func (r *PostgresUserRepository) CreateProfile(user_id int64) error {
	_, err := r.db.Exec("INSERT INTO profiles (user_id) VALUES ($1)", user_id)
	log.Println(err)
	return err
}

func (r *PostgresUserRepository) CreateRole(roleType models.RoleType, roleName string) error {
	_, err := r.db.Exec("INSERT INTO roles (id,name) VALUES ($1,$2)", roleType, roleName)
	log.Println(err)
	return err
}

func (r *PostgresUserRepository) GetUserByID(user_id int64) (*models.User, error) {
	var user models.User
	row := r.db.QueryRow("SELECT id, email, password, role_id FROM users WHERE id=$1", user_id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.RoleID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GeUserByEmail(email string) (*models.User, error) {
	var user models.User
	row := r.db.QueryRow("SELECT id, email, password, role_id FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.RoleID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetProfileByUserID(user_id int64) (*models.Profile, error) {
	var name, bio, avatar sql.NullString

	var profile models.Profile
	err := r.db.QueryRow(`SELECT profile.id, profile.user_id, profile.name, profile.avatar, profile.bio, profile.created_at FROM profiles profile WHERE profile.user_id = $1`, user_id).Scan(
		&profile.ID,
		&profile.UserID,
		&name,
		&avatar,
		&bio,
		&profile.CreatedAt)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	profile.Avatar = avatar.String
	profile.Bio = bio.String
	profile.Name = name.String

	return &profile, nil
}

func (r *PostgresUserRepository) GetProfileByEmail(email string) (*models.Profile, error) {
	var name, bio, avatar sql.NullString
	var profile models.Profile
	err := r.db.QueryRow(`SELECT profile.id, profile.user_id, profile.name, profile.avatar, profile.bio, profile.created_at	FROM users u 
		JOIN profiles profile ON profile.user_id = u.id WHERE u.email = $1`, email).Scan(
		&profile.ID,
		&profile.UserID,
		&name,
		&avatar,
		&bio,
		&profile.CreatedAt)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	profile.Avatar = avatar.String
	profile.Bio = bio.String
	profile.Name = name.String

	return &profile, nil
}
