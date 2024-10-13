package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Simple POST request connectivity test to google.com
func TestPostSuccess(t *testing.T) {
	reqBody := []byte(`{"url": "https://google.com"}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Received status code: %v ", status)
	}

}

// Tries an empty payload in the POST
func TestPostEmpty(t *testing.T) {
	reqBody := []byte(`{"url": ""}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Received status code: %v ", status)
	}
}
