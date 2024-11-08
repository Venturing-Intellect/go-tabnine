package model

import (
	"testing"
)

func TestFeedbackValidation(t *testing.T) {
	tests := []struct {
		email   string
		message string
		valid   bool
	}{
		{"test@example.com", "This is a feedback message.", true},
		{"invalid-email", "This is a feedback message.", false},
		{"", "This is a feedback message.", false},
		{"test@example.com", "", false},
	}

	for _, test := range tests {
		feedback := Feedback{Email: test.email, Message: test.message}
		err := feedback.Validate()
		if (err == nil) != test.valid {
			t.Errorf("expected validity: %v, got error: %v", test.valid, err)
		}
	}
}
