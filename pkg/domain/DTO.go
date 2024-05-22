package domain

import "encoding/json"

type MeasurementDTO struct {
	Temperature float32   `json:"temperatura"`
	Sound       float32   `json:"som"`
	Current     float32   `json:"corrente"`
	Vibration   []float32 `json:"vibracao"`
}

func (m *MeasurementDTO) ToByte() ([]byte, error) {
	return json.Marshal(m)
}

func (m *MeasurementDTO) FromByte(bytes []byte) error {
	return json.Unmarshal(bytes, m)
}
