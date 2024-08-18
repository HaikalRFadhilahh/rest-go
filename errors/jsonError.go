package errors

import (
	"encoding/json"
	"net/http"
)

type errorsStruct struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	ErrorResponse := errorsStruct{
		StatusCode: statusCode,
		Status:     "error",
		Message:    message,
	}

	json.NewEncoder(w).Encode(ErrorResponse)
}
