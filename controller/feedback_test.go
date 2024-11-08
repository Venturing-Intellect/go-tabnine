package controller

import (
	"bytes"
	"customer-feedback/model"
	"customer-feedback/repository"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSubmitFeedback(t *testing.T) {
	// Use an in-memory SQLite database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the database schema
	if err := repository.InitSchema(db); err != nil {
		t.Fatalf("Failed to initialize database schema: %v", err)
	}

	fc := &FeedbackController{DB: db}

	feedback := model.Feedback{Email: "test@example.com", Message: "Great service!"}
	payload, _ := json.Marshal(feedback)

	req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fc.SubmitFeedback)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
