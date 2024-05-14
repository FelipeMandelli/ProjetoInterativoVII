package service

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	provider Provider
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal("API is Healthy!")
	if err != nil {
		h.provider.Log.Printf("error marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	h.provider.Log.Printf("Health Checked by IP %s", r.RemoteAddr)
	w.Write(resp)
}
