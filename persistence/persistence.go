package main

import (
	"fmt"
	"sync"

	"pi.go/persistence/internal/service"
)

func main() {
	provider := service.NewProvider()

	err := service.ConnectDatabase(provider)
	if err != nil {
		provider.LogFatalf(fmt.Sprintf("error connecting to database: %v", err))
	}

	sub := service.NewZMQSubscriber()

	var wg sync.WaitGroup

	wg.Add(1)

	err = sub.Subscribe(provider)
	if err != nil {
		provider.LogFatalf(fmt.Sprintf("error subscribing: %v", err))
		wg.Done()
	}

	provider.Logf("waiting subscription event...")

	wg.Wait()
}
