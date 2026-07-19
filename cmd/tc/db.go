package main

import (
	"database/sql"

	sqliterepo "termcode/internal/infrastructure/database/sqlite"
)

func openDB() (*sql.DB, error) {
	return sqliterepo.Open()
}

func runMigrations(db *sql.DB) error {
	return sqliterepo.RunMigrations(db)
}

func openProviderRepo(db *sql.DB) *sqliterepo.ProviderRepo {
	return sqliterepo.NewProviderRepo(db)
}
