package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
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

	response, err := json.Marshal(dataCollections)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
