package main

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"nats-learning/configs"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		logrus.Fatalf("error while initializing config file: %s", err.Error())
	}
	logrus.Print("successfully initialized config file")

	// connect to nats streaming server
	sc, err := stan.Connect(config.ClusterId, config.ClientProducer, stan.NatsURL(config.NatsUrl))
	if err != nil {
		logrus.Fatalf("error while connnecting to nats streaming server: %s", err.Error())
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			logrus.Fatalf("error while closing publisher connection to nats streaming server: %s", err.Error())
		}
	}(sc)
	logrus.Print("successfully connected to nats streaming server")

}
