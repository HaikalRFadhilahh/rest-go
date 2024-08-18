package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/rest-go/db"
	"github.com/HaikalRFadhilahh/rest-go/models"
)

func GetAllLink(w http.ResponseWriter, r *http.Request) {
	db, err := db.CreateConnection()

	if err != nil {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": 500,
			"status":     "error",
			"message":    "Error Create Connection To Database!",
		})
	}

	defer db.Close()

	res, err := db.Query("select * from links")
	if err != nil {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": 500,
			"status":     "error",
			"message":    "Internal Server Error!",
		})
	}
	var data []models.Links
	for res.Next() {
		var links models.Links
		err := res.Scan(&links.ID, &links.Alias, &links.Url, &links.Created_at, &links.Updated_at)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, links)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode": 200,
		"status":     "success",
		"message":    "Data Link",
		"data":       data,
	})
}
