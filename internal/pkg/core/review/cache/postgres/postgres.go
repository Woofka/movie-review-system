package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	cachePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
)

func New(pool *pgxpool.Pool) cachePkg.Interface {
	return &Repository{pool}
}

type Repository struct {
	pool *pgxpool.Pool
}

func (r *Repository) List(ctx context.Context) []*models.Review {
	return nil
}

func (r *Repository) Add(ctx context.Context, review *models.Review) error {
	return nil
}

func (r *Repository) Get(ctx context.Context, id uint) (*models.Review, error) {
	return nil, nil
}

func (r *Repository) Update(ctx context.Context, review *models.Review) error {
	return nil
}

func (r *Repository) Delete(ctx context.Context, id uint) error {
	return nil
}
