package Routes

import (
	"AUTH-SERVICE/src/Controllers"
	middlewares "AUTH-SERVICE/src/Middlewares"
	"github.com/gorilla/mux"
	_ "net/http"
)

func InitialiserRoutes() *mux.Router {
	router := mux.NewRouter()

	// Définition des routes
	router.HandleFunc("/client", Controllers.GetClients).Methods("GET")
	router.HandleFunc("/client/{id}", Controllers.GetClient).Methods("GET")
	router.HandleFunc("/client", Controllers.CreateClient).Methods("POST")
	router.HandleFunc("/client/{id}", Controllers.UpdateClient).Methods("PUT")
	router.HandleFunc("/client/{id}", Controllers.DeleteClient).Methods("DELETE")
	// Apply the middleware globally for all routes
	router.Use(middlewares.VerifyToken())
	return router
}
