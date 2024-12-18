package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://NHATANH:WMGaAVaGCtsPnC1k@cluster0.6z5yhqo.mongodb.net/devops-assignment")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	DB = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("devops-assignment").Collection(collectionName)
}
