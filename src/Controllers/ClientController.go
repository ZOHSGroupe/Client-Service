package Controllers

import (
	"AUTH-SERVICE/src/Database"
	"AUTH-SERVICE/src/Models"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"gorm.io/gorm"
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

	if clientID == "" {
		http.Error(w, "ID du client manquant dans la requête", http.StatusBadRequest)
		return
	}

	var client Models.Client
	result := Database.DB.First(&client, "id = ?", clientID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Client non trouvé", http.StatusNotFound)
		} else {
			http.Error(w, "Erreur lors de la recherche du client: "+result.Error.Error(), http.StatusInternalServerError)
		}
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

	Client.ID = uuid.New().String()
	result := Database.DB.Create(&Client)

	if result.Error != nil {
		http.Error(w, "Erreur lors de la création du client: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	clientID := vars["id"]

	var updatedClient Models.Client
	err := json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	result := Database.DB.Model(&Models.Client{}).Where("id = ?", clientID).Updates(updatedClient)
	if result.Error != nil {
		http.Error(w, "Error updating client information", http.StatusInternalServerError)
		return
	}

	successMessage := map[string]string{"message": "Client information updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successMessage)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	clientID := params["id"]
	if clientID == "" {
		http.Error(w, "ID du client manquant dans les paramètres de l'URL", http.StatusBadRequest)
		return
	}

	client := Models.Client{ID: clientID}

	result := Database.DB.Delete(&client)

	if result.Error != nil {
		http.Error(w, "Erreur lors de la suppression du client: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Client non trouvé", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Client supprimé avec succès"})
}
