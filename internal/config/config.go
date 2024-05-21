package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/sethvargo/go-envconfig"
)

type ENV struct {
	AppName     string `env:"APP_NAME"`
	Environment string `env:"ENVIRONMENT"`
	Debug       bool   `env:"DEBUG"`

	// MongoDB
	MongoDB MongoDBCfg

	// Redis
	RedisURL string `env:"REDIS_URL"`

	// Sentry
	SentryDSN     string `env:"SENTRY_DSN"`
	SentryMachine string `env:"SENTRY_MACHINE"`

	// Queue
	QueueUsername    string `env:"QUEUE_USERNAME"`
	QueuePassword    string `env:"QUEUE_PASSWORD"`
	QueueConcurrency int    `env:"QUEUE_CONCURRENCY"`
}

var env ENV

// GetENV ...
func GetENV() ENV {
	return env
}

// MongoDBCfg ...
type MongoDBCfg struct {
	URI                 string `env:"MONGO_URL"`
	DBName              string `env:"MONGO_DB_NAME"`
	ReplicaSet          string `env:"MONGO_REPLICA_SET"`
	CAPem               string `env:"MONGO_CA_PEM"`
	CertPem             string `env:"MONGO_CERT_PEM"`
	CertKeyFilePassword string `env:"MONGO_CERT_KEY_FILE_PASSWORD"`
	ReadPrefMode        string `env:"MONGO_READ_PREF_MODE"`
}

// InitWithLoadENV read env config
func InitWithLoadENV() {
	var ctx = context.Background()
	envFile := ".env"
	if err := godotenv.Load(envFile); err != nil {
		fmt.Println("Load env file err: ", err)
	}
	if err := envconfig.Process(ctx, &env); err != nil {
		log.Fatal("Assign env err: ", err)
	}
}

// Init ...
func Init() {
	if getEnvStr("ENVIRONMENT") == "" {
		panic(errors.New("ENVIRONMENT is not set"))
	}

	// from .env
	env = ENV{
		AppName:     getEnvStr("APP_NAME"),
		Environment: getEnvStr("ENVIRONMENT"),
		Debug:       getEnvBool("DEBUG"),

		RedisURL: getEnvStr("REDIS_URL"),

		SentryDSN:     getEnvStr("SENTRY_DSN"),
		SentryMachine: getEnvStr("SENTRY_MACHINE"),

		QueueUsername:    getEnvStr("QUEUE_USERNAME"),
		QueuePassword:    getEnvStr("QUEUE_PASSWORD"),
		QueueConcurrency: getEnvInt("QUEUE_CONCURRENCY"),
		MongoDB: MongoDBCfg{
			URI:                 getEnvStr("MONGO_URL"),
			DBName:              getEnvStr("MONGO_DB_NAME"),
			ReplicaSet:          getEnvStr("MONGO_REPLICA_SET"),
			CAPem:               getEnvStr("MONGO_CA_PEM"),
			CertPem:             getEnvStr("MONGO_CERT_PEM"),
			CertKeyFilePassword: getEnvStr("MONGO_CERT_KEY_FILE_PASSWORD"),
			ReadPrefMode:        getEnvStr("MONGO_READ_PREF_MODE"),
		},
	}

	fmt.Printf("⚡️ [env]: %s \n", env.Environment)
	fmt.Printf("⚡️ [config]: loaded \n")
}
