package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Improve test to check for different methods, input values, status codes
func TestHome(t *testing.T) {
	// Request home
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create mock response recorder
	rr := httptest.NewRecorder()

	// start a test server and set handler
	handler := http.HandlerFunc(home)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code: %v, got %v", http.StatusOK, status)
	}

	// Clean up any log folders created
	os.RemoveAll("log")
}
