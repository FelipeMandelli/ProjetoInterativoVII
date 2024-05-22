package main

import (
	"fmt"
	"log"
	"net/http"

	"pi.go/api/internal/service"
)

func main() {
	publisher, err := service.NewZMQPublisher()
	if err != nil {
		log.Fatalf("error initializing publisher: %v", err)
	}

	provider := service.NewProvider(publisher)

	httpServer := &http.Server{
		Addr:    "0.0.0.0:9015",
		Handler: service.CreateRouter(provider),
	}

	provider.Logf(fmt.Sprintf("Starting API on: %s!", httpServer.Addr))
	if err := httpServer.ListenAndServe(); err != nil {
		provider.LogFatalf(fmt.Sprintf("error serving api: %s", err.Error()))
	}
}
