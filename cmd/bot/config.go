package main

import "github.com/antonpriyma/RSCC/internal/pkg/handlers/message/fuji"

type Config struct {
	BotConfig      `mapstructure:"bot"`
	HandlersConfig `mapstructure:"handlers"`
}

type BotConfig struct {
	BotKey string `mapstructure:"bot_key"`
	Debug  bool   `mapstructure:"debug"`
}

type HandlersConfig struct {
	Fuji fuji.Config `mapstructure:"fuji"`
}
