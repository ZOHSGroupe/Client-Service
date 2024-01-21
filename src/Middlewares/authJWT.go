package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func VerifyToken() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized!", http.StatusUnauthorized)
				return
			}
		})
	}
}
