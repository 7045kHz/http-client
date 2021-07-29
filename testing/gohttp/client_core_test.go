package gohttp

import (
	"net/http"
	"testing"
)

// White box testing
// One test per return
func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	requestHeaders.Set("Content-Type", "application/json;charset=utf-8")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("expected 3 headers")
	}
	if finalHeaders.Get("Content-Type") != "application/json;charset=utf-8" {
		t.Error("custom header for Content-Type not set")
	}

}
