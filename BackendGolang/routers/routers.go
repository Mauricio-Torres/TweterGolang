package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Error en los datos de recibidos "+err.Error(), 400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "Email Usuario Requerido ", 400)
		return
	}

	if len(usuario.Password) < 4 {
		http.Error(w, "Password debe contener mas de 4 caracteres ", 400)
		return
	}

	_, encontrado, _ := db.ChequeoUsuario(usuario.Email)

	if encontrado {
		http.Error(w, "Usuario ya registrado ", 400)
		return
	}

	_, status, err := db.InsertarRegistro(usuario)

	if err != nil {
		http.Error(w, "No se pudo realizar el registro del Usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo realizar el registro del Usuario ", 400)
		return
	}

	if status {

		return
	}

	w.WriteHeader(http.StatusCreated)
}
