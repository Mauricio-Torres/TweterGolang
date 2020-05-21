package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Mauricio-Torres/TweterGolang/db"
)

/*VerPerfil  retorna un json del perfil del usuario*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Se encontro un error al buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
