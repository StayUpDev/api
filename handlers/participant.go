package handlers

import (
	"encoding/json"
	"jwt_go_server/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateParticipant(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var participant models.Participant
        if err := json.NewDecoder(r.Body).Decode(&participant); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        if err := db.Create(&participant).Error; err != nil {
            http.Error(w, "Could not create participant", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(participant)
    }
}

func GetParticipant(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var participant models.Participant
        if err := db.Preload("Events").First(&participant , vars["id"]).Error; err != nil {
            http.Error(w, "Participant not found", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(participant)
    }
}

func UpdateParticipant(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var particpant models.Participant
        if err := db.First(&particpant.ID, vars["id"]).Error; err != nil {
            http.Error(w, "Participant not found", http.StatusNotFound)
            return
        }

        var updatedParticipant models.Participant
        if err := json.NewDecoder(r.Body).Decode(&updatedParticipant); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        db.Model(&particpant).Updates(updatedParticipant)
        json.NewEncoder(w).Encode(updatedParticipant)
    }
}

func DeleteParticipant(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        if err := db.Delete(&models.Participant{}, vars["id"]).Error; err != nil {
            http.Error(w, "Participant not found", http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    }
}

func GetLikedParticipantEvents(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        var likedParticipantEvents []models.ParticipantEvent

        // Query to get participant events where idParticipant matches and IsLiked is true
        if err := db.Where("id_participant = ? AND is_liked = ?", vars["idParticipant"], true).
            Find(&likedParticipantEvents).Error; err != nil {
            http.Error(w, "No liked participant events found", http.StatusNotFound)
            return
        }

        // Encode the result as JSON and send the response
        json.NewEncoder(w).Encode(likedParticipantEvents)
    }
}
