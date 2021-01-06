package tools

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword Hash Password
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

//CompareHashedPassword Comparator
func CompareHashedPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
