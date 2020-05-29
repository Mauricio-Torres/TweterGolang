package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mauricio-Torres/TweterGolang/db"
)

// ListaUsuarios retorna los usuarios relacionados o no relacionados con el usuario logueado
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	tipoUsuarios := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	serch := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina ", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := db.LeerTodosUsuarios(IDUsuario, pag, serch, tipoUsuarios)

	if status == false {
		http.Error(w, "Ocurrio un error inesperado consultando los datos ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
