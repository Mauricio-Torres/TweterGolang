package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario coleccion de DB*/
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Nombre          string             `bson:"nombre, " json:"nombre, omitempty"`
	Apellidos       string             `bson:"apellido, omitempty" json:"apellido, omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento, omitempty" json:"fechaNacimiento, omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password, omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar, omitempty"`
	Banner          string             `bson:"banner" json:"banner, omitempty"`
	Biografica      string             `bson:"biografia" json:"biografia, omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb, omitempty"`
}
