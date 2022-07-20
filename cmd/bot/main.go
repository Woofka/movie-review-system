package main

import (
	"log"
	"movie-review-system/config"
	"movie-review-system/internal/commander"
)

func main() {
	cmd, err := commander.Init(config.ApiToken, false)
	if err != nil {
		log.Panic(err)
	}

	err = cmd.Run()
	if err != nil {
		log.Panic(err)
	}
}
