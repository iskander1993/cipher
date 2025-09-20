package middleware

import (
	"ave_project/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type contextKey string

var userIDKey = contextKey("userID")

// AuthMiddleware проверяет JWT токен
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// TODO: Переписать без " "
		//  Использовать trim или replace
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		userID, err := jwtpkg.ParseToken(parts[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// кладём userID в context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID достаёт userID из context
func GetUserID(r *http.Request) int {
	if v := r.Context().Value(userIDKey); v != nil {
		return v.(int)
	}
	return 0
}
