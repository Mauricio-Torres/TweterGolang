package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ConsultarRelacionesTweeter consulta la relacion entre 2 usuarios
func ConsultarRelacionesTweeter(relacion models.Relaciones) (bool, error) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         relacion.UsuarioId,
		"usuariorelacionid": relacion.UsuarioRelacionId,
	}

	var resultado models.Relaciones

	fmt.Println(resultado)

	err := coleccion.FindOne(contexto, condicion).Decode(resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, err
}
