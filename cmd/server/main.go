package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"radish/internal/healthcheck"

	"github.com/gorilla/mux"
)

var flagConfig = flag.String("config", "./config/local.env", "path to config file")

func main() {
	log.Println("Loading config vars")
	flag.Parse()

	err := godotenv.Load(*flagConfig)

	if err != nil {
		log.Fatal("Failed to load config file")
	}

	log.Println("Starting api server with", os.Getenv("MESSAGE"))
	router := mux.NewRouter()
	router.HandleFunc("/health-check", healthcheck.Handler)

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)

}
