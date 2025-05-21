package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, *mongo.Collection, error) {
	err := godotenv.Load(); 
	if err != nil {
		log.Println("Warning: No .env file found, falling back to environment variables")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return nil, nil, fmt.Errorf("MONGODB_URI not set in environment")
	}

	dbName := os.Getenv("MONGODB_DB")
	if dbName == "" {
		dbName = "todo_db"
	}

	collectionName := os.Getenv("MONGODB_COLLECTION")
	if collectionName == "" {
		collectionName = "todos"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err 
	}

	fmt.Println("Connected to mongoDB!")

	collection := client.Database(dbName).Collection(collectionName)

	return client, collection, nil 
}

