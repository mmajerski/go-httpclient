package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "http-client")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ASD-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers")
	}
}
