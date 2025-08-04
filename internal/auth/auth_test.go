package auth

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mybirthday"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Hashed password doesn't match")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "yourbirthday"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("Error hashing password")
	}
	err = CheckPasswordHash(password, string(hash))
	if err != nil {
		t.Errorf("CheckPasswordHash function not accurate")
	}
}
