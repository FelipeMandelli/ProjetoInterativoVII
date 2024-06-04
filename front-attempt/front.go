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
		http.ServeFile(w, r, "index.html")
	})

	r.Get("/ws", handleWebSocket)

	http.ListenAndServe(":8080", r)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	for {
		_, motorID, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var dataCollections []domain.DataCollection
		DB.Where("motor_identification = ?", string(motorID)).Find(&dataCollections)

		response, err := json.Marshal(dataCollections)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Error: "+err.Error()))
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
			return
		}
	}
}
