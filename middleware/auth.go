package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func (m *AuthMiddleware) JWTMiddleware(w http.ResponseWriter, r *http.Request) {
	// Define your secret key, which should be kept secret in a real application
	secretKey := []byte("ef74HBH3nf34x34ry7HBSAsdaaDMXdasdUHNghn327zr2")

	// Parse the JWT token from the request header
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
    
}
