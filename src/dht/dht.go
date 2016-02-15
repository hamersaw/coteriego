package dht

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
)

type DHTService struct {
	tokens             []uint64
	address            string
	applicationAddress string
	seeds              []Seed
	lookupTable        map[uint64]string
	peerTable          map[string][]uint64
}

type DHTConfig struct {
	Tokens         []uint64
	Address        string
	HeartbeatDelay uint
	Seeds          []Seed
}

type Seed struct {
	Address string
}

func NewDHTService(tokens []uint64, address string, applicationAddress string, seeds []Seed) *DHTService {
	dhtService := DHTService { tokens, address, applicationAddress, seeds, make(map[uint64]string), make(map[string][]uint64) }
	for _, token := range tokens {
		dhtService.AddToken(token, applicationAddress)
	}

	return &dhtService
}

func (d *DHTService) Start() error {
	for _, seed := range d.seeds {
		d.AddPeer(seed.Address, nil)
	}

	go func(){
		for executions := 0; ; executions++ {
			dhtMsg := new(DHTMsg)
			dhtMsg.Type = DHTMsg_HEARTBEAT
			dhtMsg.HeartbeatMsg = &HeartbeatMsg { d.tokens, d.address, d.applicationAddress, executions % 10 == 0 }
			for address, _ := range d.peerTable {
				conn, err := net.Dial("tcp", address)
				if err != nil {
					tokens, _ := d.peerTable[address]
					for _, token := range tokens {
						_ = d.RemoveToken(token)
					}
					_ = d.RemovePeer(address)
					continue
				}

				if err = writeDHTMsg(dhtMsg, conn); err != nil {
					panic(err)
				}

				rtnMsg, err := readDHTMsg(conn)
				if err != nil {
					panic(err)
				}

				switch rtnMsg.Type {
				case DHTMsg_RESULT:
				case DHTMsg_LOOKUP_TABLE_DUMP:
					for token, address := range rtnMsg.GetLookupTableDumpMsg().GetLookupTable() {
						_ = d.AddToken(token, address)
					}
				default:
					panic(errors.New("Expecting RESULT or LOOKUP_TABLE_DUMP type"))
				}
			}

			time.Sleep(time.Duration(time.Second * 3))
		}
	}()

	listener, err := net.Listen("tcp", d.address)
	if err != nil {
		return err
	}

	for {
		if conn, err := listener.Accept(); err == nil {
			go d.handleConn(conn)
		} else {
			continue
		}
	}

	return nil
}

func (d *DHTService) handleConn(conn net.Conn) {
	defer conn.Close()
	dhtMsg, err := readDHTMsg(conn)
	if err != nil {
		panic(err)
	}

	switch(dhtMsg.Type) {
	case DHTMsg_HEARTBEAT:
		heartbeatMsg := dhtMsg.GetHeartbeatMsg()
		rtnMsg := new(DHTMsg)
		if heartbeatMsg.RequestTableDump {
			rtnMsg.Type = DHTMsg_LOOKUP_TABLE_DUMP
			rtnMsg.LookupTableDumpMsg = &LookupTableDumpMsg { d.lookupTable }
		} else {
			rtnMsg.Type = DHTMsg_RESULT
			rtnMsg.ResultMsg = &ResultMsg { true, "" }
		}

		if err = writeDHTMsg(rtnMsg, conn); err != nil {
			panic(err)
		}

		d.AddPeer(heartbeatMsg.Address, heartbeatMsg.Tokens)
		for _, token := range heartbeatMsg.Tokens {
			_ = d.AddToken(token, heartbeatMsg.ApplicationAddress)
		}

	default:
		fmt.Printf("TODO - handle dht messsage type: %v\n", dhtMsg.Type)
	}
}

func (d *DHTService) Lookup(token uint64) (string, error) {
	tokens := []uint64{}
	for t, _ := range d.lookupTable {
		tokens = append(tokens, t)
	}

	for i:=1; i < len(tokens); i++ {
		for j:=0; j < len(tokens)-i; j++ {
			if (tokens[j] > tokens[j+1]) {
				tokens[j], tokens[j+1] = tokens[j+1], tokens[j]
			}
		}
	}

	var address string
	if token < tokens[0] {
		address, _ := d.lookupTable[tokens[len(tokens)-1]]
		return address, nil
	}

	var tokenKey uint64
	for _, t := range tokens {
		if token <= t {
			continue
		} else {
			tokenKey = t
			break
		}
	}

	address, _ = d.lookupTable[tokenKey]
	return address, nil
}

func (d *DHTService) AddPeer(address string, tokens []uint64) error {
	//fmt.Printf("Adding peer with address '%s' and tokens '%v'\n", address, tokens)
	if value, ok := d.peerTable[address]; ok && value != nil {
		return errors.New(fmt.Sprintf("Unable to insert, address %v already exists in peer table", address))
	}
	d.peerTable[address] = tokens
	return nil
}

func (d *DHTService) RemovePeer(address string) error {
	//fmt.Printf("Removing peer with address '%s'\n", address)
	if _, ok := d.peerTable[address]; !ok {
		return errors.New(fmt.Sprintf("Unable to delete, address %v doesn't exist in peer table", address))
	}
	delete(d.peerTable, address)
	return nil
}

func (d *DHTService) AddToken(token uint64, address string) error {
	//fmt.Printf("Adding token '%d' for peer '%s'\n", token, address)
	if _, ok := d.lookupTable[token]; ok {
		return errors.New(fmt.Sprintf("Unable to insert, token %d already exists in lookup table", token))
	}
	d.lookupTable[token] = address
	return nil
}

func (d *DHTService) RemoveToken(token uint64) error {
	//fmt.Printf("Removing token '%d'\n", token)
	if _, ok := d.lookupTable[token]; !ok {
		return errors.New(fmt.Sprintf("Unable to delete, token %d doesn't exist in lookup table", token))
	}
	delete(d.lookupTable, token)
	return nil
}

func writeDHTMsg(dhtMsg *DHTMsg, conn net.Conn) error {
	bytes, err := proto.Marshal(dhtMsg)
	if err != nil {
		return err
	}
	lengthBytes := make([]byte, 8)
	binary.PutUvarint(lengthBytes, uint64(len(bytes)))

	if _, err := conn.Write(lengthBytes); err != nil {
		return err
	}

	if _, err := conn.Write(bytes); err != nil {
		return err
	}

	return nil
}

func readDHTMsg(conn net.Conn) (*DHTMsg, error) {
	buf := make([]byte, 4096)
	_, err := conn.Read(buf[:8])
	if err != nil {
		return nil, err
	}

	length, bytesRead := binary.Uvarint(buf[:8])
	if bytesRead < 0 {
		return nil, errors.New("Unable to parse length of dht protobuf")
	}

	_, err = conn.Read(buf[:length])
	if err != nil {
		return nil, err
	}

	dhtMsg := new(DHTMsg)
	err = proto.Unmarshal(buf[:length], dhtMsg)
	if err != nil {
		return nil, err
	}

	return dhtMsg, nil
}
