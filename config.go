package main

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	BotToken string `yaml:"bot_token"`
	OwnerId  int64  `yaml:"owner_id"`
}

func parseConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse config file")
	}

	return config, nil
}
