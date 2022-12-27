package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection is the function that connects to the database
var MongoConn = MongoConnection()

// ClientOptions is the options for the connection
var clientOptions = options.Client().ApplyURI("mongodb+srv://diegodapo77:<Dapo0917>@twitt.fcdyfbk.mongodb.net/?retryWrites=true&w=majority")

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
