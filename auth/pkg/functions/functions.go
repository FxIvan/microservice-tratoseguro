package functions

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordMatch(passwordRequest string, passwordDB string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordRequest), []byte(passwordDB))
	if err != nil {
		return false
	}
	return true
}
