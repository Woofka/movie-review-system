package main

import (
	"log"

	"gitlab.ozon.dev/Woofka/movie-review-system/internal/config"
	botPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot"
	cmdAddPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/add"
	cmdDeletePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/delete"
	cmdGetPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/get"
	cmdHelpPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/help"
	cmdListPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/list"
	cmdUpdatePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/update"
	reviewPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review"
)

func main() {
	var review reviewPkg.Interface
	{
		review = reviewPkg.New()
	}

	var bot botPkg.Interface
	{
		bot, err := botPkg.New(config.ApiToken, false)
		if err != nil {
			log.Fatal(err)
		}

		commandAdd := cmdAddPkg.New(review)
		bot.RegisterHandler(commandAdd)

		commandGet := cmdGetPkg.New(review)
		bot.RegisterHandler(commandGet)

		commandUpdate := cmdUpdatePkg.New(review)
		bot.RegisterHandler(commandUpdate)

		commandDelete := cmdDeletePkg.New(review)
		bot.RegisterHandler(commandDelete)

		commandList := cmdListPkg.New(review)
		bot.RegisterHandler(commandList)

		// TODO: update descriptions
		commandHelp := cmdHelpPkg.New(map[string]string{
			commandAdd.Name():    commandAdd.Description(),
			commandGet.Name():    commandGet.Description(),
			commandUpdate.Name(): commandUpdate.Description(),
			commandDelete.Name(): commandDelete.Description(),
			commandList.Name():   commandList.Description(),
		})
		bot.RegisterHandler(commandHelp)
	}

	go runBot(bot)
	err := runGRPCServer(review)
	if err != nil {
		log.Fatal(err)
	}
}

func runBot(bot botPkg.Interface) {
	err := bot.Run()
	if err != nil {
		log.Fatal(err)
	}
}
