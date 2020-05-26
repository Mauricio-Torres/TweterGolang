package routers

import (
	"net/http"

	"github.com/Mauricio-Torres/TweterGolang/db"
)

/*EliminarTweet borra los tweet del usuario logueado*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "debe enviar el id del tweet", http.StatusBadRequest)
	}

	err := db.BorrarTweet(ID, IDUsuario)

	if err != nil {
		http.Error(w, "error al borrar tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
