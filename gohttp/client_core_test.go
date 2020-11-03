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

	if finalHeaders.Get("X-Request-Id") != "ASD-123" {
		t.Error("invalid request id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}

	if finalHeaders.Get("User-Agent") != "http-client" {
		t.Error("invalid user agent received")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("no body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body received")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/xml", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as xml")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("invalid xml body received")
		}
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body received")
		}
	})
}
