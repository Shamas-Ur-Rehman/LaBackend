package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// MongoDB connection URI
	// uri := "mongodb+srv://shamasurrehman398:enEr1ytql29Q6N1d@cluster0.frjl1bw.mongodb.net/"
	uri := "mongodb+srv://shamasurrehman398:enEr1ytql29Q6N1d@cluster0.frjl1bw.mongodb.net/?retryWrites=true&w=majority&tls=false"

	// Create a new Mongo client
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Assign database (change "productdb" to your DB name)
	DB = client.Database("productdb")

	log.Println("✅ MongoDB connected successfully.")
}

// GetCollection returns a MongoDB collection reference
func GetCollection(name string) *mongo.Collection {
	if DB == nil {
		log.Fatal("❌ MongoDB not initialized. Call ConnectDB() first.")
	}
	return DB.Collection(name)
}
