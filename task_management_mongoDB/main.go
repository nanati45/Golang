package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"task_management_mongoDB/data"
	"task_management_mongoDB/router"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client instance
var mongoClient *mongo.Client


func main() {
	
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
	
	mongoURI       := os.Getenv("MONGO_URI")
	dbName         := os.Getenv("MONGO_DB_NAME")

	fmt.Println(mongoURI , dbName)
	if mongoURI == "" || dbName == "" {
        log.Fatal("Environment variables MONGO_URI and MONGO_DB_NAME are required")
    }
	
	// MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create a context with timeout for MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection by pinging MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB")

	// Set the global MongoDB client
	data.InitMongoDB(client, dbName)

	// Defer disconnecting from MongoDB
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB")
	}()

	// Initialize the router
	r := router.SetupRouter(mongoClient)

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
