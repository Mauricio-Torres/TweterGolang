package db

import (
	"context"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertTweet inserta los tweet a la base de datos */
func InsertTweet(tweet models.Tweet) (string, bool, error) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("tweet")

	registro := bson.M{
		"userid":  tweet.UserId,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}

	result, err := coleccion.InsertOne(contexto, registro)

	if err != nil {
		return "", false, err
	}

	ObjetID := result.InsertedID.(primitive.ObjectID).Hex()

	return ObjetID, true, nil
}
