package configuration

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client holds the MongoDB client connection
var Client *mongo.Client

// Connect establishes a connection to the MongoDB database
func Connect() {
	var err error
	// Replace the URI with your MongoDB connection string
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to ensure connection is established
	if err = Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("MongoDB is unreachable: %v", err)
	}
	log.Println("Connected to MongoDB!")
}
