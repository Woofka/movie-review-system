package get

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"
	commandPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command"
	reviewPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review"
	cachePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache"
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
	return "get"
}

func (c *command) Description() string {
	return "show review. Usage: `/get <id>`."
}

func (c *command) Process(ctx context.Context, argsString string) string {
	if len(argsString) == 0 {
		return "No arguments were given. See /help for details."
	}
	args, err := commandPkg.ParseArguments(argsString)
	if err != nil {
		return "Invalid arguments. Make sure you don't use quotes in a quoted part."
	}
	if len(args) != 1 {
		return fmt.Sprintf("Invalid amount of arguments. Expected 1, but got %d instead.", len(args))
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return "Argument should be integer."
	}

	review, err := c.review.Get(ctx, uint(id))
	if err != nil {
		if errors.Is(err, reviewPkg.ErrValidation) {
			return err.Error()
		}
		if errors.Is(err, cachePkg.ErrReviewNotExists) {
			return "Review does not exist"
		}
		log.Print(err)
		return "internal error"
	}

	return review.String()
}
