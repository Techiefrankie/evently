package security

import (
	"evently/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const SecretKey = "hU8w7z!9@KqXy&3N4fT*5bV6pQ#A1RsGdL"

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
		"roles":   []string{"USER"},
	})
	return token.SignedString([]byte(SecretKey))
}
func GetEncryptedPassword(rawPassword string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	if err != nil {
		return "nil", err
	}

	return string(password), nil
}

func PasswordMatches(plainPassword, encryptedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(plainPassword))
	return err == nil
}
