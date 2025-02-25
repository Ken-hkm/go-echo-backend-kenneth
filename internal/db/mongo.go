package db

import (
	"context"
	"fmt"
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/secrets"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectMongoDB initializes a connection to MongoDB.
func ConnectMongoDB() {
	mongoURI, err := secrets.GetParameter("mongo_uri")
	if err != nil {
		log.Printf("Failed to get parameter: %v", err)
	}
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	MongoClient = client
}

// GetCollection returns a MongoDB collection handle.
func GetCollection(collectionName string) *mongo.Collection {
	// Replace "online_cv" with your database name
	return MongoClient.Database("Resume").Collection(collectionName)
}
