package lruntime

import (
	"log"
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

// RegisterRuntimeMiddleware applies all runtime middleware to the handler function
func RegisterRuntimeMiddleware(middlewares []Middleware, handler http.HandlerFunc) http.HandlerFunc {
	RuntimeHandler := handler

	for _, middleware := range middlewares {
		RuntimeHandler = middleware(RuntimeHandler)
	}

	return RuntimeHandler
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

// SabaresuContext adds the context object Sabaresu will expose to Lua functions
func SabaresuContext(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SetSabaresu(r, "hi")
		next.ServeHTTP(w, r)
	})
}
