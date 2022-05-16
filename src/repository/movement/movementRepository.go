package movement

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"personalFinanceManager/src/model"
	"personalFinanceManager/src/utils"
)

func getMovementsCollection() *mgm.Collection {
	movementsCollection := mgm.Coll(&model.Movement{})
	return movementsCollection
}

func AddMovement(movement model.Movement) model.Movement {
	err := getMovementsCollection().Create(&movement)
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
		{"id", utils.SanitizeString(movementId)},
		{"user", utils.SanitizeString(userId)},
	}
	element := getMovement(filter)
	err := movementsCollection.Delete(&element)
	if err != nil {
		log.Println(err)
		log.Printf("The movement %v was not deleted", utils.SanitizeString(movementId))
		return false
	}
	return true
}

func getMovement(filter bson.D) model.Movement {
	result := model.Movement{}
	err := getMovementsCollection().First(filter, &result)
	if err != nil {
		log.Println(err)
		return model.Movement{}
	}
	return result
}

func getMovementList(filter bson.D) []*model.Movement {
	movementsCollection := getMovementsCollection()
	var userMovements []*model.Movement
	err := movementsCollection.SimpleFind(&userMovements, filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found multiple documents (array of pointers): %+v\n", userMovements)
	return userMovements
}
