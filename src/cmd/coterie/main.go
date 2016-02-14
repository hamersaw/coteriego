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
	var config CoterieConfig

	//parse configuration file
	if _, err := toml.DecodeFile(configurationFile, &config); err != nil {
		panic(err)
	}

	//start dht service
	dhtService := dht.NewDHTService(config.DHT.Tokens, config.DHT.Address, config.Address, config.DHT.Seeds)
	dhtService.Start()

	//TMP print out tomlConfig
	fmt.Printf("tokens: %v\n", config.DHT.Tokens)
	fmt.Printf("application address: '%s'\n", config.Address)
	fmt.Printf("dht address: '%s'\n", config.DHT.Address)
	for _, address := range config.DHT.Seeds {
		fmt.Printf("seed address: '%s'\n", address)
	}
}

type CoterieConfig struct {
	Address     string
	DHT         dht.DHTConfig
}
