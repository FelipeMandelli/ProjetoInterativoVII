package main

import (
	"encoding/json"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/mjibson/go-dsp/fft"
	"gorm.io/gorm"

	service "pi.go/front-attempt/services"
	"pi.go/pkg/domain"
)

// DB é uma variável global para o banco de dados
var DB *gorm.DB

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ProcessedData representa os dados processados para o frontend
type ProcessedData struct {
	MotorID       string    `json:"motor_id"`
	Temperature   float32   `json:"temperature"`
	Sound         float32   `json:"sound"`
	Current       float32   `json:"current"`
	Vibration     []float64 `json:"vibration"`
	VibrationFreq []float64 `json:"vibration_freq"`
	DateTime      time.Time `json:"datetime"`
}

func main() {
	db, err := service.ConnectDatabase()
	if err != nil {
		log.Default().Fatalf("error connecting to database: %v", err)
	}

	DB = db

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./front-attempt/index.html")
	})

	r.Get("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "./front-attempt/styles.css")
	})

	r.Get("/data", getData)

	log.Println("Server started at :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Default().Fatalf("error starting server: %v", err)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	var dataCollections []domain.DataCollection
	DB.Find(&dataCollections)

	var processedData []ProcessedData

	for _, item := range dataCollections {
		vibrationStrings := strings.Split(item.Vibration, ",")
		vibrationFloats := make([]float64, len(vibrationStrings))
		for i, v := range vibrationStrings {
			vibrationFloats[i], _ = strconv.ParseFloat(v, 64)
		}

		// Apply FFT
		vibrationFreq := fft.FFTReal(vibrationFloats)
		vibrationMagnitudes := make([]float64, len(vibrationFreq))
		for i, c := range vibrationFreq {
			vibrationMagnitudes[i] = cmplx.Abs(c)
		}

		processedData = append(processedData, ProcessedData{
			MotorID:       item.MotorID,
			Temperature:   item.Temperature,
			Sound:         item.Sound,
			Current:       item.Current,
			Vibration:     vibrationFloats,
			VibrationFreq: vibrationMagnitudes,
			DateTime:      item.DateTime,
		})
	}

	response, err := json.Marshal(processedData)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
