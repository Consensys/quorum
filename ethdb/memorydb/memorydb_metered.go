package memorydb

import (
	"errors"
	"strconv"
	"sync"
)

type MeteredDatabase struct {
	Database

	readCount   int
	writeCount  int
	deleteCount int

	countLock sync.RWMutex
}

func NewMetered() *MeteredDatabase {
	return &MeteredDatabase{
		Database: Database{
			db: make(map[string][]byte),
		},
		readCount:  0,
		writeCount: 0,
	}
}

func (db *MeteredDatabase) Get(key []byte) ([]byte, error) {
	db.countLock.RLock()
	defer db.countLock.RUnlock()
	db.readCount++
	return db.Database.Get(key)
}

func (db *MeteredDatabase) Put(key []byte, value []byte) error {
	db.countLock.RLock()
	defer db.countLock.RUnlock()

	db.writeCount++
	return db.Database.Put(key, value)
}

func (db *MeteredDatabase) Delete(key []byte) error {
	db.countLock.RLock()
	defer db.countLock.RUnlock()

	db.deleteCount++
	return db.Database.Delete(key)
}

func (db *MeteredDatabase) Stat(property string) (string, error) {
	switch property {
	case "readCount":
		return strconv.Itoa(db.readCount), nil
	case "writeCount":
		return strconv.Itoa(db.writeCount), nil
	case "deleteCount":
		return strconv.Itoa(db.deleteCount), nil
	default:
		return "", errors.New("unknown property")
	}
}
