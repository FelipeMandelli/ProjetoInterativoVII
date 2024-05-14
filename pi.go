package main

import (
	"fmt"
	"net/http"

	"pi.go/internal/service"
)

func main() {

	fmt.Println("Hello World!")

	provider := service.NewProvider()

	httpServer := &http.Server{
		Addr:    "0.0.0.0:9015",
		Handler: service.CreateRouter(provider),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		provider.Log.Fatalf("error serving api: %s", err.Error())
	}
}
