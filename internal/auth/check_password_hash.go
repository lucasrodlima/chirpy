package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
