package service

import (
	"context"
	"log"

	zmq "github.com/go-zeromq/zmq4"
	"pi.go/pkg/domain"
)

type Publisher interface {
	Publish(domain.MeasurementDTO) error
}

type ZMQ struct {
	pubChan chan (domain.MeasurementDTO)
}

func NewZMQPublisher() (Publisher, error) {
	// pub := zmq.NewPub(context.Background())
	// defer pub.Close()

	// err := pub.Listen("tcp://*:5563")
	// if err != nil {
	// 	return nil, fmt.Errorf("could not listen: %w", err)
	// }

	log.Default().Println("listener started...")

	zmq := ZMQ{
		pubChan: make(chan domain.MeasurementDTO),
	}

	go zmq.startPublisher()

	return &zmq, nil
}

func (z ZMQ) startPublisher() {
	pub := zmq.NewPub(context.Background())
	defer pub.Close()

	err := pub.Listen("tcp://*:5563")
	if err != nil {
		log.Default().Fatalf("could not listen: %v\n", err)
	}

	log.Default().Println("listener started...")

	for msg := range z.pubChan {
		log.Default().Println("sending message to publisher...")

		msgBytes, err := msg.ToByte()
		if err != nil {
			log.Default().Printf("could not parse message into bytes to be sent: %v\n", err.Error())
		}

		pub.Send(zmq.NewMsgFrom(
			[]byte("API-PUB"),
			msgBytes,
		))
	}
}

func (zmq *ZMQ) Publish(payload domain.MeasurementDTO) error {
	zmq.pubChan <- payload

	return nil
}
