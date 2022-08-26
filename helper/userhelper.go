package helper

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

//Encrypt Password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

//Check Password
func CheckPassword(providedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return err == nil
	}
	return true
}

func ValidMailAddress(email string) (string, bool) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}
