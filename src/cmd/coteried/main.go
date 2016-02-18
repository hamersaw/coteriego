package main

import (
	"flag"
	"fmt"
	"net"

	"coterie"
	"dht"
	"recordStore"

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

	recordStore := recordStore.NewRecordStore()
	for {
		if conn, err := listener.Accept(); err == nil {
			go handleConn(conn, recordStore, dhtService)
		} else {
			continue
		}
	}
}

func handleConn(conn net.Conn, recStore *recordStore.RecordStore, dhtService *dht.DHTService) {
	defer conn.Close()
	conns := make(map[string]*net.TCPConn)

	for {
		coterieMsg, err := coterie.ReadCoterieMsg(conn)
		if err != nil {
			panic(err)
		}

		switch(coterieMsg.Type) {
		case coterie.CoterieMsg_CLOSE_CONNECTION:
		case coterie.CoterieMsg_INSERT_ENTRY:
			insertEntryMsg := coterieMsg.GetInsertEntryMsg()
			if err := recStore.InsertEntry(insertEntryMsg.Token, insertEntryMsg.Key, insertEntryMsg.Value); err != nil {
				panic(err)
			}
			continue
		case coterie.CoterieMsg_INSERT_RECORD:
			insertRecordMsg := coterieMsg.GetInsertRecordMsg()
			if err := recStore.InsertRecord(insertRecordMsg.Token, insertRecordMsg.GetRecord().GetEntries()); err != nil {
				panic(err)
			}
			continue
		case coterie.CoterieMsg_INSERT_RECORDS:
			insertRecordsMsg := coterieMsg.GetInsertRecordsMsg()
			for _, record := range insertRecordsMsg.GetRecords() {
				recordKey := recordStore.ComputeRecordToken(record.GetEntries())
				recordAddress, err := dhtService.Lookup(recordKey)
				if err != nil {
					panic(err)
				}

				insertRecordMsg := new(coterie.CoterieMsg)
				insertRecordMsg.Type = coterie.CoterieMsg_INSERT_RECORD
				insertRecordMsg.InsertRecordMsg = &coterie.InsertRecordMsg { recordKey, record }

				recordConn := getConn(recordAddress, conns)
				err = coterie.WriteCoterieMsg(insertRecordMsg, recordConn)
				if err != nil {
					panic(err)
				}

				for key, value := range record.GetEntries() {
					entryKey := recordStore.ComputeEntryToken(key, value)
					entryAddress, err := dhtService.Lookup(entryKey)
					if err != nil {
						panic(err)
					}

					insertEntryMsg := new(coterie.CoterieMsg)
					insertEntryMsg.Type = coterie.CoterieMsg_INSERT_ENTRY
					insertEntryMsg.InsertEntryMsg = &coterie.InsertEntryMsg { entryKey, key, value }

					entryConn := getConn(entryAddress, conns)
					err = coterie.WriteCoterieMsg(insertEntryMsg, entryConn)
					if err != nil {
						panic(err)
					}
				}
			}

			resultMsg := new(coterie.CoterieMsg)
			resultMsg.Type = coterie.CoterieMsg_RESULT
			resultMsg.ResultMsg = &coterie.ResultMsg { true, "" }

			if err = coterie.WriteCoterieMsg(resultMsg, conn); err != nil {
				panic(err)
			}
			continue
		case coterie.CoterieMsg_QUERY:
			queryMsg := coterieMsg.GetQueryMsg()

			for _, filter := range queryMsg.Filters {
				fmt.Printf("Filter on %s ~%s(%v) %s\n", filter.FieldName, filter.Type, filter.Arguments, filter.Value)
			}

			fmt.Println("TODO - execute query")

			queryResultMsg := new(coterie.QueryResultMsg)
			queryResultMsg.Records = []*coterie.Record{}

			coterieMsg := new(coterie.CoterieMsg)
			coterieMsg.Type = coterie.CoterieMsg_QUERY_RESULT
			coterieMsg.QueryResultMsg = &coterie.QueryResultMsg { []*coterie.Record{} }

			if err = coterie.WriteCoterieMsg(coterieMsg, conn); err != nil {
				panic(err)
			}
		default:
			fmt.Printf("TODO - handle coterie messsage type: %v\n", coterieMsg.Type)
		}

		break
	}

	closeConnectionMsg := new(coterie.CoterieMsg)
	closeConnectionMsg.Type = coterie.CoterieMsg_CLOSE_CONNECTION
	closeConnectionMsg.CloseConnectionMsg = &coterie.CloseConnectionMsg { "finished writes" }
	for _, conn := range conns {
		coterie.WriteCoterieMsg(closeConnectionMsg, conn)
		conn.Close()
	}
}

func getConn(address string, conns map[string]*net.TCPConn) *net.TCPConn {
	if _, ok := conns[address]; ok {
		conn, _ := conns[address]
		return conn
	} else {
		tcpAddress, err := net.ResolveTCPAddr("tcp", address)
		if err != nil {
			panic(err)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddress)
		if err != nil {
			panic(err)
		}

		conns[address] = conn
		return conn
	}
}
