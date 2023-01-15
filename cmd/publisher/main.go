package main

import (
	"github.com/Owner-maker/nats-learning/internal/configs"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		logrus.Fatalf("error while initializing config file: %s", err.Error())
	}
	logrus.Print("successfully initialized config file")

	// connect to the nats streaming server
	sc, err := stan.Connect(
		config.ClusterId,
		config.ClientProducer,
		stan.NatsURL(config.NatsUrl))
	if err != nil {
		logrus.Fatalf("error while connnecting to the nats streaming server: %s", err.Error())
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			logrus.Fatalf("error while closing publisher connection to the nats streaming server: %s", err.Error())
		}
	}(sc)
	logrus.Print("successfully connected to the nats streaming server")

	// parse static json file
	dataJson, err := os.Open(config.JsonStaticModelPath)
	if err != nil {
		logrus.Fatalf("error while opening json file: %s", err.Error())
	}
	defer func(dataJson *os.File) {
		err = dataJson.Close()
		if err != nil {
			logrus.Fatalf("error while closing json fie: %s", err.Error())
		}
	}(dataJson)
	byteValue, _ := io.ReadAll(dataJson)

	// send json static repository to the nats streaming server
	err = sc.Publish(config.NatsSubject, byteValue)
	if err != nil {
		logrus.Fatalf("error while publishing json static file to the nats streaming server: %s", err.Error())
	}
	logrus.Print("successfully published json static file to the nats streaming server")
}
