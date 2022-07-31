package help

import (
	"fmt"

	commandPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command"
)

func New(extendedMap map[string]string) commandPkg.Interface {
	if extendedMap == nil {
		extendedMap = map[string]string{}
	}

	return &command{
		extended: extendedMap,
	}
}

type command struct {
	extended map[string]string
}

func (c *command) Name() string {
	return "help"
}

func (c *command) Description() string {
	return "show this message."
}

func (c *command) Process(_ string) string {
	result := fmt.Sprintf("/%s - %s\n", c.Name(), c.Description())
	for cmd, description := range c.extended {
		result += fmt.Sprintf("/%s - %s\n", cmd, description)
	}
	return result
}
