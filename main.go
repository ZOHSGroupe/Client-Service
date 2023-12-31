package main

import (
	"AUTH-SERVICE/src/Database"
	"AUTH-SERVICE/src/Routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	Database.InitializeDatabase()
	router := Routes.InitialiserRoutes()
	fmt.Println("Serveur écoutant sur le port :" + os.Getenv("GO_DOCKER_PORT"))
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("GO_DOCKER_PORT"), router))
}
