package main

import (
	"flag"
	"github.com/antonpriyma/RSCC/internal/pkg/handlers/message"
	"github.com/antonpriyma/RSCC/internal/pkg/handlers/message/fuji"
	"github.com/antonpriyma/RSCC/pkg/config"
	"github.com/antonpriyma/RSCC/pkg/log"
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
	must(logger, err, "failed to load config")
	logger.Infof("read config: %v", cfg)

	// init bot
	pref := tele.Settings{
		Token:  cfg.BotKey,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	must(logger, err, "failed to init telegram bot")

	// message handlers
	fujiMessageHandler := fuji.New(cfg.Fuji, logger)

	messageHandlers := []message.Handler{
		fujiMessageHandler,
	}

	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		logger.Printf("[%s] %s", ctx.Message().OriginalSenderName, ctx.Message().Text)

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

func must(logger log.Logger, err error, msg string) {
	if err != nil {
		logger.WithError(err).Fatal(msg)
	}
}
