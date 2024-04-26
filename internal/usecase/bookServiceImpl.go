package usecase

import (
	"context"
	"scaleX/internal/dto"
	"scaleX/internal/repository"
)

type bookService struct {
	userRepo repository.UserRepo
}

func (b bookService) FetchBook(ctx context.Context, userId string) (*dto.FetchBookResp, error) {
	userInfo, err := b.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	switch userInfo.Role {
	case "regular":

	case "admin":

	}

	panic("implement me")
}

func NewBookService() BookService {
	return &bookService{}
}
