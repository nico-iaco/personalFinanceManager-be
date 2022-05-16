package user

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"personalFinanceManager/src/model"
)

func getUserCollection() *mgm.Collection {
	userCollection := mgm.Coll(&model.User{})
	return userCollection
}

func AddUser(user model.User) model.User {
	userCollection := getUserCollection()
	err := userCollection.Create(&user)
	if err != nil {
		log.Fatal(err)
		return model.User{}
	}
	log.Print("Inserted one user : ", user.ID)
	return user
}

func GetUser(email string) model.User {
	userCollection := getUserCollection()
	var result model.User
	filter := bson.D{{"email", email}}
	err := userCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return model.User{}
	}
	return result
}

func GetUserById(id string) model.User {
	userCollection := getUserCollection()
	var result model.User
	filter := bson.D{{"id", id}}
	err := userCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return model.User{}
	}
	return result
}

func UpdateUser(user model.User, updatedFields bson.D) model.User {
	userCollection := getUserCollection()
	var result model.User
	filter := bson.D{{"id", user.ID}}
	err := userCollection.FindOneAndUpdate(context.Background(), filter, updatedFields).Decode(&result)
	log.Printf("Updated user %v", result)
	if err != nil {
		return model.User{}
	}
	return result
}

func CheckEmailExists(email string) bool {
	userCollection := getUserCollection()
	filter := bson.D{{"email", email}}
	err := userCollection.FindOne(context.Background(), filter).Err()
	return err == nil
}
