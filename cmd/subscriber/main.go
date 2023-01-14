package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"nats-learning/internal/configs"
	"nats-learning/internal/delivery/http"
	"nats-learning/internal/delivery/nats"
	"nats-learning/internal/repository/cache"
	"nats-learning/internal/repository/postgres"
	"nats-learning/internal/service"
	"sync"
)

func main() {
	// parse configuration
	config, err := configs.LoadConfig(".")
	if err != nil {
		logrus.Fatalf("error while parsing config file: %s", err.Error())
	}
	logrus.Print("successfully parsed config file")

	//init cache
	orderCache := cache.NewOrderCache(cache.NewCache())

	logrus.Print("successfully initialized cache")

	//connect to the postgres
	db, err := postgres.ConnectDB(
		postgres.Config{
			Host:     config.PostgresHost,
			Port:     config.PostgresPort,
			Username: config.PostgresUser,
			Password: config.PostgresPassword,
			DbName:   config.PostgresDb,
			SslMode:  config.PostgresSslMode,
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	orderPostgres := postgres.NewOrderPostgres(db)

	//create service
	s := service.NewService(*orderCache, *orderPostgres)
	//fill the cache from postgres
	err = s.PutOrdersFromDbToCache()
	if err != nil {
		logrus.Fatal(err)
	}

	// connect to the nats streaming server
	natsStreaming := nats.NewNats(s, validator.New())

	sc, err := natsStreaming.Connect(
		config.ClusterId,
		config.ClientSubscriber,
		config.NatsUrl)
	if err != nil {
		return
	}
	defer func(sc stan.Conn) {
		err = sc.Close()
		if err != nil {
			logrus.Errorf("error while closing subscriber connection to the nats streaming server: %s", err.Error())
		}
	}(sc)

	//subscribe to the nats subject "orders"
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err = natsStreaming.Subscribe(&wg, sc, config.NatsSubject)
		if err != nil {
			return
		}
	}()
	logrus.Print("successfully subscribed to the nats streaming subject orders")
	logrus.Print("Service is successfully started...")

	// init handler
	httpHandler := http.NewHandler(s)
	err = httpHandler.InitRoutes().Run()
	if err != nil {
		logrus.Fatal(err)
	}

	wg.Wait()
}
