package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	host     string
	username string
	password string
	dbname   string
)

func getConnectionString() string {
	host = os.Getenv("DB_HOST")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", username, password, host)
	return mongoURI
}

func ConfigDB() *mongo.Database {
	env := os.Getenv("APP_ENV")
	connectionString := getConnectionString()
	if env == "production" {
		connectionString = getConnectionString()
	} else {
		connectionString = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	dbname = os.Getenv("EXERCISE_DB_NAME")
	return client.Database(dbname)
}
