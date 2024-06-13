package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/resource", getResource).Methods("GET")
	r.HandleFunc("/resource", createResource).Methods("POST")
	r.HandleFunc("/resource/{id}", updateResource).Methods("PUT")
	r.HandleFunc("/resource/{id}", deleteResource).Methods("DELETE")

	r.HandleFunc("/home", Home).Methods("GET")
}

func getResource(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get resource"))
}

func createResource(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create resource"))
}

func updateResource(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update resource"))
}

func deleteResource(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete resource"))
}
