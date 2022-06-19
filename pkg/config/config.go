package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func LoadConfig(path string, config interface{}) error {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read config")
	}

	err = viper.Unmarshal(&config)
	return nil
}
