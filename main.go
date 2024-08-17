package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HaikalRFadhilahh/rest-go/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	godotenv.Load()

	// Mux API Routung
	router := mux.NewRouter()

	// Routing
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/link", controllers.GetAllLink).Methods("GET")

	// Serve Http Server
	host := "0.0.0.0"
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Print("Application Running On", host, ":", port)
	err := http.ListenAndServe(host+":"+port, router)
	if err != nil {
		log.Fatalln(err)
	}
}
