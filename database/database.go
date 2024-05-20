package database

import (
	"os"
	"path/filepath"

	"database/sql"

	_ "modernc.org/sqlite"
)

func Database() (*sql.DB, error) {
	dbDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dbPath := filepath.Join(dbDir, "Data", "DJ", "dj.db")
	db, err := sql.Open("sqlite", dbPath)
	return db, err
}
