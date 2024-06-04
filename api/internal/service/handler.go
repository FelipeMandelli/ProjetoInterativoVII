package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"pi.go/api/internal/domain"
	domainpkg "pi.go/pkg/domain"
)

type Handler struct {
	provider *Provider
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

func (h *Handler) DataReceiverHandler(w http.ResponseWriter, r *http.Request) {
	h.provider.Logf(fmt.Sprintf("received data request from %s, processing...", r.RemoteAddr))

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.provider.Logf(fmt.Sprintf("could not read request body: %s", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	req := new(domain.Request)

	err = json.Unmarshal(body, req)
	if err != nil {
		h.provider.Logf(fmt.Sprintf("could not unmarshall request body: %s", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = validatePayload(*req)
	if err != nil {
		h.provider.Logf(fmt.Sprintf("invalid payload: %s", err.Error()))

		resp, _ := json.Marshal("Invalid Payload")

		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)

		return
	}

	err = h.provider.Publisher.Publish(domainpkg.MeasurementDTO{
		MotorID:     req.MotorID,
		Temperature: req.Temperature,
		Sound:       req.Sound,
		Current:     req.Current,
		Vibration:   req.Vibration,
		DateTime:    time.Now(),
	})
	if err != nil {
		h.provider.Logf(fmt.Sprintln("could not send req to persistence"))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// DataReceiverHandler manipulador para receber dados com filtros opcionais
func (h *Handler) DataGetterHandler(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]string)

	if motorID := r.URL.Query().Get("motor_id"); motorID != "" {
		filters["motor_identification"] = motorID
	}

	dataCollections, err := GetDataCollections(h.provider, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataCollections)
}

func validatePayload(req domain.Request) error {
	if req.Current == 0 {
		return fmt.Errorf("")
	}
	return nil
}
