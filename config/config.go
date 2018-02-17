package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"strconv"
)

var config *Config

type Config struct {
	Application APIConfig `toml:"application"`
	Database    DBConfig  `toml:"database"`
}

type APIConfig struct {
	Port int
}

func GetAPIPort() int {
	return config.Application.Port
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
}

func GetDBConnInfo() string {
	return config.Database.User + ":" + config.Database.Password + "@" + "tcp(" + config.Database.Host + ":" + strconv.Itoa(config.Database.Port) + ")/" + config.Database.Schema
}

func init() {
	if config == nil {
		config = &Config{}
		_, err := toml.DecodeFile("config.toml", &config)
		if err != nil {
			fmt.Errorf("err: %v", err)
		}
	}
}
