package repository

import (
	"database/sql"
	"log"
	"servicerepository/config"
	"servicerepository/internal/migrations"
	"servicerepository/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	DeleteUserByID(user_id int64) error
	DeleteUserByEmail(email string) error

	CreateProfile(user_id int64) error
	DeleteProfileByUserID(user_id int64) error

	CreateRole(roleType models.RoleType, roleName string) error
	GetUserByID(user_id int64) (*models.User, error)
	GeUserByEmail(email string) (*models.User, error)
	GetProfileByUserID(user_id int64) (*models.Profile, error)
	GetProfileByEmail(email string) (*models.Profile, error)
}

func ConnectToSql(cfg *config.Config) *sql.DB {

	dbURL := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"
	migrations.RunMigrations("file://./internal/migrations", dbURL)

	dsn := "host=" + cfg.DBHost + " port=" + cfg.DBPort + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " sslmode=disable"
	db, err := sql.Open(cfg.DBDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	connectionerror := db.Ping()

	if connectionerror != nil {
		log.Fatal(err)
	}

	return db
}
