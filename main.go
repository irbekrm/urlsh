package main

import (
	"fmt"
	"net/http"
	"github/irbekrm/urlsh/handlers"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/ny": "https://www.newyorker.com",
		"/rl": "https://www.rigaslaiks.lv",
		"/gr": "https://www.goodreads.com",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	yaml := `
- "path": "/clouds"
  "url": "https://en.wikipedia.org/wiki/List_of_cloud_types"
- "path": "/medium"
  "url": "https://medium.com"
- "path": "/godocs"
  "url": "https://golang.org/doc/"
- "path": "/yml"
  "url": "https://en.wikipedia.org/wiki/YAML"
`
	yamlHandler, err := handlers.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
