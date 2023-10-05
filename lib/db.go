package lib

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opt := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opt)

	if err != nil {
		log.Fatalf("Error connecting MongoDB: %v", err)
	}
	return client
}
