package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"personalFinanceManager/src/utils"
)

var Client *mongo.Client

func CreateConnection() {
	databaseUrl := utils.GetStringProperty("database.url")
	username := utils.GetStringProperty("database.username")
	pwd := utils.GetStringProperty("database.password")
	credential := options.Credential{
		Username: username,
		Password: pwd,
	}
	var err error
	clientOptions := options.Client().ApplyURI(databaseUrl).SetAuth(credential)
	Client, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Started mongodb connection...")
}

func Disconnect() {
	err := Client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
