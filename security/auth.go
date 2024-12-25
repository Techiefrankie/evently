package security

import (
	"errors"
	"evently/api"
	"evently/config"
	"evently/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var SecretKey = config.GetEnv(config.SecretKey, "")

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
		"roles":   []string{"USER"},
	})
	return token.SignedString([]byte(SecretKey))
}

func ValidateToken(token string) (api.AuthResponse, error) {
	if token == "" {
		return api.AuthResponse{}, errors.New("access token is required")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return api.AuthResponse{}, errors.New("could not parse token")
	}

	if parsedToken.Valid {
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if ok {
			if time.Unix(int64(claims["exp"].(float64)), 0).Sub(time.Now()) > 0 {
				return api.AuthResponse{
					Email:  claims["email"].(string),
					UserId: int(claims["user_id"].(float64)),
					//Roles:     claims["roles"].([]string),
					//ExpiresIn: claims["exp"].(int) - int(time.Now().Unix()),
				}, nil
			} else {
				return api.AuthResponse{}, errors.New("token expired")
			}
		} else {
			return api.AuthResponse{}, errors.New("invalid token")
		}
	} else {
		return api.AuthResponse{}, errors.New("invalid token")
	}
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
