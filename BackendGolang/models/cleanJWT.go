package models

import (
	jwtl "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Clain estructura usada para procesar los JWT*/
type Clain struct {
	Email string             `json: "email"`
	ID    primitive.ObjectID `bson: "_id" json: "_id, omitempty"`
	jwtl.StandardClaims
}
