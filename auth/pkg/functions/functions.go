package functions

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordMatch(passwordRequest string, passwordDB string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordRequest), []byte(passwordDB))
	if err != nil {
		return false
	}
	return true
}

func GenerateJWTSecret(length int) (string, error) {
	randomByte := make([]byte, length)
	_, err := rand.Read(randomByte)
	if err != nil {
		return "Error al crear el byte", err
	}

	JWTString := base64.URLEncoding.EncodeToString(randomByte)
	return JWTString, nil
}
