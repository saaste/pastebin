package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/saaste/pastebin/pkg/config"
)

type JwtParser struct {
	appConfig *config.AppConfig
}

func NewJwtParser(appConfig *config.AppConfig) *JwtParser {
	return &JwtParser{
		appConfig: appConfig,
	}
}

func (j *JwtParser) CreateJWT() (string, error) {
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(j.appConfig.JwtSecret))
	return tokenString, err
}

func (j *JwtParser) ParseJWT(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.appConfig.JwtSecret), nil
	})
	if err != nil {
		return fmt.Errorf("invalid token: %v", err)
	}

	return nil
}
