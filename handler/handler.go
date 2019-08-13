package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/irbekrm/urlsh/datastore"
)

type IHandler interface {
	Handle() http.HandlerFunc
}

type Handler struct {
	ds       datastore.DataStore
	fallback http.Handler
}

func New(dsType string, fallback http.Handler) (IHandler, error) {
	ds, err := datastore.New(dsType)
	if err != nil {
		return nil, fmt.Errorf("Failed initializing datastore [%s], error: [%v]", dsType, err)
	}
	return &Handler{
		ds:       ds,
		fallback: fallback,
	}, nil
}

func (h *Handler) Handle() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		realURL, found, err := h.ds.Get(r.URL.Path)
		if err != nil {
			log.Fatalf("Failed searching path [%s], error: [%v]", r.URL.Path, err)
		}
		if found {
			http.Redirect(w, r, realURL, http.StatusSeeOther)
		}
		h.fallback.ServeHTTP(w, r)
	})
}
