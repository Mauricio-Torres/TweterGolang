package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

// SubirAvatar almacena una imagen en el sistema
func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar") // obtiene un archivo de la ruta

	var extension = strings.Split(handler.Filename, ".")[1] // obtiene la extension del archivo

	crearDirectorioSiNoExiste("uploads/avatar")

	var archivo string = "uploads/avatar" + IDUsuario + "." + extension

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

	usuario.Avatar = IDUsuario + "." + extension

	status, err := db.ModificarRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "No se pudo actualizar el perfil de usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// exists returns whether the given file or directory exists or not

func exists(path string) (bool, error) {

	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// crearDirectorioSiNoExiste crea un directorio si no existe y coloca permisos para carpetas.
// El propietario, el grupo y otros pueden leer el directorio, pero solo el propietario puede cambiar su contenido.
func crearDirectorioSiNoExiste(directorio string) {

	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.Mkdir(directorio, 0755)

		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	}
}
