package Controllers

import (
	"AUTH-SERVICE/src/Database"
	"AUTH-SERVICE/src/Models"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"gorm.io/gorm"
)

var client []Models.Client

func GetTest(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("client service work")
}
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

	var newClient Models.Client
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Check if a client with the same email exists
	existingEmailClient := Models.Client{}
	if err := Database.DB.Where("email = ?", newClient.Email).First(&existingEmailClient).Error; err == nil {
		http.Error(w, "A client with the same email already exists", http.StatusBadRequest)
		return
	}

	// Check if a client with the same national ID exists
	existingNationalIDClient := Models.Client{}
	if err := Database.DB.Where("national_id = ?", newClient.NationalID).First(&existingNationalIDClient).Error; err == nil {
		http.Error(w, "A client with the same national ID already exists", http.StatusBadRequest)
		return
	}

	// Initialize CreateDate and LastModificationDate with the current date
	currentDate := time.Now().Format("2006-01-02")
	newClient.CreateDate = currentDate
	newClient.LastModificationDate = currentDate

	// Generate a new UUID for the client
	newClient.ID = uuid.New().String()

	// Create the new client in the database
	result := Database.DB.Create(&newClient)
	if result.Error != nil {
		http.Error(w, "Error creating the client: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the newly created client
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newClient)
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
