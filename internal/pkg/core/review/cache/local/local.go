package local

import (
	"context"
	"sort"
	"sync"

	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
)

const pollSize = 10

func New() cachePkg.Interface {
	return &cache{
		mu:     sync.RWMutex{},
		data:   map[uint]*models.Review{},
		lastId: uint(0),
		poolCh: make(chan struct{}, pollSize),
	}
}

type cache struct {
	mu     sync.RWMutex
	data   map[uint]*models.Review
	lastId uint
	poolCh chan struct{}
}

func (c *cache) List(_ context.Context, limit, offset uint, orderDesc bool) ([]*models.Review, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
	}()

	result := make([]*models.Review, 0, len(c.data))
	for _, v := range c.data {
		result = append(result, v)
	}
	if orderDesc {
		sort.Slice(result, func(i, j int) bool { return result[i].Id > result[j].Id })
	} else {
		sort.Slice(result, func(i, j int) bool { return result[i].Id < result[j].Id })
	}
	return result[offset : offset+limit], nil
}

func (c *cache) Add(_ context.Context, review *models.Review) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	c.lastId++
	review.Id = c.lastId
	c.data[review.Id] = review
	return nil
}

func (c *cache) Get(_ context.Context, id uint) (*models.Review, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
	}()

	review, ok := c.data[id]
	if !ok {
		return &models.Review{}, errors.Wrapf(cachePkg.ErrReviewNotExists, "review with id %d does not exists", id)
	}
	return review, nil
}

func (c *cache) Update(_ context.Context, review *models.Review) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	if _, ok := c.data[review.Id]; !ok {
		return errors.Wrapf(cachePkg.ErrReviewNotExists, "review with id %d does not exists", review.Id)
	}
	c.data[review.Id] = review
	return nil
}

func (c *cache) Delete(_ context.Context, id uint) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	if _, ok := c.data[id]; !ok {
		return errors.Wrapf(cachePkg.ErrReviewNotExists, "review with id %d does not exists", id)
	}
	delete(c.data, id)
	return nil
}
