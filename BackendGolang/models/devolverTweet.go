package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevolverTweet devuelve los tweet de un usuario en particular */
type DevolverTweet struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id, omitempty"`
	UserId  string             `bson:"userid" json:"userId, omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje, omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha, omitempty"`
}
