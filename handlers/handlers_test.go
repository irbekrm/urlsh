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