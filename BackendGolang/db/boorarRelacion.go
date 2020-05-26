package db

import (
	"context"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
)

// BorrarRelacionTweet borra la relacion entre tweeter existentes
func BorrarRelacionTweet(relacion models.Relaciones) (bool, error) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("relacion")

	_, err := coleccion.DeleteOne(contexto, relacion)

	if err != nil {
		return false, err
	}

	return true, nil
}
