package model

import (
	"database/sql"
	"errors"
	"regexp"
)

type Feedback struct {
	ID      int
	Email   string
	Message string
}

func (f *Feedback) Validate() error {
	if f.Email == "" || f.Message == "" {
		return errors.New("email and message are required")
	}

	// Validate email format
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(f.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

func (f *Feedback) Save(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO feedback(email, message) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(f.Email, f.Message)
	return err
}
