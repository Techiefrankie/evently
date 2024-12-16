package util

import "golang.org/x/crypto/bcrypt"

func GetEncryptedPassword(rawPassword string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	if err != nil {
		return "nil", err
	}

	return string(password), nil
}
