package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	provider Provider
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal("API is Healthy!")
	if err != nil {
		h.provider.Logf(fmt.Sprintf("error marshalling response: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	h.provider.Logf(fmt.Sprintf("Health Checked by IP %s", r.RemoteAddr))
	w.Write(resp)
}
