package domain

import (
	"encoding/json"
	"time"
)

type MeasurementDTO struct {
	MotorID     string    `json:"motorId"`
	Temperature float32   `json:"temperatura"`
	Sound       float32   `json:"som"`
	Current     float32   `json:"corrente"`
	Vibration   []float32 `json:"vibracao"`
	DateTime    time.Time `json:"datetime"`
}

func (m *MeasurementDTO) ToByte() ([]byte, error) {
	return json.Marshal(m)
}

func (m *MeasurementDTO) FromByte(bytes []byte) error {
	return json.Unmarshal(bytes, m)
}
