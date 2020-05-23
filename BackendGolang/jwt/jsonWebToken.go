package jwt

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	conv "github.com/dgrijalva/jwt-go"
)

/*GenerarToken genera token para el login del usuario*/
func GenerarToken(usuario models.Usuario) (string, error) {

	clave := []byte("generar token valido")

	// signKey, err := conv.ParseRSAPrivateKeyFromPEM(clave)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	payload := conv.MapClaims{
		"email":            usuario.Email,
		"nombre":           usuario.Nombre,
		"apellidos":        usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia":        usuario.Biografia,
		"sitioweb":         usuario.SitioWeb,
		"_id":              usuario.ID.Hex(),
		"exp":              json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10)),
	}

	token := conv.NewWithClaims(conv.SigningMethodHS256, payload)

	tokenSrt, err := token.SignedString(clave)
	if err != nil {
		fmt.Println(err.Error())
		return tokenSrt, err
	}

	return tokenSrt, nil
}
