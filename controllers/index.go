package controllers

import (
	"net/http"
)

type Result struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Haik Shorterner Link With <b>Go</b>,See My Portofolio Web <a href='https://haik.my.id'>Here!</a>"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
