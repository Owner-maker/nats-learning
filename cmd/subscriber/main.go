package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"nats-learning/internal/transport"
	"nats-learning/internal/util"
	"sync"
)

func main() {
	var validate = validator.New()
	var wg sync.WaitGroup

	// parse configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		logrus.Fatalf("error while initializing config file: %s", err.Error())
	}
	logrus.Print("successfully initialized config file")

	// connect to the nats streaming server
	sc, err := transport.Connect(
		config.ClusterId,
		config.ClientSubscriber,
		config.NatsUrl)
	if err != nil {
		return
	}
	defer func(sc stan.Conn) {
		err = sc.Close()
		if err != nil {
			logrus.Fatalf("error while closing subscriber connection to the nats streaming server: %s", err.Error())
		}
	}(sc)

	//subscribe to the nats subject "orders"
	wg.Add(1)
	go func() {
		err = transport.Subscribe(&wg, validate, sc, config.NatsSubject)
		if err != nil {
			return
		}
	}()

	wg.Wait()
}
