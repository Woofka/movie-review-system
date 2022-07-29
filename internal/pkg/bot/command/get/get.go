package get

import (
	"fmt"
	"strconv"

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
	return "get"
}

func (c *command) Description() string {
	return "show review"
}

func (c *command) Process(argsString string) string {
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

	review, err := c.review.Get(uint(id))
	if err != nil {
		// if errors.Is(err, reviewPkg.ErrValidation) { return "invalid args" }
		// TODO: error handling
		return "internal error"
	}

	return review.String()
}
