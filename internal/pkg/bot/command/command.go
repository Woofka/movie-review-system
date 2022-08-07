package command

import (
	"context"
	"encoding/csv"
	"strings"
)

type Interface interface {
	Name() string
	Description() string
	Process(ctx context.Context, args string) string
}

func ParseArguments(args string) ([]string, error) {
	reader := csv.NewReader(strings.NewReader(args))
	reader.Comma = ' '
	res, err := reader.Read()
	if err != nil {
		return nil, err
	}
	return res, nil
}
