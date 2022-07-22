package commander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type CmdHandler func(string) string

type Commander struct {
	bot    *tgbotapi.BotAPI
	routes map[string]CmdHandler
}

func Init(apiToken string, debug bool) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	bot.Debug = debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	commander := Commander{
		bot:    bot,
		routes: make(map[string]CmdHandler),
	}

	return &commander, nil
}

func (c *Commander) RegisterHandler(cmd string, handler CmdHandler) {
	c.routes[cmd] = handler
}

func (c *Commander) Run() error {
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
			cmd := update.Message.Command()
			if h, ok := c.routes[cmd]; ok {
				msg.Text = h(update.Message.CommandArguments())
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
