package usecase

import (
	"context"
	"scaleX/internal/dto"
	"scaleX/internal/repository"
	"scaleX/utils"
)

type authService struct {
	userRepo repository.UserRepo
}

func (a authService) Login(ctx context.Context, req dto.LoginReq) (res dto.LoginResp, err error) {
	userInfo, err := a.userRepo.GetUserByUserName(ctx, req.UserName)
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

func NewAuthService(userRepo repository.UserRepo) AuthService {
	return authService{
		userRepo: userRepo,
	}
}
