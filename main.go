package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rochdi/fizz-buzz/controllers"
	"github.com/rochdi/fizz-buzz/services/fizzbuzz"
	"github.com/rochdi/fizz-buzz/services/health"
)

func main() {

	healthService := health.NewService()
	healthController := controllers.NewHealthController(healthService)
	fizzbuzzService := fizzbuzz.NewService()
	fizzbuzzController := controllers.NewFizzBuzzController(fizzbuzzService, healthService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", healthController.Health).Methods("GET")
	router.HandleFunc("/stats", healthController.Stats).Methods("GET")
	router.HandleFunc("/fizzbuzz", fizzbuzzController.FizzBuzz).Methods("POST") // POST since we are change the server state: updating the stats

	log.Fatal(http.ListenAndServe(":10000", router))
}
