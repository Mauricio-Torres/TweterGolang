package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mauricio-Torres/TweterGolang/db"
)

// LeoTweetsSeguidores trae todos los tweets de los seguidores
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Error transformando string a num "+err.Error(), http.StatusBadRequest)
		return
	}

	respuesta, correcto := db.LeerTweetAllSeguidores(IDUsuario, pagina)

	if correcto == false {
		http.Error(w, "Error al traer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
