package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN is the database connection object
var MongoCN = ConnectDatabase()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user1:contraseña1@cluster-careerpath.1yinv.mongodb.net/test")

//ConnectDatabase performs the db connection
func ConnectDatabase() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successful connection to the Database")
	return client
}

//ConnectionCheck is the db ping
func ConnectionCheck() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
