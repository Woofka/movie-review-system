package add

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	commandPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command"
	reviewPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
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
	return "add"
}

func (c *command) Description() string {
	return "create review"
}

func (c *command) Process(argsString string) string {
	if len(argsString) == 0 {
		return "No arguments were given. See /help for details."
	}
	args, err := commandPkg.ParseArguments(argsString)
	if err != nil {
		return "Invalid arguments. Make sure you don't use quotes in a quoted part."
	}
	if len(args) != 4 {
		return fmt.Sprintf("Invalid amount of arguments. Expected 4, but got %d instead.", len(args))
	}
	rating, err := strconv.Atoi(args[3])
	if err != nil {
		return "4th argument should be integer."
	}

	err = c.review.Create(models.Review{
		Reviewer:   args[0],
		MovieTitle: args[1],
		Text:       args[2],
		Rating:     uint8(rating),
	})
	if err != nil {
		if errors.Is(err, reviewPkg.ErrValidation) {
			return err.Error()
		}
		return "internal error"
	}

	return fmt.Sprintf("Review was successfully added.")
}
