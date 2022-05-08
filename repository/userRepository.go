package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"personalFinanceManager/model"
)

func getUserCollection() (*mongo.Collection, *mongo.Client) {
	client := createConnection()
	userCollection := client.Database("personal-finance").Collection("users")
	return userCollection, client
}

func AddUser(user model.User) model.User {
	userCollection, client := getUserCollection()
	insertOneResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return model.User{}
	}
	log.Print("Inserted one user : ", insertOneResult.InsertedID)
	disconnect(client)
	return user
}

func GetUser(email string) model.User {
	userCollection, client := getUserCollection()
	var result model.User
	filter := bson.D{{"email", email}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return model.User{}
	}
	disconnect(client)
	return result
}

func CheckEmailExists(email string) bool {
	userCollection, client := getUserCollection()
	filter := bson.D{{"email", email}}
	err := userCollection.FindOne(context.TODO(), filter).Err()
	disconnect(client)
	return err == nil
}
