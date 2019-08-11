package handler

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func populateTestDB() error {
	testData := map[string]string{
		"rocks":  "https://en.wikipedia.org/wiki/List_of_rock_types",
		"trees":  "http://bhort.bh.cornell.edu/tree/list.htm",
		"clouds": "https://scied.ucar.edu/webweather/clouds/cloud-types",
		"cycle":  "https://www.cycle-route.com/",
		"bike":   "https://www.cyclist.co.uk/tags/bike-maintenance",
		"read":   "https://www.goodreads.com/",
	}
	db, err := bolt.Open("testdb", 0666, nil)
	if err != nil {
		return fmt.Errorf("Failed setting up testdb, error: [%v]\n", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("URLMappings"))
		if err != nil {
			return fmt.Errorf("Failed creating a bucket URLMappings, error: [%v]\n", err)
		}
		return nil
	})
	for key, value := range testData {
		err = db.Update(func(tx *bolt.Tx) error {
			err := tx.Bucket([]byte("URLMappings")).Put([]byte(key), []byte(value))
			if err != nil {
				return fmt.Errorf("Failed inserting key %s, value %s into bucket URLMappings, error: [%v]\n", key, value, err)
			}
			return nil
		})
	}
	return nil
}
