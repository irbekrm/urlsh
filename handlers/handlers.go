package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if realURL, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, realURL, http.StatusSeeOther)
		}
		fallback.ServeHTTP(w, r)
	})
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	tempSlice := make([]map[string]string, 10)
	pathsToUrls := make(map[string]string)
    err := yaml.Unmarshal(yml, &tempSlice)
    if err != nil {
    	return nil, err
	}
	for _, v := range tempSlice {
		pathsToUrls[v["path"]] = v["url"]
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	tempSlice := make([]map[string]string, 10)
	pathToUrls := make(map[string]string)
	err := json.Unmarshal(jsn, &tempSlice)
	if err != nil {
		return nil, err
	}
	for _, v := range tempSlice {
		pathToUrls[v["path"]] = v["url"]
	}
	return MapHandler(pathToUrls, fallback), nil
}