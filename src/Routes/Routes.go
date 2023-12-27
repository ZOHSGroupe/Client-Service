package Routes

import (
	"AUTH-SERVICE/src/Controllers"
	"github.com/gorilla/mux"
	_ "net/http"
)

func InitialiserRoutes() *mux.Router {
	router := mux.NewRouter()

	// Définition des routes
	router.HandleFunc("/clients", Controllers.GetClients).Methods("GET")
	router.HandleFunc("/client/{id}", Controllers.GetClient).Methods("GET")
	router.HandleFunc("/client", Controllers.CreateClient).Methods("POST")

	return router
}
