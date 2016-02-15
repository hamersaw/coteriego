package dht

import (
	"fmt"
	"testing"
)

func TestDHTService(t *testing.T) {
	t.Log("Creating DHT Service object")
	dhtService := DHTService { []uint64{}, "", "", []Seed{}, make(map[uint64]string), make(map[string][]uint64) }
	dhtService.AddToken(0, "local_addr")
	dhtService.AddToken(5, "remote_addr")
	dhtService.AddToken(10, "local_addr")
	dhtService.AddToken(15, "remote_addr")
	dhtService.AddToken(20, "local_addr")

	lookup := make(map[uint64]string)
	lookup[1] = "remote_addr"
	lookup[5] = "remote_addr"
	lookup[11] = "remote_addr"
	lookup[25] = "local_addr"
	for token, address := range lookup {
		rtnAddress, err := dhtService.Lookup(token)
		if err != nil {
			t.Fail()
		}
		if rtnAddress != address {
			t.Log(fmt.Sprintf("Lookup for token %d expected '%s' but got '%s'", token, address, rtnAddress))
			t.Fail()
		}
	}
}
