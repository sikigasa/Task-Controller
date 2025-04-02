package usecase

import (
	"context"

	"github.com/sikigasa/task-controller/internal/domain"
)

type TagService interface {
	CreateTag(ctx context.Context, param domain.CreateTagParam) (domain.Tag, error)
	ListTag(ctx context.Context, param domain.ListTagParam) ([]domain.Tag, error)
	DeleteTag(ctx context.Context, param domain.DeleteTagParam) error
}
