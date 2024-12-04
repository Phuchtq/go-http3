package routes

import (
	"http3-integrate/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializePetApi() *mux.Router {
	var router = mux.NewRouter().PathPrefix("/pets").Subrouter()

	// Get methods
	router.HandleFunc("", handlers.GetAllPets).Methods(http.MethodGet)
	router.HandleFunc("/keyword/{keyword}", handlers.GetPetsByKeyword).Methods(http.MethodGet)
	router.HandleFunc("/{id}", handlers.GetPetById).Methods(http.MethodGet)

	// Post methods
	router.HandleFunc("", handlers.CreatePet).Methods(http.MethodPost)

	// Put methods
	router.HandleFunc("", handlers.EditPet).Methods(http.MethodPut)

	// Delete methods
	router.HandleFunc("/{id}", handlers.RemovePet).Methods(http.MethodDelete)

	return router
}
