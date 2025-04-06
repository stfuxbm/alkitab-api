package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(" Error loading .env file")
	}
}

// GetMongoURI returns the MongoDB connection URI from env
func GetMongoURI() string {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal(" MONGO_URI not set in environment")
	}
	return uri
}

// GetMongoDBName returns the MongoDB database name from env
func GetMongoDBName() string {
	db := os.Getenv("MONGO_DB")
	if db == "" {
		log.Fatal("MONGO_DB not set in environment")
	}
	return db
}
