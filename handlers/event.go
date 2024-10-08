package handlers

import (
	"encoding/json"
	"jwt_go_server/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreateEvent creates a new event
func CreateEvent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var event models.Event
        if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        if err := db.Create(&event).Error; err != nil {
            http.Error(w, "Could not create event", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(event)
    }
}

// GetEvent retrieves an event by ID
func GetEvent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]

        var event models.Event
        if err := db.First(&event, id).Error; err != nil {
            http.Error(w, "Event not found", http.StatusNotFound)
            return
        }

        json.NewEncoder(w).Encode(event)
    }
}

// UpdateEvent updates an existing event
func UpdateEvent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]

        var event models.Event
        if err := db.First(&event, id).Error; err != nil {
            http.Error(w, "Event not found", http.StatusNotFound)
            return
        }

        var updatedEvent models.Event
        if err := json.NewDecoder(r.Body).Decode(&updatedEvent); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        // Update the event fields
        event.Name = updatedEvent.Name
        event.Description = updatedEvent.Description
        event.EventStartDatetime= updatedEvent.EventStartDatetime
        event.EventEndDatetime = updatedEvent.EventEndDatetime
        event.Price = updatedEvent.Price
        event.Location = updatedEvent.Location
        event.Parking = updatedEvent.Parking
        event.DressCode = updatedEvent.DressCode
        event.HotelsNearby = updatedEvent.HotelsNearby
        event.Tags = updatedEvent.Tags

        if err := db.Save(&event).Error; err != nil {
            http.Error(w, "Could not update event", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(event)
    }
}

// DeleteEvent deletes an event by ID
func DeleteEvent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]

        if err := db.Delete(&models.Event{}, id).Error; err != nil {
            http.Error(w, "Event not found", http.StatusNotFound)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}

// ListEvents retrieves all events
func ListEvents(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var events []models.Event
        if err := db.Find(&events).Error; err != nil {
            http.Error(w, "Could not retrieve events", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(events)
    }
}
