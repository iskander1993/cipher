// internal/middleware/auth.go
package middleware

import (
	"ave_project/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type contextKey string

var userIDKey = contextKey("userID")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Fields(authHeader) // автоматически разделяет по пробелам
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimSpace(parts[1])
		userID, err := jwt.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

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
