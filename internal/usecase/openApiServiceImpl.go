package usecase

import (
	"context"
)

type openApiService struct {
}

func (b openApiService) SummarizeChapters(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewOpenApiService() OpenApiService {
	return &openApiService{}
}
