package repository

import (
	"context"
	"scaleX/internal/dto"
)

type UserRepo interface {
	GetUserById(ctx context.Context, userId string) (res dto.UserInfo, err error)
	GetUserByUserName(ctx context.Context, userName string) (res dto.UserInfo, err error)
}
