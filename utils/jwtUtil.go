package utils

import (
	"github.com/golang-jwt/jwt"
	"scaleX/internal/dto"
	"time"
)

const (
	JWTSecretKey  = "JwtSecretKey"
	timeToExpired = 60
)

func GenerateJWTToken(user *dto.UserInfo) (string, error) {
	claims := &dto.UserClaimJwt{
		UserId:   user.UserId,
		Role:     user.Role,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * timeToExpired).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
