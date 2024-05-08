package config

import (
	"github.com/OpenIMSDK/tools/errs"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"strings"
)

func LoadConfig(path string, envPrefix string, config any) error {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return errs.Wrap(err, "failed to read config file", "path", path, "envPrefix", envPrefix)
	}

	if err := v.Unmarshal(config, func(config *mapstructure.DecoderConfig) {
		config.TagName = "mapstructure"
	}); err != nil {
		return errs.Wrap(err, "failed to unmarshal config", "path", path, "envPrefix", envPrefix)
	}
	return nil
}
