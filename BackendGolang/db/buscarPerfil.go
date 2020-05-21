package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca el perfil dentro de la base de datos */
func BuscoPerfil(ID string) (models.Usuario, error) {

	contexto, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoC.Database("tweter")
	coleccion := db.Collection("usuarios")

	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := coleccion.FindOne(contexto, condicion).Decode(&perfil)

	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}
	perfil.Password = ""

	return perfil, nil
}
