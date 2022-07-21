package handlers

import (
	"encoding/csv"
	"fmt"
	"movie-review-system/internal/commander"
	"movie-review-system/internal/storage"
	"strconv"
	"strings"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
	addCmd    = "add"
	updateCmd = "update"
	deleteCmd = "delete"
)

func parseArgs(args string) ([]string, error) {
	reader := csv.NewReader(strings.NewReader(args))
	reader.Comma = ' '
	res, err := reader.Read()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func helpHandler(_ string) string {
	return "/help - show this message.\n" +
		"/list - show all reviews.\n" +
		"/add `<reviewer> \"<movie title>\" \"<review text>\" <rating>` - add new review. " +
		"`<rating>` should be an integer between 0 and 10.\n" +
		"/update `<id> <reviewer> \"<movie title>\" \"<review text>\" <rating>` - update review.\n" +
		"/delete `<id>` - removes review."
}

func listHandler(_ string) string {
	reviews := storage.List()
	if len(reviews) > 0 {
		res := make([]string, 0, len(reviews))
		for _, review := range reviews {
			res = append(res, review.String())
		}
		return strings.Join(res, "\n---\n")
	} else {
		return "No reviews yet"
	}
}

func addHandler(s string) string {
	if len(s) == 0 {
		return "No arguments were given. See /help for details."
	}
	args, err := parseArgs(s)
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
	r, err := storage.NewReview(args[0], args[1], args[2], uint8(rating))
	if err != nil {
		return err.Error()
	}
	err = storage.Add(r)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Review #%d was successfully added.", r.GetId())
}

func updateHandler(s string) string {
	if len(s) == 0 {
		return "No arguments were given. See /help for details."
	}
	args, err := parseArgs(s)
	if err != nil {
		return "Invalid arguments. Make sure you don't use quotes in a quoted part."
	}
	if len(args) != 5 {
		return fmt.Sprintf("Invalid amount of arguments. Expected 5, but got %d instead.", len(args))
	}
	idInt, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Sprintf("1st argument should be integer.")
	}
	id := uint(idInt)
	if !storage.Exist(id) {
		return fmt.Sprintf("Review #%d does not exist.", id)
	}
	rating, err := strconv.Atoi(args[4])
	if err != nil {
		return "5th argument should be integer."
	}
	r, err := storage.MakeReview(id, args[1], args[2], args[3], uint8(rating))
	if err != nil {
		return err.Error()
	}
	err = storage.Update(r)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Review #%d was successfully updated.", r.GetId())
}

func deleteHandler(s string) string {
	if len(s) == 0 {
		return "No arguments were given. See /help for details."
	}
	args, err := parseArgs(s)
	if err != nil {
		return "Invalid arguments. Make sure you don't use quotes in a quoted part."
	}
	if len(args) != 1 {
		return fmt.Sprintf("Invalid amount of arguments. Expected 1, but got %d instead.", len(args))
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Sprintf("Argument should be integer.")
	}
	err = storage.Delete(uint(id))
	if err != nil {
		return fmt.Sprintf("No review to delete.")
	}
	return fmt.Sprintf("Review #%d was successfully deleted.", uint(id))
}

func RegisterHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpHandler)
	c.RegisterHandler(listCmd, listHandler)
	c.RegisterHandler(addCmd, addHandler)
	c.RegisterHandler(updateCmd, updateHandler)
	c.RegisterHandler(deleteCmd, deleteHandler)
}
