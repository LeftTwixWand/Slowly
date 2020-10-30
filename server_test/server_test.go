package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/lefttwixwand/slowly/server"
)

// Test method for server.Slol function
func Test_Slow(t *testing.T) {
	rout := mux.NewRouter()                   // Create a route
	rout.HandleFunc("/api/slow", server.Slow) // Set server.Slow mathod as route handler

	ts := httptest.NewServer(rout)
	defer ts.Close() // Auto dispose after test ending

	res, err := http.Get(ts.URL + "/api/slow?timeout=500") // Normal request

	if err != nil { // If there are any errors
		t.Error(res, err)
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != 200 { // Check status code for successful result
		t.Error(res, err)
		t.Errorf("Expected status code = 200 %s", err.Error())
	}

	res, err = http.Get(ts.URL + "/api/slow?timeout=6000") // Wrong request

	if err != nil {
		t.Error(res, err)
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != 400 {
		t.Error(res, err)
		t.Errorf("Expected status code = 200 %s", err.Error())
	}
}
