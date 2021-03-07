package config

import (
	env "github.com/caarlos0/env/v6"
)

func init() {
	if err := env.Parse(&LINE); err != nil {
		panic(err)
	}
}

var LINE line

type line struct {
	ChannelSecret string `env:"LINE_CHANNEL_SECRET"`
	AccessToken   string `env:"LINE_ACCESS_TOKEN"`
}
