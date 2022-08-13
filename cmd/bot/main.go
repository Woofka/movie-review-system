package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/config"
	botPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot"
	cmdAddPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/add"
	cmdDeletePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/delete"
	cmdGetPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/get"
	cmdHelpPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/help"
	cmdListPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/list"
	cmdUpdatePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command/update"
	reviewPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache/postgres"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	psqlConn := "host=localhost port=5432 user=user password=password dbname=movie_review sslmode=disable"
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("ping database error", err)
	}

	pgConfig := pool.Config()
	pgConfig.MaxConnIdleTime = time.Minute
	pgConfig.MaxConnLifetime = time.Hour
	pgConfig.MinConns = 2
	pgConfig.MaxConns = 4

	var review reviewPkg.Interface
	review = reviewPkg.New(postgres.New(pool))
	// review = reviewPkg.New(local.New())

	var bot botPkg.Interface
	bot, err = botPkg.New(config.ApiToken, false)
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

	commandHelp := cmdHelpPkg.New(map[string]string{
		commandAdd.Name():    commandAdd.Description(),
		commandGet.Name():    commandGet.Description(),
		commandUpdate.Name(): commandUpdate.Description(),
		commandDelete.Name(): commandDelete.Description(),
		commandList.Name():   commandList.Description(),
	})
	bot.RegisterHandler(commandHelp)

	go runBot(ctx, bot)
	err = runGRPCServer(review)
	if err != nil {
		log.Fatal(err)
	}
}

func runBot(ctx context.Context, bot botPkg.Interface) {
	err := bot.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
