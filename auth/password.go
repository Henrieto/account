package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

func ComparePassword(HashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(HashedPassword, password)
	return err == nil
}
