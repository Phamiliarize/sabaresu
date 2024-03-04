package lruntime

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const requestIDContextKey contextKey = "requestID"

type Middleware func(next http.HandlerFunc) http.HandlerFunc

// RegisterRuntimeMiddleware applies all runtime middleware to the handler function
func RegisterRuntimeMiddleware(middlewares []Middleware, handler http.HandlerFunc) http.HandlerFunc {
	RuntimeHandler := handler

	for _, middleware := range middlewares {
		RuntimeHandler = middleware(RuntimeHandler)
	}

	return RuntimeHandler
}

// RequestLogging logs basic information on all requests and adds a canonical request ID
func RequestLogging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.NewString()
		r = r.WithContext(context.WithValue(r.Context(), requestIDContextKey, requestId))
		log.Printf("[%s] %s RequestID: %s", r.Method, r.RequestURI, requestId)
		next.ServeHTTP(w, r)
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
