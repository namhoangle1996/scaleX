package repository

import (
	"context"
	"errors"
	"scaleX/internal/dto"
)

type userRepo struct {
}

func (u userRepo) GetUserById(ctx context.Context, userId string) (res dto.UserInfo, err error) {
	res.UserId = userId

	switch userId {
	case "regularUserId":
		res.Role = "regular"

	case "adminUserId":
		res.Role = "admin"
	default:
		return res, errors.New("Unknown User")
	}

	return res, nil

}

func (u userRepo) GetUserByUserName(ctx context.Context, userName string) (res dto.UserInfo, err error) {
	res.UserName = userName

	switch userName {
	case "regular":
		res.Role = "regular"
		res.UserId = "regularUserId"

	case "admin":
		res.Role = "admin"
		res.UserId = "adminUserId"

	default:
		return res, errors.New("Unknown User")
	}

	return res, nil

}

func NewUserRepo() UserRepo {
	return userRepo{}
}
