package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURL = string(os.Getenv("MONGO_HOST"))
	Client   *mongo.Client
)

func ConnectMongo() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:passwordunchanged@cluster0.jc8dwgb.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting:", err)
	}

	log.Println("Connected to mongo db server!")

	return client
}
