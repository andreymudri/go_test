package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/andreymudri/go_test/src/controller"
)

func TestHelloWorld(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Helloworld))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code error, got %v want %v", resp.StatusCode, http.StatusOK)
	}

}

func TestListSave(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controller.SaveListHandler))
	body := map[string][]int{
		"List1": {1, 2, 3},
		"List2": {4, 5, 6},
	}
	encoded := new(bytes.Buffer)
	err := json.NewEncoder(encoded).Encode(body)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(server.URL, "application/json", encoded)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code error, got %v want %v", resp.StatusCode, http.StatusMethodNotAllowed)
	}
}

func TestListMerge(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controller.SaveListHandler))
	body := map[string][]int{
		"List1": {1, 2, 3},
		"List2": {4, 5, 6},
	}
	encoded := new(bytes.Buffer)
	err := json.NewEncoder(encoded).Encode(body)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(server.URL, "application/json", encoded)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code error, got %v want %v", resp.StatusCode, http.StatusMethodNotAllowed)
	}
	serverGet := httptest.NewServer(http.HandlerFunc(controller.MergeHandler))
	respGet, errGet := http.Get(serverGet.URL)
	if errGet != nil {
		t.Error(errGet)
	}
	if respGet.StatusCode != http.StatusOK {
		t.Errorf("Status code error, got %v want %v", respGet.StatusCode, http.StatusOK)
	}
}
