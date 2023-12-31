package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// Middleware function for token verification
func VerifyToken() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip token verification for POST requests
			if r.Method == "POST" {
				next.ServeHTTP(w, r)
				return
			}

			tokenHeader := r.Header.Get("Authorization")

			if tokenHeader == "" {
				http.Error(w, "No token provided!", http.StatusForbidden)
				return
			}

			token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET_JWT")), nil
			})

			if err != nil {
				http.Error(w, "Unauthorized!", http.StatusUnauthorized)
				return
			}

			// If token is valid, you can access the claims
			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// You can access the claims using claims["key"]
				// For example, req.UserId = claims["id"].(string)
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized!", http.StatusUnauthorized)
				return
			}
		})
	}
}
