package join

import (
	"github.com/antonpriyma/RSCC/internal/app/handlers"
	"github.com/antonpriyma/RSCC/pkg/log"
	"github.com/pkg/errors"
	tele "gopkg.in/telebot.v3"
)

type Config struct {
	// welcome message
	WelcomeMessage string `mapstructure:"welcome_message"`
}

type joinHandler struct {
	Config
	Logger log.Logger
}

func New(c Config, logger log.Logger) handlers.Handler {
	return joinHandler{
		Config: c,
		Logger: logger,
	}
}

func (h joinHandler) Handle(c tele.Context) error {
	err := h.sendWelcomeMessage(c)

	return errors.Wrap(err, "failed to send welcome message")
}

func (h joinHandler) sendWelcomeMessage(c tele.Context) error {
	if h.WelcomeMessage != "" {
		text := h.WelcomeMessage

		err := c.Reply(text)
		if err != nil {
			h.Logger.Errorf("failed to send welcome message: %s", err)
			return errors.Wrap(err, "failed to reply")
		}
	}

	return nil
}
