package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/rest-go/db"
	"github.com/HaikalRFadhilahh/rest-go/errors"
	"github.com/HaikalRFadhilahh/rest-go/models"
	"github.com/gorilla/mux"
)

type response struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func responseLink(w http.ResponseWriter, statusCode int, status string, message string, data any) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	res := response{
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
		Data:       data,
	}

	json.NewEncoder(w).Encode(res)
}

func GetAllLink(w http.ResponseWriter, r *http.Request) {
	db, err := db.CreateConnection()

	if err != nil {
		errors.ResponseError(w, http.StatusInternalServerError, "Cannot Connect Database!")
		return
	}

	defer db.Close()

	res, err := db.Query("select * from links")
	if err != nil {
		errors.ResponseError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	var data []models.Links
	for res.Next() {
		var links models.Links
		err := res.Scan(&links.ID, &links.Alias, &links.Url, &links.Created_at, &links.Updated_at)
		if err != nil {
			errors.ResponseError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		data = append(data, links)
	}

	responseLink(w, http.StatusOK, "success", "Data All Links", data)
}

func GoToUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := db.CreateConnection()
	if err != nil {
		errors.ResponseError(w, http.StatusInternalServerError, "Cannot Connect Database!")
		return
	}
	defer db.Close()

	var data models.Links
	err = db.QueryRow("select * from links as l where l.alias = ?", params["alias"]).Scan(&data.ID, &data.Alias, &data.Url, &data.Created_at, &data.Updated_at)
	if err != nil {
		errors.ResponseError(w, http.StatusNotFound, "Link Not Found!")
		return
	}

	http.Redirect(w, r, data.Url, http.StatusSeeOther)
}

func AddLinks(w http.ResponseWriter, r *http.Request) {
	db, err := db.CreateConnection()
	if err != nil {
		errors.ResponseError(w, 500, "Cannot Connect Database")
	}

	defer db.Close()

	var link models.Links

	err = json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		errors.ResponseError(w, 500, "Internal Server Error!")
		return
	}

	if link.Alias == "" && link.Url == "" {
		errors.ResponseError(w, 400, "Invalid Request!!")
		return
	}

	_, err = db.Exec("insert into links(alias,url) value (?,?)", link.Alias, link.Url)
	if err != nil {
		errors.ResponseError(w, 500, "Internal Server Error!")
		return
	}

	responseLink(w, 200, "success", "Data Success Added!", nil)
}
