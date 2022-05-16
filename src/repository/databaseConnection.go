package repository

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"personalFinanceManager/src/utils"
)

func CreateConnection() {
	databaseUrl := utils.GetStringProperty("database.url")
	username := utils.GetStringProperty("database.username")
	pwd := utils.GetStringProperty("database.password")
	credential := options.Credential{
		Username: username,
		Password: pwd,
	}
	clientOptions := options.Client().ApplyURI(databaseUrl).SetAuth(credential)
	//Client, err = mongo.Connect(context.Background(), clientOptions)
	err := mgm.SetDefaultConfig(nil, "personal-finance", clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started mongodb connection...")
}
