package list

import (
	"context"
	"fmt"
	"log"
	"strconv"
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
	return "show all reviews. Usage: `/list <limit> <offset> <order_desc>`."
}

func (c *command) Process(ctx context.Context, argsString string) string {
	if len(argsString) == 0 {
		return "No arguments were given. See /help for details."
	}
	args, err := commandPkg.ParseArguments(argsString)
	if err != nil {
		return "Invalid arguments. Make sure you don't use quotes in a quoted part."
	}
	if len(args) != 3 {
		return fmt.Sprintf("Invalid amount of arguments. Expected 3, but got %d instead.", len(args))
	}

	limit, err := strconv.Atoi(args[0])
	if err != nil {
		return "1st argument should be integer."
	}
	if limit < 0 {
		return "1st argument should be positive."
	}

	offset, err := strconv.Atoi(args[1])
	if err != nil {
		return "2nd argument should be integer."
	}
	if offset < 0 {
		return "2nd argument should be positive."
	}

	var orderDesc bool
	switch strings.ToLower(args[2]) {
	case "true":
		orderDesc = true
	case "false":
		orderDesc = false
	default:
		return "3rd argument should be true or false."
	}

	reviews, err := c.review.List(ctx, uint(limit), uint(offset), orderDesc)
	if err != nil {
		log.Print(err)
		return "internal error"
	}

	if len(reviews) == 0 {
		return "No reviews yet"
	}

	result := make([]string, 0, len(reviews))
	for _, review := range reviews {
		result = append(result, review.String())
	}

	return strings.Join(result, "\n---\n")
}
