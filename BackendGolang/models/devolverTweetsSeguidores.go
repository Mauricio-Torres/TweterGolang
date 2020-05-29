package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DevolverTweetSeguidores retorna los tweets de los seguidores
type DevolverTweetSeguidores struct {
	ID               primitive.ObjectID `bson: "_id" json:"_id,omitempty"`
	UserId           string             `bson: "userid" json:"userid,omitempty"`
	UserioRelacionId string             `bson: "usuariorelacionid" json:"usuariorelacionid,omitempty"`
	Tweet            struct {
		Mensaje string    `bson: "mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson: "fecha" json:"fecha,omitempty"`
		ID      string    `bson: "_id" json:"_id,omitempty"`
	}
}
