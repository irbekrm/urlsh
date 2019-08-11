package handler

import (
	"net/http"
)

type BoltDBHandler struct {
	dbPath string
}

func (b *BoltDBHandler) Handle(fallback http.Handler) http.HandlerFunc {
	return nil
}

func dbOpener(path string) error {
	return nil
}
