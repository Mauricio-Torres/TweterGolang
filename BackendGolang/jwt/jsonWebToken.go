package jwt

import (
	"time"

	"github.com/Mauricio-Torres/TweterGolang/models"
	conv "github.com/dgrijalva/jwt-go"
)

/*GenerarToken genera token para el login del usuario*/
func GenerarToken(usuario models.Usuario) (string, error) {

	clave := []byte("generarToken")

	payload := conv.MapClaims{
		"email":            usuario.Email,
		"nombre":           usuario.Nombre,
		"apellidos":        usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia":        usuario.Biografica,
		"sitioweb":         usuario.SitioWeb,
		"_id":              usuario.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := conv.NewWithClaims(conv.SigningMethodES256, payload)

	tokenSrt, err := token.SignedString(clave)
	if err != nil {
		return tokenSrt, err
	}

	return tokenSrt, nil
}
