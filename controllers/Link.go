package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/rest-go/db"
)

func GetAllLink(w http.ResponseWriter, r *http.Request) {
	_, err := db.CreateConnection()
	if err != nil {
		w.Header().Set("content-type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": 500,
			"status":     "error",
			"message":    "Error Create Connection To Database!",
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode": 200,
		"status":     "success",
		"message":    "Success Connect to Database!",
	})
}
