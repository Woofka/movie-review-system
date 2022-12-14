package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	commandPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/bot/command"
)

type Interface interface {
	Run(ctx context.Context) error
	RegisterHandler(cmd commandPkg.Interface)
}

func New(apiToken string, debug bool) (Interface, error) {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	bot.Debug = debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &commander{
		bot:    bot,
		routes: make(map[string]commandPkg.Interface),
	}, nil
}

type commander struct {
	bot    *tgbotapi.BotAPI
	routes map[string]commandPkg.Interface
}

// RegisterHandler - not thread-safe
func (c *commander) RegisterHandler(cmd commandPkg.Interface) {
	c.routes[cmd.Name()] = cmd
}

func (c *commander) Run(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			cmdName := update.Message.Command()
			if cmd, ok := c.routes[cmdName]; ok {
				msg.Text = cmd.Process(ctx, update.Message.CommandArguments())
				msg.ParseMode = "markdown"
			} else {
				msg.Text = "Unknown command. Use /help to see available commands."
			}
		} else {
			msg.Text = "I can work only with commands. Use /help to see available commands."
		}

		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
