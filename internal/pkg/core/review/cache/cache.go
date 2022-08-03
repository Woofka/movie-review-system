package cache

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
)

var ErrReviewNotExists = errors.New("review does not exist")

type Interface interface {
	Add(review *models.Review) error
	Delete(id uint) error
	Update(review *models.Review) error
	List() []*models.Review
	Get(id uint) (*models.Review, error)
}
