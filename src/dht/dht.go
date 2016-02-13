package dht

import (
	"errors"
	"fmt"
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
	return &DHTService {
		tokens: tokens,
		address: address,
		applicationAddress: applicationAddress,
		seedAddresses: seedAddresses,
		lookupTable: make(map[uint64]string),
		peerTable: make(map[string][]uint64),
	}
}

func (d *DHTService) Start() error {
	fmt.Println("starting dht service")

	//send join messages to seed addresses
	for _, ipAddress := range d.seedAddresses {
		fmt.Printf("sending join message %v", ipAddress)
	}

	return nil
}

func (d *DHTService) Lookup(token uint64) (string, error) {
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
