package main

import (
	"fmt"
	"github/irbekrm/urlsh/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/ny": "https://www.newyorker.com",
		"/rl": "https://www.rigaslaiks.lv",
		"/gr": "https://www.goodreads.com",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory, error: [%v]\n", err)
	}

	// Uncomment to read urls from /data/urls.yaml
    //filePath := dir + "/data/urls.yaml"
    //handler, err := DataFromFile(filePath, "yaml", mapHandler)

    //Uncomment to read urls from /data/urls.json
    filePath := dir + "/data/urls.json"
    handler, err := DataFromFile(filePath, "json", mapHandler)


  if err != nil {
  	log.Fatalf("Failed parsing urls, file: [%v], error: [%v]\n", filePath, err)
  }
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func DataFromFile(path, fileType string, fallback http.Handler) (http.HandlerFunc, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	switch fileType {
	case "json":
		return handlers.JSONHandler(fileContents, fallback)
	case "yaml":
		return handlers.YAMLHandler(fileContents, fallback)
	default:
		return nil, fmt.Errorf("Cannot parse file type [%v]\n", fileType)
	}
	return nil, nil
	}
