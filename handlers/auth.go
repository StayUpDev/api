package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"jwt_go_server/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var mySigningKey = []byte("secret")

// Struct for JWT claims
type Claims struct {
	Email  string `json:"email"`
	UserID uint   `json:"user_id"`
	jwt.StandardClaims
}

// Register a new user
func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Email == "" || user.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login an existing user
func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		fmt.Println("Invalid input")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {

		fmt.Println("Invalid credentials 01")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {

		fmt.Println("Invalid credentials 02")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	claims := Claims{
		Email:  user.Email,
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {

		fmt.Println("Could not create token")
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}
	fmt.Println("user found")
	fmt.Println(user.Role)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString, "role": existingUser.Role})
}
