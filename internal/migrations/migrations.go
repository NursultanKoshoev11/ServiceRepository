package migrations

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(migrationsPath, dbURL string) {
	migrateData, err := migrate.New(
		migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(migrateData)

	if err := migrateData.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")
}
