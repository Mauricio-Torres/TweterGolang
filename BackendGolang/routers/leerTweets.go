package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mauricio-Torres/TweterGolang/db"
)

/*LeoTweets forma puente entre la firma del cervicio y la extraccion de la DB*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) // convierte string a numeros

	if err != nil {
		http.Error(w, "Error transformando string a num "+err.Error(), http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := db.LeoTweet(ID, pag)

	if correcto == false {
		http.Error(w, "Error al leer los tweets ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
