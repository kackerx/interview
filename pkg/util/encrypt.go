package util

import "golang.org/x/crypto/bcrypt"

func EncryptPass(password string) (string, error) {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 11)
	return string(encryptPassword), err
}

func ComparePass(password, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(raw))
}
