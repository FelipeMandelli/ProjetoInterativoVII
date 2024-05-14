package main

import (
	"fmt"
	"net/http"

	"pi.go/internal/service"
)

func main() {
	provider := service.NewProvider()

	httpServer := &http.Server{
		Addr:    "0.0.0.0:9015",
		Handler: service.CreateRouter(provider),
	}

	provider.Logf(fmt.Sprintf("Starting API on: %s!", httpServer.Addr))
	if err := httpServer.ListenAndServe(); err != nil {
		provider.LogFatalf(fmt.Sprintf("error serving api: %s", err.Error()))
	}
}
