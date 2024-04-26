package dto

import "github.com/golang-jwt/jwt"

type UserClaimJwt struct {
	UserId   string `json:"user_id"`
	Role     string `json:"role"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}
