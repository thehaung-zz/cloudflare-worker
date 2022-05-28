package config

import (
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	ServerHost    string
	MongoUser     string
	MongoPassword string
	MongoHost     string
	MongoPort     string
}

func (config *Config) initialize() {
	config.ServerHost = os.Getenv("MONGO_SERVER_HOST")
	config.MongoPort = os.Getenv("MONGO_PORT")
	config.MongoUser = os.Getenv("MONGO_USER")
	config.MongoPassword = os.Getenv("MONGO_PASSWORD")
	config.MongoHost = os.Getenv("MONGO_HOST")
}

func (config *Config) MongoURI() string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s",
		config.MongoUser,
		url.QueryEscape(config.MongoPassword),
		config.MongoHost,
	)
}

func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
