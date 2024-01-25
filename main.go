package main

import (
	"AUTH-SERVICE/src/Database"
	"AUTH-SERVICE/src/Routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hudl/fargo"
)

func main() {
	// Initialize Database
	Database.InitializeDatabase()

	// Initialize routes
	router := Routes.InitialiserRoutes()

	// Print message about server listening
	fmt.Println("Serveur écoutant sur le port :" + os.Getenv("GO_DOCKER_PORT"))

	// Register with Eureka
	eurekaClient := fargo.NewConn("http://localhost:8761/eureka/v2")
	instance := &fargo.Instance{
		App:      "client-service", // Remplacez par le nom de votre application
		Port:     8000,             // Port sur lequel votre service écoute dans le conteneur Docker
		HostName: "localhost",      // Nom d'hôte de votre service
		Status:   fargo.UP,         // Status du service
	}
	if err := eurekaClient.RegisterInstance(instance); err != nil {
		log.Fatalf("Failed to register with Eureka: %v", err)
	}

	// Start the server
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("GO_DOCKER_PORT"), router))
}
