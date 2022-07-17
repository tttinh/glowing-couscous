package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	VNPhoneCode = "84"
)

type Config struct {
	Env                 string
	ServicePort         string
	MongoURL            string
	MongoDatabase       string
	MongoReplSet        string
	SentryDSN           string
	SecretAuthKey       string
	AuthServiceEndpoint string
}

// NewConfig :
func NewConfig() *Config {
	godotenv.Load()
	conf := &Config{
		Env:                 os.Getenv("ENV"),
		ServicePort:         os.Getenv("SERVICE_PORT"),
		MongoURL:            os.Getenv("MONGO_URL"),
		MongoDatabase:       os.Getenv("MONGO_DATABASE"),
		MongoReplSet:        os.Getenv("MONGO_REPLICA_SET"),
		AuthServiceEndpoint: os.Getenv("AUTH_SERVICE_ENDPOINT"),
		SentryDSN:           os.Getenv("SENTRY_DSN"),
		SecretAuthKey:       os.Getenv("SECRET_AUTH_KEY"),
	}
	if conf.Env == "dev" {
		c, _ := json.MarshalIndent(conf, "", "\t")
		fmt.Println("Config:", string(c))
	}

	return conf
}
