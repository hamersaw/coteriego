package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
	"os"

	"coterie"

	"github.com/codegangsta/cli"
)

var ipAddress, port string

func main() {
	app := cli.NewApp()
	app.Name = "coterie-cli"
	app.Usage = "provide an easy user interface to the coteried program"

	app.Flags= []cli.Flag {
		cli.StringFlag {
			Name:        "ipAddress",
			Value:       "127.0.0.1",
			Usage:       "IP Address of the coteried host",
			Destination: &ipAddress,
		},
		cli.StringFlag {
			Name:        "port",
			Value:       "15605",
			Usage:       "Port of the coteried host",
			Destination: &port,
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "load",
			Aliases: []string{"LOAD"},
			Usage: "load data from a csv file",
			Action: loadFile,
		},
		{
			Name: "query",
			Aliases: []string{"q"},
			Usage: "submit a sql like query",
			Action: query,
		},
	}

	app.Run(os.Args)
}

func loadFile(c *cli.Context) {
	fmt.Printf("loading file '%s'\n", c.Args()[0])
	file, err := os.Open(c.Args()[0])
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	keys, err := reader.Read()
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ipAddress, port))
	if err != nil {
		panic(err)
	}

	startTime := time.Now()
	recordCount := 0
	records := []*coterie.Record{}
	for {
		values, err := reader.Read()
		if err != nil {
			break
		}

		record := make(map[string]string)
		for i, key := range keys {
			record[key] = values[i]
		}

		records = append(records, &coterie.Record{ record })
		recordCount++
		if len(records) == 50 {
			if err = sendRecordBatchMsg(records, conn); err != nil {
				panic(err)
			}
			records = nil
		}
	}

	if len(records) != 0 {
		if err = sendRecordBatchMsg(records, conn); err != nil {
			panic(err)
		}
	}

	fmt.Printf("loaded %d records in %v", recordCount, time.Since(startTime))

	coterieMsg := new(coterie.CoterieMsg)
	coterieMsg.Type = coterie.CoterieMsg_CLOSE_CONNECTION
	coterieMsg.CloseConnectionMsg = &coterie.CloseConnectionMsg { "finished writes" }
	if err = coterie.WriteCoterieMsg(coterieMsg, conn); err != nil {
		panic(err)
	}
	conn.Close()
}

func sendRecordBatchMsg(records []*coterie.Record, conn net.Conn) error {
	coterieMsg := new(coterie.CoterieMsg)
	coterieMsg.Type = coterie.CoterieMsg_RECORD_BATCH
	coterieMsg.RecordBatchMsg = &coterie.RecordBatchMsg { records }

	if err := coterie.WriteCoterieMsg(coterieMsg, conn); err != nil {
		return err
	}

	rtnMsg, err := coterie.ReadCoterieMsg(conn)
	if err != nil {
		return err
	}

	if rtnMsg.Type != coterie.CoterieMsg_RESULT {
		return errors.New("Expecting result message")
	}

	return nil
}

func query(c *cli.Context) {
	fmt.Println("TODO - query '", strings.Join(c.Args(), " "), "'")
}
