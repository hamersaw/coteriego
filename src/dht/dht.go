package dht

import (
	"errors"
	"fmt"
	"net"
)

type DHTService struct {
	tokens []uint64
	address string
	applicationAddress string
	seedAddresses []string
	lookupTable map[uint64]string
	peerTable map[string][]uint64
}

func NewDHTService(tokens []uint64, address string, applicationAddress string, seedAddresses []string) *DHTService {
	lookupTable := make(map[uint64]string)
	for _, token := range tokens {
		lookupTable[token] = applicationAddress
	}

	return &DHTService {
		tokens: tokens,
		address: address,
		applicationAddress: applicationAddress,
		seedAddresses: seedAddresses,
		lookupTable: lookupTable,
		peerTable: make(map[string][]uint64),
	}
}

func (d *DHTService) Start() error {
	fmt.Println("starting dht service")

	//send node join messages to all seeds
	for _, ipAddress := range d.seedAddresses {
		fmt.Printf("TODO - send join message %v", ipAddress)
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

	fmt.Printf("TODO - handle client connections")
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
