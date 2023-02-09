package main

import (
	"log"
	"machine-test-1/routes"
	"machine-test-1/usecase"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".ENV ERROR - ", err)
	}

	usecase.NewDbInstance()

	mux := mux.NewRouter()

	routes.InitRoutes(mux)

	log.Println("Starting mail service on port :3000")
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal("SERVER ERROR - ", err)
	}
}
