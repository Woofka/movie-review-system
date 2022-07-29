package cache

import (
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
)

type Interface interface {
	Add(review models.Review) error
	Delete(id uint) error
	Update(review models.Review) error
	List() []models.Review
	Get(id uint) (models.Review, error)
}
