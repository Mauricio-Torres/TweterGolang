package models

// Relaciones relaciona los tweet del usuario con los demas usuarios en la Db
type Relaciones struct {
	UsuarioId         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionId string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
