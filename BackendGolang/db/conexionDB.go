package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC conecta a la Data Base */
var MongoC = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://prueba:Aleja2014.@cluster0-syui7.mongodb.net/test?retryWrites=true&w=majority")

/*ConectarDB conecta a la base de datos  */
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	// realiza un pin a la DB si esta arriba o no ...
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa a la DB ")

	return client
}

/*CheckConnection verifica si la conexion a la base de datos esta bien*/
func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	return 1
}
