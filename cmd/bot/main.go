package main

import (
	"log"
	"movie-review-system/config"
	"movie-review-system/internal/commander"
	"movie-review-system/internal/handlers"
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
