package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)	
	if err != nil {
		log.Fatal(err)
	}
	DB := client.Database("go_news")
	return DB
}