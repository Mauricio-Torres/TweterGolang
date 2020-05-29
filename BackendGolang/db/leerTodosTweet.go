package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LeerTodosUsuarios trae todos los usuarios que se relacionan con la persona logueada
func LeerTodosUsuarios(ID string, page int64, serch string, tipo string) ([]*models.Usuario, bool) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("usuarios")

	var result []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20) // el orden importa
	findOptions.SetLimit(20)

	//`(?i)` sirve para decirle que no importa entre minusculas y mayusculas
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + serch},
	}

	cursor, err := coleccion.Find(contexto, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	var encontrado, incluir bool

	for cursor.Next(contexto) {
		var us models.Usuario

		err := cursor.Decode(&us)

		if err != nil {
			fmt.Println(err.Error())
			return result, false
		}

		// se consultan las relaciones
		var relacionado models.Relaciones
		relacionado.UsuarioId = ID
		relacionado.UsuarioRelacionId = us.ID.Hex()

		incluir = false

		encontrado, err = ConsultarRelacionesTweeter(relacionado)

		if tipo == "new" && encontrado == false {
			incluir = true
		}

		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if relacionado.UsuarioRelacionId == ID {
			incluir = false
		}

		if incluir == true {

			us.Password = ""
			us.Biografia = ""
			us.SitioWeb = ""
			us.Banner = ""
			us.Email = ""

			result = append(result, &us)
		}
	}

	err = cursor.Err()

	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	cursor.Close(contexto)
	return result, true
}
