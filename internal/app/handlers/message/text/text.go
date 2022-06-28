package text

import (
	"github.com/antonpriyma/RSCC/internal/app/handlers/message"
	"github.com/antonpriyma/RSCC/pkg/log"
	"github.com/pkg/errors"
	tele "gopkg.in/telebot.v3"
	"math/rand"
	"strings"
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

	for text, reply := range h.Replies {
		if strings.Contains(strings.ToLower(ctx.Text()), text) && len(reply) > 0 {
			err := ctx.Reply(reply[rand.Intn(len(reply))])
			return errors.Wrap(err, "failed to send reply")
		}
	}

	return nil
}
