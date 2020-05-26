package routers

import (
	"net/http"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

// EliminarRelacionTweeter elimina la relacion de usuarios de tweets
func EliminarRelacionTweeter(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "el id de usuario tweet es requerido ", http.StatusBadRequest)
		return
	}

	var tweetRelacion models.Relaciones // relaciones tweet

	tweetRelacion.UsuarioId = IDUsuario
	tweetRelacion.UsuarioRelacionId = ID

	status, err := db.BorrarRelacionTweet(tweetRelacion)

	if err != nil {
		http.Error(w, "No pudo eliminar la relacion entre los tweeters "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "error inesperado al eliminar una relacion de tweets  ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
