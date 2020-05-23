package db

import (
	"context"
	"log"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweet sirve como puente entre la base de datos y el middware */
func LeoTweet(Id string, pagina int64) ([]*models.DevolverTweet, bool) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("tweet")

	var resultado []*models.DevolverTweet

	condicion := bson.M{
		"userid": Id,
	}

	// options permite definir y filtrar comportamientos en la base de datos mongo
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := coleccion.Find(contexto, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevolverTweet

		err := cursor.Decode(&registro)

		if err != nil {
			return resultado, false
		}

		resultado = append(resultado, &registro)
	}

	return resultado, true
}
