package handlers

import (
	"encoding/json"
	"jwt_go_server/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Create a new promoter
func CreatePromoter(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var promoter models.Promoter
        if err := json.NewDecoder(r.Body).Decode(&promoter); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        if err := db.Create(&promoter).Error; err != nil {
            http.Error(w, "Could not create promoter", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(promoter)
    }
}

// Retrieve a single promoter by ID
func GetPromoter(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var promoter models.Promoter
        if err := db.Preload("Events").First(&promoter, vars["id"]).Error; err != nil {
            http.Error(w, "Promoter not found", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(promoter)
    }
}

// Update a promoter
func UpdatePromoter(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var promoter models.Promoter
        if err := db.First(&promoter, vars["id"]).Error; err != nil {
            http.Error(w, "Promoter not found", http.StatusNotFound)
            return
        }

        var updatedPromoter models.Promoter
        if err := json.NewDecoder(r.Body).Decode(&updatedPromoter); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        db.Model(&promoter).Updates(updatedPromoter)
        json.NewEncoder(w).Encode(promoter)
    }
}

// Delete a promoter
func DeletePromoter(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        if err := db.Delete(&models.Promoter{}, vars["id"]).Error; err != nil {
            http.Error(w, "Promoter not found", http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    }
}

// Fetch previous events for a promoter
func GetPreviousEvents(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var events []models.Event
        currentTime := time.Now()

        if err := db.Where("promoter_id = ? AND event_end_datetime < ?", vars["id"], currentTime).Find(&events).Error; err != nil {
            http.Error(w, "Error fetching events", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(events)
    }
}

// Fetch current events for a promoter
func GetCurrentEvents(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var events []models.Event
        currentTime := time.Now()

        if err := db.Where("promoter_id = ? AND event_start_datetime <= ? AND event_end_datetime >= ?", vars["id"], currentTime, currentTime).Find(&events).Error; err != nil {
            http.Error(w, "Error fetching events", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(events)
    }
}

// Fetch scheduled events for a promoter (future events)
func GetScheduledEvents(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var events []models.Event
        currentTime := time.Now()

        if err := db.Where("promoter_id = ? AND event_start_datetime > ?", vars["id"], currentTime).Find(&events).Error; err != nil {
            http.Error(w, "Error fetching events", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(events)
    }
}