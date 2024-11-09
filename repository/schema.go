package repository

import (
	"database/sql"
)

// InitSchema used for integration test
func InitSchema(db *sql.DB) error {
	createTable := `
	CREATE TABLE IF NOT EXISTS feedback (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		message TEXT NOT NULL
	);`

	_, err := db.Exec(createTable)
	return err
}
