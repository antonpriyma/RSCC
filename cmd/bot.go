package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/antonpriyma/RSCC/pkg/log"
)

func main() {
	logger := log.Default()

	// TODO: add config

	// init bot
	bot, err := tgbotapi.NewBotAPI("1652329832:AAHSL2CrNQt19fCGHzJcle_kjktyNofpmVs")
	must(logger, err, "failed to init telegram bot")

	bot.Debug = true

	u := tgbotapi.NewUpdate(0) // TODO: to cfg
	u.Timeout = 60             // TODO: to cfg

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		logger.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

}

func must(logger log.Logger, err error, msg string) {
	if err != nil {
		logger.WithError(err).Fatal(msg)
	}
}
