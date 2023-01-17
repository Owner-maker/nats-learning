package main

import (
	"context"
	"github.com/Owner-maker/nats-learning/internal/configs"
	"github.com/Owner-maker/nats-learning/internal/delivery/http"
	"github.com/Owner-maker/nats-learning/internal/delivery/nats"
	"github.com/Owner-maker/nats-learning/internal/repository/cache"
	"github.com/Owner-maker/nats-learning/internal/repository/postgres"
	"github.com/Owner-maker/nats-learning/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// @title Nats learning service
// @version 1.0
// @description This service uses a nats streaming server as message broker to get model Order from it and stores into the postgres db & app's cache. Provides a way to get information about orders from cache via the HTTP requests.

// @host localhost:8080
// @basePath /

// @contact.name Artem Lisitsyn
// @contact.email artem.lisitsynn@gmail.com
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
		config.NatsUrlSub)
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

	// init handler
	httpHandler := http.NewHandler(s)
	//init server
	srv := new(http.Server)
	go func() {
		if err = srv.Run(config.AppPort, httpHandler.InitRoutes()); err != nil {
			logrus.Fatal(err)
		}
	}()
	logrus.Print("Service is successfully started...")

	// graceful shutdown

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	wg.Wait()
}
