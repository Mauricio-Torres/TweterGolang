package db

import (
	"context"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// LeerTweetAllSeguidores lee todos los tweets de los seguidores
func LeerTweetAllSeguidores(ID string, pagina int) ([]models.DevolverTweetSeguidores, bool) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreingField": "userid",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": 1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := coleccion.Aggregate(contexto, condiciones)

	var result []models.DevolverTweetSeguidores

	err = cursor.All(contexto, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
