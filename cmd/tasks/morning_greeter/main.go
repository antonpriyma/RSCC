package main

import (
	"flag"
	"github.com/antonpriyma/RSCC/internal/app/tasks/morning_greeter"
	"github.com/antonpriyma/RSCC/internal/pkg/config"
	"github.com/antonpriyma/RSCC/pkg/log"
	"github.com/antonpriyma/RSCC/pkg/utils"
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

	// task
	task := morning_greeter.New(cfg.GreeterConfig, logger, *bot)
	err = task.Run()
	if err != nil {
		logger.WithError(err).Fatal("failed to run task")
	}

	return
}
