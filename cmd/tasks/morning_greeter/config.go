package main

import (
	"github.com/antonpriyma/RSCC/internal/app/tasks/morning_greeter"
	"github.com/antonpriyma/RSCC/internal/pkg/config"
	"github.com/pkg/errors"
)

type Config struct {
	config.BotConfig `mapstructure:"bot"`
	GreeterConfig    morning_greeter.Config `mapstructure:"morning_greeter"`
}

// TODO: normal validating in one place
func (c Config) Validate() error {
	if len(c.GreeterConfig.GreetingStickers) != 7 {
		return errors.New("wrong amount of stickers")
	}

	return nil
}
