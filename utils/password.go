package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	cost:=14
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),cost)
	return string(bytes),err
}

func CompareHashedPassword(password string,hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
		return false
    }
	return true
}
