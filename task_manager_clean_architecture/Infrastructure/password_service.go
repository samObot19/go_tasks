package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)


func EncryptPassword(string password) (string, err) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil{
		return "", err
	}

	return string(hashedPassword), nil

}

