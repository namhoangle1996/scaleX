package usecase

import (
	"context"
	"scaleX/internal/dto"
)

type BookService interface {
	FetchBook(ctx context.Context, userId string) (dto.FetchBookResp, error)
	AddBook(ctx context.Context, request dto.AddBookRequest) error
}
