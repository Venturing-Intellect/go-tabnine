package main

import (
	"customer-feedback/controller"
	"customer-feedback/repository"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db := repository.InitDB()
	defer db.Close()

	feedbackController := &controller.FeedbackController{DB: db}

	router := mux.NewRouter()
	router.HandleFunc("/feedback", feedbackController.SubmitFeedback).Methods("POST")

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow requests from the React dev server
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(router)

	http.ListenAndServe(":8080", handler)
}
