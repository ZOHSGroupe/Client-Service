package Routes

import (
	"AUTH-SERVICE/src/Controllers"
	_ "net/http"

	"github.com/gorilla/mux"
)

func InitialiserRoutes() *mux.Router {
	router := mux.NewRouter()

	// Définition des routes
	router.HandleFunc("/", Controllers.GetTest).Methods("GET")
	router.HandleFunc("/client", Controllers.GetClients).Methods("GET")
	router.HandleFunc("/client/{id}", Controllers.GetClient).Methods("GET")
	router.HandleFunc("/client", Controllers.CreateClient).Methods("POST")
	router.HandleFunc("/client/{id}", Controllers.UpdateClient).Methods("PUT")
	router.HandleFunc("/client/{id}", Controllers.DeleteClient).Methods("DELETE")
	// Apply the middleware globally for all routes
	// router.Use(middlewares.VerifyToken())
	return router
}
