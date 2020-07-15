package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

var Cfg Config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &Cfg); err != nil {
		log.Fatal(err)
	}
}
