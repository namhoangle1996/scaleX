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
	case "userA":
		res.Role = "regular"

	case "userB":
		res.Role = "admin"
	default:
		return res, errors.New("Unknown User")
	}

	return res, nil

}

func NewUserRepo() UserRepo {
	return userRepo{}
}
