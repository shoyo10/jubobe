package config

import (
	"jubobe/pkg/echorouter"
	"jubobe/pkg/postgres"
	"jubobe/pkg/zerolog"

	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var _config *Config

// Get get global config
func Get() *Config {
	return _config
}

// Set global config
func Set(config *Config) {
	_config = config
}

// Config ...
type Config struct {
	fx.Out

	Log      *zerolog.Config    `yaml:"log"`
	HTTP     *echorouter.Config `yaml:"http"`
	Postgres *postgres.Config   `yaml:"postgres"`
}

// New read config from file
func New() (*Config, error) {
	viper.AutomaticEnv()

	configPath := viper.GetString("CONFIG_DIR")
	if configPath == "" {
		configPath = "./"
	}

	configName := viper.GetString("CONFIG_NAME")
	if configName == "" {
		configName = "app"
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Msgf("Error reading config file, %s", err)
		return nil, err
	}

	var config Config
	err := viper.Unmarshal(&config, func(cfg *mapstructure.DecoderConfig) {
		cfg.TagName = "yaml"
	})
	if err != nil {
		log.Error().Msgf("Unable to decode into struct, %v", err)
		return nil, err
	}

	Set(&config)

	return _config, nil
}
