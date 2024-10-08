package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

// Middleware to validate the JWT token
func ValidateToken(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return mySigningKey, nil
        })

        if err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            fmt.Println(claims["email"]) // Access claims here if needed
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
        }
    })
}
