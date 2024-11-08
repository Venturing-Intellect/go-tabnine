package repository

import (
	"database/sql"
)

func InitSchema(db *sql.DB) error {
	createTable := `
	CREATE TABLE IF NOT EXISTS feedback (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		message TEXT NOT NULL
	);`

	_, err := db.Exec(createTable)
	return err
}
