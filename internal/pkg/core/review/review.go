package review

import (
	"context"

	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache"
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
	Create(ctx context.Context, review *models.Review) error
	Update(ctx context.Context, review *models.Review) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*models.Review, error)
	Get(ctx context.Context, id uint) (*models.Review, error)
}

func New(cache cachePkg.Interface) Interface {
	return &core{
		cache: cache,
	}
}

type core struct {
	cache cachePkg.Interface
}

func (c *core) Create(ctx context.Context, review *models.Review) error {
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
		return errors.Wrapf(ErrValidation, "invalid rating value: %d. Should be 0..10", review.Rating)
	}

	return c.cache.Add(ctx, review)
}

func (c *core) Update(ctx context.Context, review *models.Review) error {
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
		return errors.Wrapf(ErrValidation, "invalid rating value: %d. Should be 0..10", review.Rating)
	}

	return c.cache.Update(ctx, review)
}

func (c *core) Delete(ctx context.Context, id uint) error {
	return c.cache.Delete(ctx, id)
}

func (c *core) Get(ctx context.Context, id uint) (*models.Review, error) {
	return c.cache.Get(ctx, id)
}

func (c *core) List(ctx context.Context) ([]*models.Review, error) {
	return c.cache.List(ctx)
}
