package Controllers

import (
	"AUTH-SERVICE/src/Database"
	"AUTH-SERVICE/src/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
)

var client []Models.Client

func GetClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := []Models.Client{}
	Database.DB.Find(&client)

	json.NewEncoder(w).Encode(client)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	clientID := params["id"]

	// Vérifier si l'ID est vide
	if clientID == "" {
		http.Error(w, "ID du client manquant dans la requête", http.StatusBadRequest)
		return
	}

	client := Models.Client{}
	result := Database.DB.First(&client, clientID)

	if result.Error != nil {

		http.Error(w, "Client non trouvé", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(client)
}
func CreateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Client Models.Client
	err := json.NewDecoder(r.Body).Decode(&Client)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	result := Database.DB.Create(&Client)

	if result.Error != nil {

		//http.Error(w, "Erreur lors de la création du client", http.StatusInternalServerError)
		http.Error(w, "Erreur lors de la création du client: "+result.Error.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Client)
}
