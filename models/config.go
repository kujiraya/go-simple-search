package models

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	Title  string `toml:"title"`
	Server serverConfig
	DB     dbConfig
}

type serverConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type dbConfig struct {
	Name string `toml:"name"`
	User string `toml:"user"`
}

func getConfig() config {
	var config config
	_, err := toml.DecodeFile("./model.toml", &config)
	if err != nil {
		panic(err)
	}
	return config
}
