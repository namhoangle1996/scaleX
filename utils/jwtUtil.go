package utils

import (
	"github.com/golang-jwt/jwt"
	"scaleX/internal/dto"
	"time"
)

func GenerateJWTToken(user *dto.UserInfo) (string, error) {
	claims := &dto.UserClaimJwt{
		UserId:   user.UserId,
		Role:     user.Role,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("JwtSecretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
