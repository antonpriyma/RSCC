package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config interface {
	Validate() error
}

type BotConfig struct {
	BotKey string `mapstructure:"bot_key"`
	Debug  bool   `mapstructure:"debug"`
}

func LoadConfig(path string, config Config) error {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read config")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}

	err = config.Validate()
	return errors.Wrap(err, "failed to validate config")
}
