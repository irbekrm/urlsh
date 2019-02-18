package handler

import (
	"fmt"
	"github/irbekrm/urlsh/datastore"
	"log"
	"net/http"
)

type IHandler interface {
	Handle(fallback http.Handler) http.HandlerFunc
}

type Handler struct {
	ds datastore.DataStore
}

func NewHandler(dsType, dsPath string) (IHandler, error) {
	ds, err := datastore.NewDataStore(dsType, dsPath)
	if err != nil {
		return nil, fmt.Errorf("Failed initializing datastore [%s], error: [%v]", dsPath, err)
	}
	return &Handler{
		ds: ds,
	}, nil
}

func (h *Handler) Handle(fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		realURL, found, err := h.ds.Search(r.URL.Path)
		if err != nil {
			log.Fatalf("Failed searching path [%s], error: [%v]", r.URL.Path, err)
		}
		if found {
			http.Redirect(w, r, realURL, http.StatusSeeOther)
		}
		fallback.ServeHTTP(w, r)
	})
}