package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initializing config file: %s", err.Error())
	}

	// connect to nats streaming server
	sc, err := stan.Connect(viper.GetString("nats.clusterID"), viper.GetString("nats.clientProducer"), stan.NatsURL(viper.GetString("nats.serverURL")))
	if err != nil {
		logrus.Fatalf("error while connnecting to nats streaming server: %s", err.Error())
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			logrus.Fatalf("error while closing publisher connection to nats streaming server: %s", err.Error())
		}
	}(sc)

	fmt.Println(sc)

}
