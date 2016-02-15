package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"net"
	"strings"
	"os"

	"coterie"
)

var address string
var batch_size int

func init() {
	flag.StringVar(&address, "address", "", "Address of coteried application")
	flag.IntVar(&batch_size, "batch_size", 50, "Size of record batchs to send on LOAD")
}

func main() {
	flag.Parse()
	fmt.Printf("welcome to coterie-cli!\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		fields := strings.Fields(cmd)
		switch fields[0] {
		case "exit", "EXIT":
			os.Exit(1)
		case "help", "HELO":
			fmt.Printf("\texit : exit the application\n\tload <filename> : load the given csv filename\n")
		case "load", "LOAD":
			fmt.Printf("loading file '%s'\n", fields[1])
			file, err := os.Open(fields[1])
			if err != nil {
				panic(err)
			}

			reader := csv.NewReader(file)
			keys, err := reader.Read()
			if err != nil {
				panic(err)
			}

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

				if len(records) % batch_size == 0 {
					coterieMsg := new(coterie.CoterieMsg)
					coterieMsg.Type = coterie.CoterieMsg_RECORD_BATCH
					coterieMsg.RecordBatchMsg = &coterie.RecordBatchMsg { records }

					conn, err := net.Dial("tcp", address)
					if err != nil {
						panic(err)
					}

					coterie.WriteCoterieMsg(coterieMsg, conn)
					records = nil
				}
			}

			if len(records) != 0 {
				coterieMsg := new(coterie.CoterieMsg)
				coterieMsg.Type = coterie.CoterieMsg_RECORD_BATCH
				coterieMsg.RecordBatchMsg = &coterie.RecordBatchMsg { records }

				conn, err := net.Dial("tcp", address)
				if err != nil {
					panic(err)
				}

				coterie.WriteCoterieMsg(coterieMsg, conn)
			}
		default:
		}

		fmt.Printf("> ")
	}
}
