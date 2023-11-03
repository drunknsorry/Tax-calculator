package apiconsumer

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJson(t *testing.T) {

	// Start a mock http server and sample Json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"tax_brackets":[{"max": 10000, "min": 0, "rate": 10.0}]}`))
	}))
	defer server.Close()

	// Call GetJson
	var response TaxBracketResults
	err := GetJson(server.URL, &response)

	// Check for errors
	if err != nil {
		t.Errorf("GetJson error: %v", err)
	}

	// Verify the structure of the response
	if len(response.TaxBrackets) != 1 {
		t.Error("Expected 1 tax bracket, but got a different number")
	}

}

// More detail testing needs to be implemented to cover other functions and with larger varying data sets
