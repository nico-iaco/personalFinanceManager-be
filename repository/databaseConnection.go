package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"personalFinanceManager/utils"
)

func createConnection() *mongo.Client {
	databaseUrl := utils.GetStringProperty("database.url")
	username := utils.GetStringProperty("database.username")
	pwd := utils.GetStringProperty("database.password")
	credential := options.Credential{
		Username: username,
		Password: pwd,
	}
	clientOptions := options.Client().ApplyURI(databaseUrl).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

func disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
