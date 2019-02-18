package datastore

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type BoltDataStore struct {
	path string
}

//TODO: validate database here
func NewBoltDataStore(path string) (*BoltDataStore, error) {
	return &BoltDataStore{
		path: path,
	}, nil
}

func (bd BoltDataStore) Search(url string) (string, bool, error) {
	var tempPath []byte
	var path string
	var found bool
	db, err := bolt.Open(bd.path, 0666, nil)
	if err != nil {
		return "", false, fmt.Errorf("Failed opening database [%s], error: [%v]", bd.path, err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		tempPath = tx.Bucket([]byte("URLMappings")).Get([]byte(url))
		return nil
	})
	if err != nil {
		return path, found, fmt.Errorf("Failed viewing database [%s], error: [%v]", bd.path, err)
	}
	if tempPath != nil {
		path = string(tempPath)
		found = true
	}
	return path, found, nil
}
