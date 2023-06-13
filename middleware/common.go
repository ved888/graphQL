package middleware

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"grapgQL/dbhelper"
	"net/http"
)

type ContextKeys string

const (
	userContext ContextKeys = "__userContext"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding,Authorization, X-api-key")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// If it's an OPTIONS request, return immediately with a 200 status code
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// add recovery middleware
		RecoveryMiddleware(next)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic error
				logrus.Errorf("Recovered from panic: %v", err)

				// Set a 500 status code and write a message to the response
				w.WriteHeader(http.StatusInternalServerError)
				_, err = fmt.Fprint(w, "Internal Server Error")
				if err != nil {
					logrus.Errorf("error occured %v", err)
					return
				}
			}
		}()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request path is not the login route
		apiKey := r.Header.Get("x-api-key")
		if apiKey == "" {
			logrus.Errorf("empty token: %s", apiKey)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// verify session with token
		user, err := dbhelper.VerifySession(&apiKey)
		if err != nil || user == nil {
			logrus.WithError(err).Errorf("failed to get user with token: %s", apiKey)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userContext, user)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
