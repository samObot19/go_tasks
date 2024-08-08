package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)


func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil{
		return "", err
	}

	return string(hashedPassword), nil

}
func IsValidPassword(loggedUser string, user string) bool{
        if bcrypt.CompareHashAndPassword([]byte(loggedUser), []byte(user)) != nil {
                return false
        }

        return true
}

