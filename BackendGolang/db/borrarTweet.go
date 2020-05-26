package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BorrarTweet borra los tweet segun el usuario conectado ...*/
func BorrarTweet(ID string, IDUser string) error {

	contexto, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoC.Database("tweter")
	coleccion := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": IDUser,
	}

	_, err := coleccion.DeleteOne(contexto, condicion)

	if err != nil {
		return err
	}

	return nil
}
