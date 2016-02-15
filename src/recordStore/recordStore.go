package recordStore

import (
	"errors"
	"fmt"
	"hash/crc64"
	"sync"
)

type RecordStore struct {
	records     map[uint64]map[string]string
	recordsLock sync.RWMutex
	entries     map[string]map[string][]uint64
	entriesLock sync.RWMutex
}

func NewRecordStore() *RecordStore {
	return &RecordStore {
		make(map[uint64]map[string]string),
		sync.RWMutex{},
		make(map[string]map[string][]uint64),
		sync.RWMutex{},
	}
}

func (d *RecordStore) InsertRecord(token uint64, record map[string]string) error {
	//fmt.Printf("inserting record '%d': '%v'\n", token, record)
	d.recordsLock.Lock()
	if _, ok := d.records[token]; ok {
		d.recordsLock.Unlock()
		return errors.New(fmt.Sprintf("Token %d already exists in records map", token))
	}

	d.records[token] = record
	d.recordsLock.Unlock()
	return nil
}

func (d *RecordStore) InsertEntry(token uint64, key string, value string) error {
	//fmt.Printf("inserting entity '%s': '%s'\n", key, value)
	d.entriesLock.Lock()
	m, ok := d.entries[key]
	if !ok {
		m = make(map[string][]uint64)
		d.entries[key] = m
	}

	s, ok := m[value]
	if !ok {
		s = []uint64{}
		m[value] = s
	}

	for _, t := range s {
		if t == token {
			d.entriesLock.Unlock()
			return errors.New(fmt.Sprintf("Token %d already exists in entry map", token))
		}
	}

	s = append(s, token)
	d.entriesLock.Unlock()
	return nil
}

func (d *RecordStore) GetMatchingRecordKeys(field string, value string, match func(string, string) bool) ([]uint64, error) {
	m, ok := d.entries[field]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unable to find field '%s' in entry map", field))
	}

	tokens := []uint64{}
	for k, v := range m {
		if match(k, value) {
			tokens = append(tokens, v...)
		}
	}

	return tokens, nil
}

func ComputeRecordToken(record map[string]string) uint64 {
	crcTable := crc64.MakeTable(crc64.ECMA)
	return crc64.Checksum([]byte(fmt.Sprintf("%v", record)), crcTable)
}

func ComputeEntryToken(key string, value string) uint64 {
	crcTable := crc64.MakeTable(crc64.ECMA)
	return crc64.Checksum([]byte(fmt.Sprintf("%s:%s", key, value)), crcTable)
}
