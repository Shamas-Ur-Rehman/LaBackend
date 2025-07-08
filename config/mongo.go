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
	uri := "mongodb+srv://shamasurrehman398:enEr1ytql29Q6N1d@cluster0.frjl1bw.mongodb.net/?retryWrites=true&w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	// Ping to test connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ MongoDB ping failed: %v", err)
	}

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
