package handlers

import (
	"encoding/json"
	"fmt"
	"jwt_go_server/models"
	"net/http"

	"gorm.io/gorm"
)

// Create a new promoter
// Retrieve a single promoter by ID
func GetAllUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Fetching all users")
		var users []models.User                       // Cambiato a slice di User per ottenere più utenti
		if err := db.Find(&users).Error; err != nil { // Rimosso Preload
			http.Error(w, "Users not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(users) // Encode tutti gli utenti
	}
}
