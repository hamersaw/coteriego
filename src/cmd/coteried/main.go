package main

import (
	"flag"
	"fmt"
	"net"

	"coterie"
	"dht"

	"github.com/BurntSushi/toml"
)

type CoterieConfig struct {
	Address     string
	DHT         dht.DHTConfig
}

var configurationFile string

func init() {
	flag.StringVar(&configurationFile, "config_file", "", "Toml configuration file")
}

func main() {
	flag.Parse()
	var config CoterieConfig

	if _, err := toml.DecodeFile(configurationFile, &config); err != nil {
		panic(err)
	}

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
	defer conn.Close()
	coterieMsg, err := coterie.ReadCoterieMsg(conn)
	if err != nil {
		panic(err)
	}

	switch(coterieMsg.Type) {
	case coterie.CoterieMsg_INSERT_RECORD:
		break
	default:
		fmt.Printf("TODO - handle coterie messsage type: %v\n", coterieMsg.Type)
		break
	}
}
