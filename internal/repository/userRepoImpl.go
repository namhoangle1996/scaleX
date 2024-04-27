package repository

import (
	"context"
	"errors"
	"scaleX/internal/constants"
	"scaleX/internal/dto"
)

type userRepo struct {
}

func (u userRepo) GetUserById(ctx context.Context, userId string) (res dto.UserInfo, err error) {
	res.UserId = userId

	switch userId {
	case constants.REGULAR_USER_ID:
		res.Role = constants.REGULAR_ROLE

	case constants.ADMIN_USER_ID:
		res.Role = constants.ADMIN_ROLE
	default:
		return res, errors.New("Unknown User")
	}

	return res, nil

}

func (u userRepo) GetUserByUserName(ctx context.Context, userName string) (res dto.UserInfo, err error) {
	res.UserName = userName

	switch userName {
	case constants.REGULAR_ROLE:
		res.Role = constants.REGULAR_ROLE
		res.UserId = constants.REGULAR_USER_ID

	case constants.ADMIN_ROLE:
		res.Role = constants.ADMIN_ROLE
		res.UserId = constants.ADMIN_USER_ID

	default:
		return res, errors.New("Unknown User")
	}

	return res, nil

}

func NewUserRepo() UserRepo {
	return userRepo{}
}
