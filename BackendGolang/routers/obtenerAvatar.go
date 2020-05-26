package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Mauricio-Torres/TweterGolang/db"
)

// ObtenerAvatar obtiene la imagen del Avatar de los usuarios
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	perfilUs, err := db.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado ", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatar" + perfilUs.Avatar)

	if err != nil {
		http.Error(w, "Imagen no encontrada ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile) // esta sentencia envia en modo binario al archivo encontrado

	if err != nil {
		http.Error(w, "Error al copiar la Imagen ", http.StatusBadRequest)
	}
}
