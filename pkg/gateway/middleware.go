package gateway

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const RequestIDContextKey contextKey = "requestID"

type Middleware func(next http.HandlerFunc) http.HandlerFunc

// RegisterRuntimeMiddleware applies all runtime middleware to the handler function
func RegisterRuntimeMiddleware(middlewares []Middleware, handler http.HandlerFunc) http.HandlerFunc {
	RuntimeHandler := handler

	// We need to reverse back the "human" order because of how wrapping functions works
	for i := len(middlewares) - 1; i >= 0; i-- {
		RuntimeHandler = middlewares[i](RuntimeHandler)
	}

	return RuntimeHandler
}

// RequestLogging logs basic information on all requests and adds a canonical request ID
func RequestLogging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.NewString()
		r = r.WithContext(context.WithValue(r.Context(), RequestIDContextKey, requestId))
		log.Printf("[%s] %s Request ID: %s", r.Method, r.RequestURI, requestId)
		next.ServeHTTP(w, r)
	})
}

// RequestValidator validates the incoming request against a schema
func RequestValidator(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do Stuff
		next.ServeHTTP(w, r)
	})
}

// ResponseSweeper builds response against the schema
func ResponseSweeper(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		// Do stuff
	})
}

// PanicRecovery handles any panic errors that may occur gracefully
func PanicRecovery(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[DEBUG] panic occurred: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
