package main

import (
	"flag"
	"github.com/antonpriyma/RSCC/internal/app/handlers/message"
	"github.com/antonpriyma/RSCC/internal/app/handlers/message/fuji"
	"github.com/antonpriyma/RSCC/internal/app/handlers/message/text"
	"github.com/antonpriyma/RSCC/internal/pkg/config"
	"github.com/antonpriyma/RSCC/pkg/log"
	"github.com/antonpriyma/RSCC/pkg/utils"
	"github.com/pkg/errors"
	tele "gopkg.in/telebot.v3"
	"time"
)

var cfgPath = flag.String("config", "./cfg/bot.yaml", "path to config")

func main() {
	flag.Parse()

	// init logger
	logger := log.Default()

	// read config
	cfg := Config{}
	err := config.LoadConfig(*cfgPath, &cfg)
	utils.Must(logger, err, "failed to load config")
	logger.Infof("read config: %v", cfg)

	// init bot
	pref := tele.Settings{
		Token:  cfg.BotKey,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	utils.Must(logger, err, "failed to init telegram bot")

	// message handlers
	fujiMessageHandler := fuji.New(cfg.Fuji, logger)

	textHandler := text.New(cfg.Text, logger)

	messageHandlers := []message.Handler{
		fujiMessageHandler,
		textHandler,
	}

	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		logger.
			WithField("chat_id", ctx.Chat().ID).
			Printf("[%s] %ds", ctx.Message().OriginalSenderName, ctx.Message().Text)

		for _, handler := range messageHandlers {
			err := handler.Handle(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to handle message")
			}
		}

		return nil
	})

	bot.Start()
}
