package main

import (
	"log"
	"net/http"

	"github.com/A-Big-Moose/go-rest/handlers"
	"github.com/A-Big-Moose/go-rest/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	// Register authentication routes without middleware
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected API routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuthMiddleware)
	handlers.RegisterRoutes(api)

	log.Fatal(http.ListenAndServe(":8080", r))
}
