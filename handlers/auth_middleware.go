package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"first-rest-api/response"
	"first-rest-api/jwt"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.JSON(w, http.StatusUnauthorized, false, "Missing Authorization header", nil)
			return
		}

		// Must be in format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.JSON(w, http.StatusUnauthorized, false, "Invalid token format", nil)
			return
		}

		tokenString := parts[1]

		// Validate token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, false, "Invalid or expired token", nil)
			return
		}
 
		r.Header.Set("X-User-ID", fmt.Sprintf("%d", claims.UserID))
		r.Header.Set("X-User-Email", claims.Email)

		next.ServeHTTP(w, r)
	}
}
