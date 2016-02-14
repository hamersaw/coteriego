package main

import (
	"flag"
	"net"

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
	go dhtService.Start()

	listener, err := net.Listen("tcp", config.Address)
	if err != nil {
		panic(err)
	}

	for {
		if conn, err := listener.Accept(); err == nil {
			go handleConn(conn)
		} else {
			continue
		}
	}
}

func handleConn(conn net.Conn) {
	
}

type CoterieConfig struct {
	Address     string
	DHT         dht.DHTConfig
}
