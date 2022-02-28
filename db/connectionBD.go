package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Mongo CN Var*/
var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://cluster0.ylcqq.mongodb.net/myFirstDatabase?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=./credentials/X509-cert-486224665158591459.pem")

/*Connect DB: Function allow to connect with the Database*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Database Conection Successfully")
	return client
}

/*CheckConnection allow check out the connection with the database.*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
