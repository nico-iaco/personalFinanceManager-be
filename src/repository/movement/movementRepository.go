package movement

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"personalFinanceManager/src/model"
	"personalFinanceManager/src/repository"
	"strings"
)

func getMovementsCollection() *mongo.Collection {
	movementsCollection := repository.Client.Database("personal-finance").Collection("movements")
	return movementsCollection
}

func AddMovement(movement model.Movement) model.Movement {
	movementsCollection := getMovementsCollection()
	_, err := movementsCollection.InsertOne(context.Background(), movement)
	if err != nil {
		log.Fatal(err)
		return model.Movement{}
	}
	return movement
}

func GetUserMovements(userId string) []*model.Movement {
	filter := bson.D{
		{"user", userId},
	}
	return getMovementList(filter)
}

func GetUserAccountMovements(userId, accountName string) []*model.Movement {
	filter := bson.D{
		{"user", userId},
		{"$or", []bson.D{
			{{"source", accountName}},
			{{"destination", accountName}},
		}},
	}
	return getMovementList(filter)
}

func EditUserMovement(userId, movementId string, updatedField bson.D) model.Movement {
	movementsCollection := getMovementsCollection()
	var result model.Movement
	filter := bson.D{
		{"id", movementId},
		{"user", userId},
	}
	err := movementsCollection.FindOneAndUpdate(context.Background(), filter, updatedField).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return model.Movement{}
	}
	return result
}

func DeleteUserMovement(userId, movementId string) bool {
	movementsCollection := getMovementsCollection()

	filter := bson.D{
		{"id", movementId},
		{"user", userId},
	}
	element, err := movementsCollection.DeleteOne(context.Background(), filter)
	if err != nil || element.DeletedCount == 0 {
		log.Println(err)
		escapedMovementId := strings.Replace(userId, "\n", "", -1)
		escapedMovementId = strings.Replace(escapedMovementId, "\r", "", -1)
		log.Printf("The movement %v was not deleted", escapedMovementId)
		return false
	}
	return true
}

func getMovementList(filter bson.D) []*model.Movement {
	movementsCollection := getMovementsCollection()
	var userMovements []*model.Movement
	cur, err := movementsCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem model.Movement
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		userMovements = append(userMovements, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	err = cur.Close(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Printf("Found multiple documents (array of pointers): %+v\n", userMovements)
	return userMovements
}
