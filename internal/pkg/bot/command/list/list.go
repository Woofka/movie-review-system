package list

import (
	"strings"

	commandPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command"
	reviewPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review"
)

func New(review reviewPkg.Interface) commandPkg.Interface {
	return &command{
		review: review,
	}
}

type command struct {
	review reviewPkg.Interface
}

func (c *command) Name() string {
	return "list"
}

func (c *command) Description() string {
	return "show all reviews."
}

func (c *command) Process(_ string) string {
	reviews := c.review.List()

	if len(reviews) == 0 {
		return "No reviews yet"
	}

	result := make([]string, 0, len(reviews))
	for _, review := range reviews {
		result = append(result, review.String())
	}

	return strings.Join(result, "\n---\n")
}
