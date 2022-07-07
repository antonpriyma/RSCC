package text

import (
	"github.com/antonpriyma/RSCC/internal/app/handlers"
	"github.com/antonpriyma/RSCC/pkg/log"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	tele "gopkg.in/telebot.v3"
	"math/rand"
	"strings"
	"time"
)

type Config struct {
	// trigger_word:reply
	Replies map[string][]string `mapstructure:"replies"`

	// limit per minute
	Limit float64 `mapstructure:"limit"`
}

type textHandler struct {
	Config
	Logger  log.Logger
	Limiter *rate.Limiter
}

func New(c Config, logger log.Logger, limiter *rate.Limiter) handlers.Handler {
	return textHandler{
		Config:  c,
		Logger:  logger,
		Limiter: limiter,
	}
}

func (h textHandler) Handle(ctx tele.Context) error {
	return h.sendTextReply(ctx)
}

func (h textHandler) sendTextReply(ctx tele.Context) error {
	rand.Seed(time.Now().UnixNano())

	for text, reply := range h.Replies {
		if strings.Contains(strings.ToLower(ctx.Text()), text) && len(reply) > 0 {
			if h.Limiter.Allow() {
				err := ctx.Reply(reply[rand.Intn(len(reply))])
				return errors.Wrap(err, "failed to send reply")
			}
		}
	}

	return nil
}
