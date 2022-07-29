package update

import (
	"fmt"
	"strconv"

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
	return "update"
}

func (c *command) Description() string {
	return "update review"
}

func (c *command) Process(argsString string) string {
	if len(argsString) == 0 {
		return "No arguments were given. See /help for details."
	}
	args, err := commandPkg.ParseArguments(argsString)
	if err != nil {
		return "Invalid arguments. Make sure you don't use quotes in a quoted part."
	}
	if len(args) != 5 {
		return fmt.Sprintf("Invalid amount of arguments. Expected 5, but got %d instead.", len(args))
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return "1th argument should be integer."
	}

	rating, err := strconv.Atoi(args[4])
	if err != nil {
		return "4th argument should be integer."
	}

	err = c.review.Update(models.Review{
		Id:         uint(id),
		Reviewer:   args[1],
		MovieTitle: args[2],
		Text:       args[3],
		Rating:     uint8(rating),
	})
	if err != nil {
		// if errors.Is(err, reviewPkg.ErrValidation) { return "invalid args" }
		// TODO: error handling
		return "internal error"
	}

	return fmt.Sprintf("Review was successfully updated.")
}
