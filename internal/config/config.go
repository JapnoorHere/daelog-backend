package config

import (
	"log"
	"os"
)

type Config struct {
	Port     string
	MongoURI string
}

func Load() *Config {
	mongoURI := os.Getenv("DAELOG_MONGO_URI")
	if mongoURI == "" {
		log.Fatal("DAELOG_MONGO_URI is required")
	}

	port := os.Getenv("DAELOG_PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:     port,
		MongoURI: mongoURI,
	}
}
