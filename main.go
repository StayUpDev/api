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

	// Protect users routes
	router.Handle("/api/users", middleware.ValidateToken(http.HandlerFunc(handlers.GetAllUsers(db)))).Methods(http.MethodGet)
	router.Handle("/api/users/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.GetUser(db)))).Methods(http.MethodGet)

	// Promoter routes with validation
	router.Handle("/api/promoters", middleware.ValidateToken(http.HandlerFunc(handlers.CreatePromoter(db)))).Methods(http.MethodGet)
	router.Handle("/api/promoters/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.GetPromoter(db)))).Methods(http.MethodGet)
	router.Handle("/api/promoters/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.UpdatePromoter(db)))).Methods(http.MethodPut)
	router.Handle("/api/promoters/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.DeletePromoter(db)))).Methods(http.MethodDelete)

	router.Handle("/api/promoters/{id}/events/previous", middleware.ValidateToken(http.HandlerFunc(handlers.GetPreviousEvents(db)))).Methods(http.MethodGet)
	router.Handle("/api/promoters/{id}/events/current", middleware.ValidateToken(http.HandlerFunc(handlers.GetCurrentEvents(db)))).Methods(http.MethodGet)
	router.Handle("/api/promoters/{id}/events/scheduled", middleware.ValidateToken(http.HandlerFunc(handlers.GetScheduledEvents(db)))).Methods(http.MethodGet)

	// Participant routes with validation
	router.Handle("/api/participants", middleware.ValidateToken(http.HandlerFunc(handlers.CreateParticipant(db)))).Methods(http.MethodGet)
	router.Handle("/api/participants/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.GetParticipant(db)))).Methods(http.MethodGet)
	router.Handle("/api/participants/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.UpdateParticipant(db)))).Methods(http.MethodPut)
	router.Handle("/api/participants/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.DeleteParticipant(db)))).Methods(http.MethodDelete)

	// Event routes with validation
	router.Handle("/api/events", middleware.ValidateToken(http.HandlerFunc(handlers.CreateEvent(db)))).Methods(http.MethodGet)
	router.Handle("/api/events/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.GetEvent(db)))).Methods(http.MethodGet)
	router.Handle("/api/events/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.UpdateEvent(db)))).Methods(http.MethodPut)
	router.Handle("/api/events/{id}", middleware.ValidateToken(http.HandlerFunc(handlers.DeleteEvent(db)))).Methods(http.MethodDelete)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
