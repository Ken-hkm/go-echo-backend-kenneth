package db

import (
	"context"
	"fmt"
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/secrets"
	"go.mongodb.org/mongo-driver/bson"
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
	return MongoClient.Database("Resume").Collection(collectionName)
}

// GetDocuments returns a MongoDB Document
func GetDocuments(collectionName string, filter bson.M) ([]bson.M, error) {
	collection := MongoClient.Database("Resume").Collection(collectionName)

	// Query the collection
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err // Return error properly
	}
	defer cursor.Close(context.TODO()) // Ensure cursor is closed

	var result []bson.M
	if err := cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, nil // Return the documents
}
