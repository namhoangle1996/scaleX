package usecase

import (
	"context"
	"scaleX/internal/dto"
)

type AuthService interface {
	Login(ctx context.Context, userInfo dto.LoginReq) (dto.LoginResp, error)
}
