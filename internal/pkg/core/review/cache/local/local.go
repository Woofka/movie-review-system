package local

import (
	"sort"
	"sync"

	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
)

func New() cachePkg.Interface {
	return &cache{
		mu:     sync.RWMutex{},
		data:   map[uint]models.Review{},
		lastId: uint(0),
	}
}

type cache struct {
	mu     sync.RWMutex
	data   map[uint]models.Review
	lastId uint
}

func (c *cache) List() []models.Review {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]models.Review, 0, len(c.data))
	for _, v := range c.data {
		result = append(result, v)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Id < result[j].Id })
	return result
}

func (c *cache) Add(review models.Review) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lastId++
	review.Id = c.lastId
	c.data[review.Id] = review
	return nil
}

func (c *cache) Get(id uint) (models.Review, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	review, ok := c.data[id]
	if !ok {
		return models.Review{}, errors.Wrapf(cachePkg.ErrReviewNotExists, "review with id %d does not exists", id)
	}
	return review, nil
}

func (c *cache) Update(review models.Review) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[review.Id]; !ok {
		return errors.Wrapf(cachePkg.ErrReviewNotExists, "review with id %d does not exists", review.Id)
	}
	c.data[review.Id] = review
	return nil
}

func (c *cache) Delete(id uint) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[id]; !ok {
		return errors.Wrapf(cachePkg.ErrReviewNotExists, "review with id %d does not exists", id)
	}
	delete(c.data, id)
	return nil
}
