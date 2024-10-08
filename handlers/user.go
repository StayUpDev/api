package handlers

import (
	"encoding/json"
	"fmt"
	"jwt_go_server/models"
	"net/http"

	"github.com/gorilla/mux"
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

func GetUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var user models.User
		if err := db.First(&user, vars["id"]).Error; err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
