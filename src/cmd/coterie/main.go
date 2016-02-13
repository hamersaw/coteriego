package main

import (
	"flag"
	"fmt"

	"dht"

	"github.com/BurntSushi/toml"
)

var configurationFile string

func init() {
	flag.StringVar(&configurationFile, "config_file", "", "Toml configuration file")
}

func main() {
	flag.Parse()
	var config tomlConfig

	//parse configuration file
	if _, err := toml.DecodeFile(configurationFile, &config); err != nil {
		panic(err)
	}

	//start dht service
	dhtService := dht.NewDHTService(config.Tokens, config.Application.Address, config.DHT.Address, config.Seeds)
	dhtService.Start()

	//TMP print out tomlConfig
	fmt.Printf("tokens: %v\n", config.Tokens)
	fmt.Printf("application address: '%s'\n", config.Application.Address)
	fmt.Printf("dht address: '%s'\n", config.DHT.Address)
	for _, address := range config.Seeds {
		fmt.Printf("seed address: '%s'\n", address)
	}
}

type tomlConfig struct {
	Tokens      []uint64
	Application address
	DHT         address
	Seeds       []string
}

type address struct {
	Address string
}
