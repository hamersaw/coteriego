package dht

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"

	"message"

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

type Seed struct {
	Address string
}

func NewDHTService(tokens []uint64, address string, applicationAddress string, seeds []Seed) *DHTService {
	lookupTable := make(map[uint64]string)
	for _, token := range tokens {
		lookupTable[token] = applicationAddress
	}

	return &DHTService {
		tokens: tokens,
		address: address,
		applicationAddress: applicationAddress,
		seeds: seeds,
		lookupTable: lookupTable,
		peerTable: make(map[string][]uint64),
	}
}

func (d *DHTService) Start() error {
	fmt.Println("starting dht service")

	//send node join messages to all seeds
	dhtMsg := new(message.DHTMsg)
	dhtMsg.Type = message.DHTMsg_JOIN
	dhtMsg.JoinMsg = &message.JoinMsg {d.tokens, d.applicationAddress}

	for _, seed := range d.seeds {
		conn, err := net.Dial("tcp", seed.Address)
		if err != nil {
			panic(err)
		}

		writeDHTMsg(dhtMsg, conn)

		fmt.Printf("TODO - recv join message reply")
	}

	//start listening for connections
	listener, err := net.Listen("tcp", d.address)
	if err != nil {
		return err
	}

	for {
		fmt.Printf("listening for connections on %s\n", d.address)
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
	case message.DHTMsg_JOIN:
		fmt.Printf("GOT JOIN MSG - TODO implement everything\n")
		break;
	default:
		fmt.Printf("dht messsage type: %v\n", dhtMsg.Type)
		break;
	}
}

func (d *DHTService) Lookup(token uint64) (string, error) {
	fmt.Printf("TODO - handle token lookup")

	return "", errors.New("unimplemented")
}

func (d *DHTService) AddPeer(address string, tokens []uint64) error {
	if _, ok := d.peerTable[address]; ok {
		return errors.New(fmt.Sprintf("Unable to insert, address %v already exists in peer table", address))
	}
	d.peerTable[address] = tokens
	return nil
}

func (d *DHTService) RemovePeer(address string) error {
	if _, ok := d.peerTable[address]; ok {
		return errors.New(fmt.Sprintf("Unable to delete, address %v doesn't exist in peer table", address))
	}
	delete(d.peerTable, address)
	return nil
}

func (d *DHTService) AddToken(token uint64, address string) error {
	if _, ok := d.lookupTable[token]; ok {
		return errors.New(fmt.Sprintf("Unable to insert, token %d already exists in lookup table", token))
	}
	d.lookupTable[token] = address
	return nil
}

func (d *DHTService) RemoveToken(token uint64) error {
	if _, ok := d.lookupTable[token]; ok {
		return errors.New(fmt.Sprintf("Unable to delete, token %d doesn't exist in lookup table", token))
	}
	delete(d.lookupTable, token)
	return nil
}

func writeDHTMsg(dhtMsg *message.DHTMsg, conn net.Conn) error {
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

func readDHTMsg(conn net.Conn) (*message.DHTMsg, error) {
	buf := make([]byte, 4096)
	_, err := conn.Read(buf[:8])
	if err != nil {
		return nil, err
	}

	length, bytesRead := binary.Uvarint(buf[:8])
	if bytesRead < 0 {
		return nil, errors.New("Unable to parse length of dht message protobuf")
	}

	_, err = conn.Read(buf[:length])
	if err != nil {
		return nil, err
	}

	dhtMsg := new(message.DHTMsg)
	err = proto.Unmarshal(buf[:length], dhtMsg)
	if err != nil {
		return nil, err
	}

	return dhtMsg, nil
}
