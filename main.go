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
	handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}), // Update this with your specific allowed origins
	)(router)

	fmt.Println("Serveur écoutant sur le port :" + os.Getenv("GO_DOCKER_PORT"))
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("GO_DOCKER_PORT"), router))
}
