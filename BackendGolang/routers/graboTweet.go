package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

func GrabarTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.GetTweet // mensaje que llega de front para ser almacenado

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.Tweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertTweet(registro)

	if err != nil {
		http.Error(w, "ocurrio un error al intentar insertar el registro, intente nuevamente"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "no se logro guardar el tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
