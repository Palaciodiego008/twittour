package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection is the function that connects to the database
var MongoConn = MongoConnection()

// ClientOptions is the options for the connection
var clientOptions = options.Client().ApplyURI(applyURI())

// MongoConnection is the function that connects to the database
func MongoConnection() *mongo.Client {

	ctx := context.TODO()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return &mongo.Client{}
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
		return &mongo.Client{}
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func CheckConnection() int {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

func applyURI() string {
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")

	return "mongodb+srv://" + user + ":" + "<" + password + ">" + "@twitt.fcdyfbk.mongodb.net/?retryWrites=true&w=majority"
}
