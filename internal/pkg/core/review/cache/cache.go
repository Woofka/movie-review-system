package cache

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
)

var ErrReviewNotExists = errors.New("review does not exist")

type Interface interface {
	Add(ctx context.Context, review *models.Review) error
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, review *models.Review) error
	List(ctx context.Context) ([]*models.Review, error)
	Get(ctx context.Context, id uint) (*models.Review, error)
}
