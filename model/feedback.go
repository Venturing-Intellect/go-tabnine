package model

import (
	"database/sql"
	"errors"
	"regexp"
)

type Feedback struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
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
	stmt, err := db.Prepare("INSERT INTO feedback(name, email, message) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(f.Name, f.Email, f.Message)
	if err != nil {
		return err
	}

	return nil
}
