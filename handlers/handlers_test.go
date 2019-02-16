package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapHandler(t *testing.T) {
	testMux := func() *http.ServeMux {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintln(w, "Testing")
		})
		return mux
	}
	testdata := map[string]string{
		"/clouds": "https://scied.ucar.edu/webweather/clouds/cloud-types",
	}
	//test cases
	req := httptest.NewRequest("GET", "/clouds", nil)

	rr := httptest.NewRecorder()
	mapHandler := MapHandler(testdata, testMux())
	mapHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("MapHandler returned wrong status, expected %v got %v", http.StatusSeeOther, status)
	}
	if location := rr.Result().Header["Location"][0]; location != testdata["/clouds"] {
		t.Errorf("Expected location %v (redirect) got %v", testdata["/clouds"], rr.Result().Header["Location"])
	}

	req = httptest.NewRequest("GET", "/random", nil)

	rr = httptest.NewRecorder()
	mapHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("MapHandler returned wrong status, expected %v got %v", http.StatusOK, status)
	}
	body, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "Testing\n" {
		t.Errorf("Expected body %q got %q", "Testing\n", string(body))
	}

}

func TestYamlHandler(t *testing.T) {
	testdata := `
- "path": "/rocks"
  "url": "https://en.wikipedia.org/wiki/List_of_rock_types"
`
	testMux := func() *http.ServeMux {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintln(w, "Testing")
		})
		return mux
	}
	req := httptest.NewRequest("GET", "/rocks", nil)
	rr := httptest.NewRecorder()
	yamlHandler, err := YAMLHandler([]byte(testdata), testMux())
	if err != nil {
		t.Fatal(err)
	}
	yamlHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Status expected %v got %v", http.StatusSeeOther, status)
	}
	if location := rr.Result().Header["Location"][0]; location != "https://en.wikipedia.org/wiki/List_of_rock_types" {
		t.Errorf("Expected location %v (redirect) got %v", "https://en.wikipedia.org/wiki/List_of_rock_types", rr.Result().Header["Location"])
	}
}

func TestJSONHandler(t *testing.T) {
	testdata := `
    [{"path": "/gl", "url": "https://docs.gitlab.com/ee/README.html"}]
    `
	testMux := func() *http.ServeMux {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintln(w, "Testing")
		})
		return mux
	}

	req := httptest.NewRequest("GET", "/gl", nil)
	rr := httptest.NewRecorder()
	jsonHandler, err := JSONHandler([]byte(testdata), testMux())
	if err != nil {
		t.Fatal(err)
	}
	jsonHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Status expected %v got %v", http.StatusSeeOther, status)
	}
	if location := rr.Result().Header["Location"][0]; location != "https://docs.gitlab.com/ee/README.html" {
		t.Errorf("Expected location %v (redirect) got %v", "https://docs.gitlab.com/ee/README.html", rr.Result().Header["Location"])
	}
}