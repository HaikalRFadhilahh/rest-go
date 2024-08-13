package controllers

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Return Json
	Res := Result{
		StatusCode: 200,
		Status:     "success",
		Message:    "RESTFull API With Go",
	}
	// Handle Error
	err := json.NewEncoder(w).Encode(Res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
