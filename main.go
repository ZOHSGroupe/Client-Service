package main

import (
	"AUTH-SERVICE/src/Database"
	"AUTH-SERVICE/src/Routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {

	Database.InitializeDatabase()
	router := Routes.InitialiserRoutes()
	// Enable CORS for all routes
	// Set up CORS middleware
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// Start the server with CORS middleware
	http.ListenAndServe(":"+os.Getenv("GO_DOCKER_PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router))

	fmt.Println("Serveur écoutant sur le port :" + os.Getenv("GO_DOCKER_PORT"))
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("GO_DOCKER_PORT"), router))
}
