package morning_greeter

import (
	"github.com/antonpriyma/RSCC/internal/app/tasks"
	"github.com/antonpriyma/RSCC/pkg/log"
	"github.com/pkg/errors"
	tele "gopkg.in/telebot.v3"
	"time"
)

type Config struct {
	GreetingText     string   `mapstructure:"greeting_text"`
	GreetingStickers []string `mapstructure:"greeting_stickers"`
	ChatID           int64    `mapstructure:"chat_id"`
}

type task struct {
	Config

	Logger log.Logger
	Bot    tele.Bot
}

func New(cfg Config, logger log.Logger, bot tele.Bot) tasks.Task {
	return task{
		Config: cfg,
		Logger: logger,
		Bot:    bot,
	}
}

func (t task) Run() error {
	// send text message
	err := t.sendTextMessage()
	if err != nil {
		return errors.Wrap(err, "failed to send text message")
	}

	// send cringe sticker
	err = t.sendCringeSticker()
	if err != nil {
		return errors.Wrap(err, "failed to send sticker")
	}

	return nil
}

func (t task) sendTextMessage() error {
	_, err := t.Bot.Send(
		&tele.Chat{ID: t.ChatID},
		t.GreetingText,
	)

	return errors.Wrap(err, "failed to send message")
}

func (t task) sendCringeSticker() error {
	fileID := t.GreetingStickers[time.Now().Weekday()]

	_, err := t.Bot.Send(
		&tele.Chat{ID: t.ChatID},
		&tele.Sticker{
			File: tele.File{
				FileID: fileID,
			},
		},
	)

	return errors.Wrap(err, "failed to send sticker")
}
