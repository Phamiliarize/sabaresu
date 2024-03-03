package lruntime

import (
	"context"
	"net/http"
)

type contextKey string

const sabaresuContextKey contextKey = "sabaresu"

type ContextKey string

func SetSabaresu(r *http.Request, value interface{}) {
	*r = *r.WithContext(context.WithValue(r.Context(), sabaresuContextKey, value))
}

func GetSabaresu(r *http.Request) interface{} {
	return r.Context().Value(sabaresuContextKey)
}
