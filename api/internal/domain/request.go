package domain

type Request struct {
	Temperature float32   `json:"temperatura"`
	Sound       float32   `json:"som"`
	Current     float32   `json:"corrente"`
	Vibration   []float32 `json:"vibracao"`
	MotorID     string    `json:"id"`
}
