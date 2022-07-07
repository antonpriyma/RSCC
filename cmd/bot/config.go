package main

import (
	"github.com/antonpriyma/RSCC/internal/app/handlers/join"
	"github.com/antonpriyma/RSCC/internal/app/handlers/message/fuji"
	"github.com/antonpriyma/RSCC/internal/app/handlers/message/text"
	"github.com/antonpriyma/RSCC/internal/pkg/config"
)

type Config struct {
	config.BotConfig `mapstructure:"bot"`
	HandlersConfig   `mapstructure:"handlers"`
}

func (c Config) Validate() error {
	return nil
}

type HandlersConfig struct {
	Fuji fuji.Config `mapstructure:"fuji"`
	Text text.Config `mapstructure:"text"`
	Join join.Config `mapstructure:"join"`
}
