package main

import (
	"jwt_go_server/config"
	"jwt_go_server/handlers"
	"jwt_go_server/middleware"
	"jwt_go_server/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.InitDB()

	// Migrate the user table
	// Migrate the user table
	if err := models.Migrate(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(db, w, r)
	}).Methods("POST")

	router.Handle("/api/protected", middleware.ValidateToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You are authorized!"))
	}))).Methods("GET")

	router.HandleFunc("/api/users", handlers.GetAllUsers(db)).Methods(http.MethodGet)
	// Promoter
	router.HandleFunc("/api/promoters", handlers.CreatePromoter(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/promoters/{id}", handlers.GetPromoter(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/promoters/{id}", handlers.UpdatePromoter(db)).Methods(http.MethodPut)
	router.HandleFunc("/api/promoters/{id}", handlers.DeletePromoter(db)).Methods(http.MethodDelete)

	router.HandleFunc("/api/promoters/{id}/events/previous", handlers.GetPreviousEvents(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/promoters/{id}/events/current", handlers.GetCurrentEvents(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/promoters/{id}/events/scheduled", handlers.GetScheduledEvents(db)).Methods(http.MethodGet)

	// Participant
	router.HandleFunc("/api/participants", handlers.CreateParticipant(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/participants/{id}", handlers.GetParticipant(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/participants/{id}", handlers.UpdateParticipant(db)).Methods(http.MethodPut)
	router.HandleFunc("/api/participants/{id}", handlers.DeleteParticipant(db)).Methods(http.MethodDelete)

	router.HandleFunc("/api/events", handlers.CreateEvent(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/events/{id}", handlers.GetEvent(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/events/{id}", handlers.UpdateEvent(db)).Methods(http.MethodPut)
	router.HandleFunc("/api/events/{id}", handlers.DeleteEvent(db)).Methods(http.MethodDelete)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
