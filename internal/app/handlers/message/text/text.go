package text

import (
	"github.com/antonpriyma/RSCC/internal/app/handlers/message"
	"github.com/antonpriyma/RSCC/pkg/log"
	"github.com/pkg/errors"
	tele "gopkg.in/telebot.v3"
	"math/rand"
	"time"
)

type Config struct {
	// trigger_word:reply
	Replies map[string][]string `mapstructure:"replies"`
}

type textHandler struct {
	Config
	Logger log.Logger
}

func New(c Config, logger log.Logger) message.Handler {
	return textHandler{
		Config: c,
		Logger: logger,
	}
}

func (h textHandler) Handle(ctx tele.Context) error {
	return h.sendTextReply(ctx)
}

func (h textHandler) sendTextReply(ctx tele.Context) error {
	rand.Seed(time.Now().UnixNano())

	if replies := h.Replies[ctx.Text()]; len(replies) > 0 {
		err := ctx.Send(replies[rand.Intn(len(replies))])
		return errors.Wrap(err, "failed to send reply")
	}

	return nil
}
