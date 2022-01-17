package services

import (
	"crm/pkg/config"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte(config.GetString("JWT_SECRET", "jwtSecret"))

type Claims struct {
	UserId     uint64    `json:"userId"`
	SystemType string `json:"systemType"`
	jwt.StandardClaims
}

type Jwt struct {
}

func (j *Jwt) GenerateToken(userId uint64, systemType string) (string,error) {
	claims := Claims{
		userId,
		systemType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenClaims.SignedString(jwtSecret)
}

func (j *Jwt) ParseToken(token string) (*Claims,error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}