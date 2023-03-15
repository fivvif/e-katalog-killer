package main

import (
	"e_katalog_killer/cmd/internal/app/apiserver"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "cmd/internal/configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
