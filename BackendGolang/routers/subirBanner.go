package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

// SubirBanner almacena una imagen en el sistema
func SubirBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar") // obtiene un archivo de la ruta

	var extension = strings.Split(handler.Filename, ".")[1] // obtiene la extension del archivo

	var archivo string = "uploads/banner" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	// var status bool

	usuario.Banner = IDUsuario + "." + extension

	status, err := db.ModificarRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "No se pudo actualizar el perfil de usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
