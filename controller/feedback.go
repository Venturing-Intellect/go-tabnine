package controller

import (
	models "customer-feedback/model"
	views "customer-feedback/view"
	"database/sql"
	"encoding/json"
	"net/http"
)

type FeedbackController struct {
	DB *sql.DB
}

func (fc *FeedbackController) SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback
	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := feedback.Validate(); err != nil {
		views.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := feedback.Save(fc.DB); err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	views.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Feedback submitted successfully"})
}
