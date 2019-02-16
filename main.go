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
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory, error: [%v]\n", err)
	}

	// Uncomment to read urls from /data/urls.yaml
    //filePath := dir + "/data/urls.yaml"
    //handler, err := DataFromFile(filePath, "yaml", mux)

    //Uncomment to read urls from /data/urls.json
    filePath := dir + "/data/urls.json"
    handler, err := DataFromFile(filePath, "json", mux, fileOpener, fileReader)


	if err != nil {
		log.Fatalf("Failed parsing urls, file: [%v], error: [%v]\n", filePath, err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func DataFromFile(path, fileType string, fallback http.Handler, fileOpener func(string)(*os.File, error), fileReader func(*os.File)([]byte, error)) (http.HandlerFunc, error) {
	if !(fileType == "json" || fileType == "yaml") {
		return nil, fmt.Errorf("Cannot parse file type [%v]\n", fileType)
	}
	file, err := fileOpener(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := fileReader(file)
	if err != nil {
		return nil, err
	}

	switch fileType {
	case "json":
		return handlers.JSONHandler(fileContents, fallback)
	case "yaml":
		return handlers.YAMLHandler(fileContents, fallback)
	}
	return nil, nil
}

func fileOpener(path string) (*os.File, error) {
	return os.Open(path)
}
func fileReader(file *os.File) ([]byte, error) {
	return ioutil.ReadAll(file)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
