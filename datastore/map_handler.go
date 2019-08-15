package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type MapHandler struct {
	data map[string]string
}

var (
	jsonFilePath = "data/urls.json"
	yamlFilePath = "data/urls.yaml"
)

func NewMapHandler(dstype string) (*MapHandler, error) {
	var data map[string]string
	switch dstype {
	case "json":
		j := getEnvVar("FILE_PATH", jsonFilePath)
		contents, err := fileReader(j)
		if err != nil {
			return nil, err
		}
		data, err = jsonMapper(contents)
		if err != nil {
			return nil, err
		}
	case "yaml":
		y := getEnvVar("FILE_PATH", yamlFilePath)
		contents, err := fileReader(y)
		if err != nil {
			return nil, err
		}
		data, err = yamlMapper(contents)
		if err != nil {
			return nil, err
		}
	}
	return &MapHandler{data: data}, nil

}

func (m *MapHandler) Get(url string) (string, bool, error) {
	return "", false, nil
}
func fileReader(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func jsonMapper(j []byte) (map[string]string, error) {
	tempSlice := make([]map[string]string, 10)
	pathToUrls := make(map[string]string)
	err := json.Unmarshal(j, &tempSlice)
	if err != nil {
		return nil, err
	}
	for _, v := range tempSlice {
		pathToUrls[v["path"]] = v["url"]
	}
	return pathToUrls, nil
}

func yamlMapper(y []byte) (map[string]string, error) {
	tempSlice := make([]map[string]string, 10)
	pathsToUrls := make(map[string]string)
	err := yaml.Unmarshal(y, &tempSlice)
	if err != nil {
		return nil, err
	}
	for _, v := range tempSlice {
		pathsToUrls[v["path"]] = v["url"]
	}
	return pathsToUrls, nil
}
