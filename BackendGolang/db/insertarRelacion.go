package db

import (
	"context"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
)

//InsertoRelacion incerta la relacion del usuario y los nuevos tweet q tiene
func InsertoRelacion(t models.Relaciones) (bool, error) {
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("relacion")

	_, err := coleccion.InsertOne(contexto, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
