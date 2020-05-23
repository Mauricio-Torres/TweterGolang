package db

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword encripta el pass del usuario*/
func EncriptarPassword(password string) (string, error) {
	costo := 8

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
