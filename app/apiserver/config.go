package apiserver

import "github.com/OldGattsu/go-rest-api/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":5555",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
