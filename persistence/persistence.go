package main

import (
	"fmt"
	"sync"

	"pi.go/persistence/internal/service"
)

func main() {
	provider := service.NewProvider()

	sub := service.NewZMQSubscriber()

	var wg sync.WaitGroup

	wg.Add(1)

	err := sub.Subscribe()
	if err != nil {
		provider.LogFatalf(fmt.Sprintf("error subscribing: %v", err))
		wg.Done()
	}

	provider.Logf("waiting subscription event...")

	wg.Wait()
}
