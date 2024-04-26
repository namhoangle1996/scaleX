package usecase

import (
	"context"
	"errors"
	"scaleX/internal/dto"
	"scaleX/utils"
)

type authService struct {
}

func (a authService) Login(ctx context.Context, req dto.LoginReq) (res dto.LoginResp, err error) {
	userInfo, err := getUserByUserName(req.UserName)
	if err != nil {
		return res, err
	}

	token, err := utils.GenerateJWTToken(&userInfo)
	if err != nil {
		return res, err
	}
	res.Token = token

	return res, nil

}

func getUserByUserName(userName string) (res dto.UserInfo, err error) {
	res.UserId = userName

	switch userName {
	case "userA":
		res.Role = "regular"

	case "userB":
		res.Role = "admin"
	default:
		return dto.UserInfo{}, errors.New("Unknown User")
	}

	return res, nil

}

func NewAuthService() AuthService {
	return authService{}
}
