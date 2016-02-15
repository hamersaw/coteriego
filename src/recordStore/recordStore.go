package recordStore

import (
	"errors"
	"fmt"
)

type RecordStore struct {
	records map[uint64]map[string]string
	entries  map[string]map[string][]uint64
}

func NewRecordStore() *RecordStore {
	return &RecordStore {
		make(map[uint64]map[string]string),
		make(map[string]map[string][]uint64),
	}
}

func (d *RecordStore) InsertRecord(token uint64, record map[string]string) error {
	if _, ok := d.records[token]; ok {
		return errors.New(fmt.Sprintf("Token %d already exists in records map", token))
	}

	d.records[token] = record
	return nil
}

func (d *RecordStore) InsertEntry(token uint64, field string, value string) error {
	m, ok := d.entries[field]
	if !ok {
		m = make(map[string][]uint64)
		d.entries[field] = m
	}

	s, ok := m[value]
	if !ok {
		s = []uint64{}
		m[value] = s
	}

	for _, t := range s {
		if t == token {
			return errors.New(fmt.Sprintf("Token %d already exists in entry map", token))
		}
	}

	s = append(s, token)
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
	return 0
}

func ComputeEntryToken(key string, value string) uint64 {
	return 0
}
