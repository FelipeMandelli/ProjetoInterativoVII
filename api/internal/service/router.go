package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	healthPath        = "/health"
	dataPath          = "/data"
	getterPath        = "/data"
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

func CreateRouter(provider *Provider) http.Handler {
	handler := Handler{
		provider: provider,
	}

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(contentTypeHeader, jsonContentType)
			next.ServeHTTP(w, r)
		})
	})

	r.Get(healthPath, handler.HealthCheckHandler)
	r.Put(dataPath, handler.DataReceiverHandler)
	r.Get(getterPath, handler.DataGetterHandler)

	return r
}
