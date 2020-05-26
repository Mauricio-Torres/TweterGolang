package db

import (
	"context"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificarRegistro modifica el registro del usuario toda su informacion */
func ModificarRegistro(usuario models.Usuario, ID string) (bool, error) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}

	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}

	registro["fechaNacimiento"] = usuario.FechaNacimiento

	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}

	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}

	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}

	if len(usuario.SitioWeb) > 0 {
		registro["sitioWeb"] = usuario.SitioWeb
	}

	updateStrin := bson.M{
		"$set": registro,
	}

	id, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": id}}

	_, err := coleccion.UpdateOne(contexto, filtro, updateStrin)

	if err != nil {
		return false, err
	}

	return true, nil
}
