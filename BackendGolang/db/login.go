package db

import (
	"github.com/Mauricio-Torres/TweterGolang/models"
	"golang.org/x/crypto/bcrypt"
)

/*Loguin metodo que permite loguear al usuario */
func Loguin(email string, password string) (models.Usuario, bool) {

	usuario, encontrado, _ := ChequeoUsuario(email)

	if !encontrado {
		return usuario, false
	}

	passwordByte := []byte(password)
	passwordByteDB := []byte(usuario.Password)
	err := bcrypt.CompareHashAndPassword(passwordByteDB, passwordByte)

	if err != nil {
		return usuario, false
	}

	return usuario, true
}
