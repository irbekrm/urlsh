package datastore

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type DataStore interface {
	Get(string) (string, bool, error)
}

func New(dsType string) (DataStore, error) {
	switch dsType {
	case "redis":
		fmt.Println("Using Redis datastore..")
		r, err := NewRedis()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to initialize a Redis datastore")
		}
		return r, nil
	case "default":
		fmt.Println("Could not determine datastore type")
		return nil, nil
	}
	return nil, nil
}

func getEnvVar(n, v string) string {
	if val := os.Getenv(n); val != "" {
		return val
	}
	return v
}
