package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultEndpoint(t *testing.T) {
	mux := CreateServerMux(DefaultVersion)
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Error making request: %s", err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading request body: %s", err.Error())
	}

	expectedData, err := json.MarshalIndent(&Response{Message: "Hello world!"}, "", "  ")
	if err != nil {
		t.Errorf("Error marshalling expected response: %s", err.Error())
	}

	dataStr := string(data)
	expStr := string(expectedData)
	if dataStr != expStr {
		t.Errorf("Data mismatched. Expected: %s, got: %s", expStr, dataStr)
	}
}
