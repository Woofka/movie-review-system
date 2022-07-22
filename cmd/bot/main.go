package main

import (
	"gitlab.ozon.dev/Woofka/movie-review-system/config"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/commander"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/handlers"
	"log"
)

func main() {
	cmd, err := commander.Init(config.ApiToken, false)
	if err != nil {
		log.Panic(err)
	}
	handlers.RegisterHandlers(cmd)

	err = cmd.Run()
	if err != nil {
		log.Panic(err)
	}
}
