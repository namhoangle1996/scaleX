package usecase

import (
	"context"
)

type OpenApiService interface {
	SummarizeChapters(ctx context.Context) (res map[string]string, error error)
}
