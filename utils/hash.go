package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}

func ComparePassword(password, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}
