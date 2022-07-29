package review

import (
	"fmt"

	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache"
	localCachePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache/local"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
)

const (
	maxReviewerLen   = 32
	maxMovieTitleLen = 50
	maxTextLen       = 200
	maxRating        = 10
)

var ErrValidation = errors.New("invalid data")

type Interface interface {
	Create(review models.Review) error
	Update(review models.Review) error
	Delete(id uint) error
	List() []models.Review
	Get(id uint) (models.Review, error)
}

func New() Interface {
	return &core{
		cache: localCachePkg.New(),
	}
}

type core struct {
	cache cachePkg.Interface
}

func (c *core) Create(review models.Review) error {
	if review.Reviewer == "" || len(review.Reviewer) > maxReviewerLen {
		return errors.Wrapf(ErrValidation, "invalid reviewer length: %d. Should be 1..32", len(review.Reviewer))
	}

	if review.MovieTitle == "" || len(review.MovieTitle) > maxMovieTitleLen {
		return errors.Wrapf(ErrValidation, "invalid movie title length: %d. Should be 1..50", len(review.MovieTitle))
	}

	if review.Text == "" || len(review.Text) > maxTextLen {
		return errors.Wrapf(ErrValidation, "invalid text length: %d. Should be 1..200", len(review.Text))
	}

	if review.Rating > maxRating {
		return fmt.Errorf("invalid rating value: %d. Should be 0..10", review.Rating)
	}

	return c.cache.Add(review)
}

func (c *core) Update(review models.Review) error {
	if review.Reviewer == "" || len(review.Reviewer) > maxReviewerLen {
		return errors.Wrapf(ErrValidation, "invalid reviewer length: %d. Should be 1..32", len(review.Reviewer))
	}

	if review.MovieTitle == "" || len(review.MovieTitle) > maxMovieTitleLen {
		return errors.Wrapf(ErrValidation, "invalid movie title length: %d. Should be 1..50", len(review.MovieTitle))
	}

	if review.Text == "" || len(review.Text) > maxTextLen {
		return errors.Wrapf(ErrValidation, "invalid text length: %d. Should be 1..200", len(review.Text))
	}

	if review.Rating > maxRating {
		return fmt.Errorf("invalid rating value: %d. Should be 0..10", review.Rating)
	}

	return c.cache.Update(review)
}

func (c *core) Delete(id uint) error {
	return c.cache.Delete(id)
}

func (c *core) Get(id uint) (models.Review, error) {
	return c.cache.Get(id)
}

func (c *core) List() []models.Review {
	return c.cache.List()
}
