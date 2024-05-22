package service

import (
	"context"
	"fmt"
	"log"

	"github.com/go-zeromq/zmq4"
	"pi.go/pkg/domain"
)

type Subscriber interface {
	Subscribe() error
}

type ZMQ struct {
	SubChan chan (domain.MeasurementDTO)
}

func NewZMQSubscriber() Subscriber {
	zmq := ZMQ{
		SubChan: make(chan domain.MeasurementDTO),
	}

	return &zmq
}

func (z *ZMQ) Subscribe() error {

	go func() {
		sub := zmq4.NewSub(context.Background())
		defer sub.Close()

		err := sub.Dial("tcp://0.0.0.0:5563")
		if err != nil {
			log.Fatalf("could not dial: %v", err)
		}

		err = sub.SetOption(zmq4.OptionSubscribe, "API-PUB")
		if err != nil {
			log.Fatalf("could not subscribe: %v", err)
		}

		for {
			msg, err := sub.Recv()
			if err != nil {
				log.Default().Printf("could not receive message: %v", err)

				continue
			} else {
				event := string(msg.Frames[0])

				if event == "API-PUB" {
					dto := domain.MeasurementDTO{}

					err := dto.FromByte(msg.Frames[1])
					if err != nil {
						log.Default().Printf("could not parse received message: %v", err)
						continue
					}

					z.SubChan <- dto
					continue
				}

				log.Default().Printf("received unknown event, ignoring...")
			}
		}
	}()

	// teste para remover o dto do chan
	go func() {
		for a := range z.SubChan {
			fmt.Printf("received dto: %+v\n", a)
		}
	}()

	return nil
}
