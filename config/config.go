package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

var configPath string

type Config struct {
	Application Application `mapstructure:"application"`
	OpenApi     OpenApi     `mapstructure:"openApi"`
}

type Application struct {
	Port             string `mapstructure:"port"`
	OpenApiMaxWorker string `mapstructure:"openApiMaxWorker"`
}

type OpenApi struct {
	ApiKey string `mapstructure:"apiKey"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}

	if configPath == "" {
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}
