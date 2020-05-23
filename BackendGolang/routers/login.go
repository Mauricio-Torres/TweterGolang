package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/jwt"
	"github.com/Mauricio-Torres/TweterGolang/models"
)

/*Loguer router para Loguin del usuario*/
func Loguer(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Password y/o Usuario invalido"+err.Error(), 400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "Email es requerido", 400)
		return
	}

	userRetorno, existe := db.Loguin(usuario.Email, usuario.Password)

	fmt.Println(userRetorno)
	if existe == false {
		http.Error(w, "Password y/o Usuario invalido", 400)
		return
	}

	jsonWebToken, err := jwt.GenerarToken(userRetorno)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jsonWebToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jsonWebToken,
		Expires: expirationTime,
	})

}
