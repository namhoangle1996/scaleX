package usecase

import (
	"context"
)

type OpenApiService interface {
	SummarizeChapters(ctx context.Context) error
}
