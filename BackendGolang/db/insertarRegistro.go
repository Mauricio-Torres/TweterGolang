package db

import (
	"context"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarRegistro inserta un usuario en la DB */
func InsertarRegistro(usuario models.Usuario) (string, bool, error) {

	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("usuarios")

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := coleccion.InsertOne(contexto, usuario)

	if err != nil {
		return "", false, err
	}

	ObjetID := result.InsertedID.(primitive.ObjectID).Hex()
	return ObjetID, true, nil
}

/*ChequeoUsuario chequeo de existencia de usuario en la DB */
func ChequeoUsuario(email string) (models.Usuario, bool, string) {
	contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoC.Database("tweter")
	coleccion := dataBase.Collection("usuarios")

	condision := bson.M{"email": email}

	var usuario models.Usuario

	err := coleccion.FindOne(contexto, condision).Decode(&usuario)
	ID := usuario.ID.Hex()

	if err != nil {

		return usuario, false, ID
	}

	return usuario, true, ID
}
