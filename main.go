package main

import (
	"fmt"
	"net/http"

        "github/irbekrm/urlsh/handlers"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/clouds":  "https://en.wikipedia.org/wiki/List_of_cloud_types",
		"/medium": "https://medium.com",
		"/godocs": "https://golang.org/doc/",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
//	yaml := `
//- path: /urlshort
//  url: https://github.com/gophercises/urlshort
//- path: /urlshort-final
//  url: https://github.com/gophercises/urlshort/tree/solution
//`
	//yamlHandler, err := handlers.YAMLHandler([]byte(yaml), mapHandler)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
