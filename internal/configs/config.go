package configs

import "github.com/spf13/viper"

type Config struct {
	PostgresHost        string `mapstructure:"POSTGRES_HOST"`
	PcPostgresPort      string `mapstructure:"PC_POSTGRES_PORT"`
	ContPostgresPort    string `mapstructure:"CONT_POSTGRES_PORT"`
	PostgresUser        string `mapstructure:"POSTGRES_USER"`
	PostgresDb          string `mapstructure:"POSTGRES_DB"`
	PostgresPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresSslMode     string `mapstructure:"POSTGRES_SSL_MODE"`
	PcNatsPort1         string `mapstructure:"PC_NATS_PORT1"`
	PontNatsPort1       string `mapstructure:"CONT_NATS_PORT1"`
	PcNatsPort2         string `mapstructure:"PC_NATS_PORT2"`
	ContNatsPort2       string `mapstructure:"CONT_NATS_PORT2"`
	NatsUrl             string `mapstructure:"NATS_URL"`
	ClusterId           string `mapstructure:"CLUSTER_ID"`
	ClientProducer      string `mapstructure:"CLIENT_PRODUCER"`
	ClientSubscriber    string `mapstructure:"CLIENT_SUBSCRIBER"`
	JsonStaticModelPath string `mapstructure:"JSON_STATIC_MODEL_PATH"`
	NatsSubject         string `mapstructure:"NATS_SUBJECT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
