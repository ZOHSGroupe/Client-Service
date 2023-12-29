package Controllers

import (
	"AUTH-SERVICE/src/Database"
	"AUTH-SERVICE/src/Models"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
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

	// Generate a new UUID for the client
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	clientID := params["id"]

	// Vérifier si l'ID est vide
	if clientID == "" {
		http.Error(w, "ID du client manquant dans la requête", http.StatusBadRequest)
		return
	}

	var updatedClient Models.Client
	err := json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Mettez à jour le client dans la base de données
	result := Database.DB.Model(&Models.Client{}).Where("id = ?", clientID).Updates(updatedClient)

	if result.Error != nil {
		http.Error(w, "Erreur lors de la mise à jour du client: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedClient)
}
func DeleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the client ID from the URL parameter
	clientID := chi.URLParam(r, "id")
	if clientID == "" {
		http.Error(w, "ID du client manquant dans les paramètres de l'URL", http.StatusBadRequest)
		return
	}

	// Supprimer le client de la base de données
	result := Database.DB.Delete(&Models.Client{}, clientID)

	if result.Error != nil {
		http.Error(w, "Erreur lors de la suppression du client: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		http.Error(w, "Client non trouvé", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Client supprimé avec succès"})
}
