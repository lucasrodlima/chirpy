package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func TestJWT(t *testing.T) {
	userID := uuid.New()
	secret := "my-super-secret-secret"
	duration := time.Hour

	token, err := MakeJWT(userID, secret, duration)
	if err != nil {
		t.Errorf("error in jwt generation: %v", err)
	}

	returnID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Errorf("error validating jwt: %v", err)
	}

	if userID != returnID {
		t.Errorf("jwt functions returning wrong id")
	}
}

func TestJWTExpiration(t *testing.T) {
	userID := uuid.New()
	secret := "my-super-secret-secret"
	duration := time.Second * 1

	token, err := MakeJWT(userID, secret, duration)
	if err != nil {
		t.Errorf("error in jwt generation: %v", err)
	}
	time.Sleep(time.Second * 2)

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Errorf("jwt expiration not enforced")
	}
}

func TestJWTWrongSecret(t *testing.T) {
	userID := uuid.New()
	secret := "correct-secret"
	wrongSecret := "wrong-secret"
	duration := time.Hour

	token, err := MakeJWT(userID, secret, duration)
	if err != nil {
		t.Errorf("error in jwt generation: %v", err)
	}

	_, err = ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Errorf("jwt should be invalid with wrong secret")
	}
}

func TestBearerToken(t *testing.T) {
	token := "123456token"
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+token)

	newToken, err := GetBearerToken(headers)
	if err != nil {
		t.Errorf("error extracting bearer token")
	}

	if newToken != token {
		t.Errorf("GetBearerToken func not working properly")
	}
}
