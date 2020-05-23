package routers

import (
	"errors"
	"strings"

	"github.com/Mauricio-Torres/TweterGolang/db"
	"github.com/Mauricio-Torres/TweterGolang/models"
	jwtTC "github.com/dgrijalva/jwt-go"
)

/*Email extraido del token */
var Email string

/*IDUsuario extraido del token */
var IDUsuario string

/*ProcesoToken extrae todos los valores del token enviado */
func ProcesoToken(token string) (*models.Clain, bool, string, error) {

	clave := []byte("generar token valido") // clave para generar token

	claims := &models.Clain{}

	splitToken := strings.Split(token, "Bearer") // divide al string en 2 en la parte de vearer y el token sifrado

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwtTC.ParseWithClaims(token, claims, func(token2 *jwtTC.Token) (interface{}, error) {
		return clave, nil
	})

	if err == nil {
		_, encontrado, _ := db.ChequeoUsuario(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
