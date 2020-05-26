package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

//ConsultarRelacionEntreUsuarios consulta la relacion entre dos usuarios
func ConsultarRelacionEntreUsuarios(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "el id de usuario tweet es requerido ", http.StatusBadRequest)
		return
	}

	var tweetRelacion models.Relaciones // relaciones tweet

	tweetRelacion.UsuarioId = IDUsuario
	tweetRelacion.UsuarioRelacionId = ID

	var resultado models.ConsultarRelacionEntreUsuarios

	status, err := db.ConsultarRelacionesTweeter(tweetRelacion)

	if status == false || err != nil {
		resultado.Status = false
	} else {
		resultado.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}
