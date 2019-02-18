package handler

import (
	"net/http"

	bolt "go.etcd.io/bbolt"
)

type BoltDBHandler struct {
	dbPath string
}

func (b *BoltDBHandler) Handle(fallback http.Handler) http.HandlerFunc {

}

func dbOpener(path string) error {
	db, err := bolt.Open
}
