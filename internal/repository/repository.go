package repository

import (
	"database/sql"
	"log"
	"servicerepository/config"
	"servicerepository/internal/migrations"
	"servicerepository/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int64) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

func ConnectToSql() *sql.DB {
	cfg := config.LoadConfig()

	dbURL := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"
	migrations.RunMigrations("file://internal/migrations", dbURL)

	dsn := "host=" + cfg.DBHost + " port=" + cfg.DBPort + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " sslmode=disable"
	db, err := sql.Open(cfg.DBDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	connectionerror := db.Ping()

	if connectionerror != nil {
		log.Println(dsn)
		log.Fatal(err)
	}

	return db
}
