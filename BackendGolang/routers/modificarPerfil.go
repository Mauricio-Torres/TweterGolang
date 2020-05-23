package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.ModificarRegistro(usuario, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
