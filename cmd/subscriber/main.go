package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"nats-learning/internal/configs"
	"nats-learning/internal/delivery/nats"
	"nats-learning/internal/repository/cache"
	"nats-learning/internal/repository/postgres"
	"sync"
)

func main() {
	var validate = validator.New()
	var wg sync.WaitGroup

	// parse configuration
	config, err := configs.LoadConfig(".")
	if err != nil {
		logrus.Fatalf("error while parsing config file: %s", err.Error())
	}
	logrus.Print("successfully parsed config file")

	// connect to the nats streaming server
	sc, err := nats.Connect(
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
		err = nats.Subscribe(&wg, validate, sc, config.NatsSubject)
		if err != nil {
			return
		}
	}()

	//connect to the postgres
	_, err = postgres.ConnectDB(
		postgres.Config{
			Host:     config.PostgresHost,
			Port:     config.PcPostgresPort,
			Username: config.PostgresUser,
			Password: config.PostgresPassword,
			DbName:   config.PostgresDb,
			SslMode:  config.PostgresSslMode,
		},
	)

	if err != nil {
		logrus.Fatal(err)
		return
	}

	//init cache
	_ = cache.NewOrderCache(cache.NewCache())
	logrus.Print("successfully initialized cache")

	wg.Wait()
}
