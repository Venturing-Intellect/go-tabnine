package controller

import (
	"customer-feedback/model"
	"customer-feedback/view"
	"database/sql"
	"encoding/json"
	"net/http"
)

type FeedbackController struct {
	DB *sql.DB
}

func (fc *FeedbackController) SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback model.Feedback

	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		view.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := feedback.Validate(); err != nil {
		view.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := feedback.Save(fc.DB); err != nil {
		view.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	view.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Feedback submitted successfully"})
}
