package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

func ValidateToken(next http.Handler) http.Handler {
	fmt.Println("validating...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			fmt.Println("Authorization header missing")
			return
		}

		// Strip the "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Print token for debugging
		fmt.Printf("token: %s\n", tokenString)

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil // Ensure `mySigningKey` is the same key used for signing
		})

		// Check if token parsing returned an error
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			fmt.Println("Invalid token: ", err)
			return
		}

		// Validate claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("Token is valid, email:", claims["email"])
			// Call the next handler
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	})
}
