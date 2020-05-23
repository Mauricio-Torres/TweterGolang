package models

/*GetTweet captura del body el mns que nos llega */
type GetTweet struct {
	Mensaje string `bson: "mensaje" json:"mensaje"`
}
