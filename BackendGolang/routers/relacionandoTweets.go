package routers

import (
	"net/http"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

//RelacionandoTweets relaciona usuarios de tweet
func RelacionandoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "el id de usuario tweet es requerido ", http.StatusBadRequest)
		return
	}

	var tweetRelacion models.Relaciones // relaciones tweet

	tweetRelacion.UsuarioId = IDUsuario
	tweetRelacion.UsuarioRelacionId = ID

	status, err := db.InsertoRelacion(tweetRelacion)

	if err != nil {
		http.Error(w, "No se logro realizar una relacion de los tweets "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "error inesperado al crear una relacion de tweets  ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
