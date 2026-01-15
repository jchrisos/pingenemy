package http

import (
	"testing"
)

const (
	timeoutMillis int64 = 1000
)

func TestCall(t *testing.T) {
	url := &UrlRequest{
		Name:               "google",
		URL:                "https://google.com",
		HttpMethod:         "GET",
		ExpectedStatusCode: 200,
		IntervalSeconds:    1,
	}

	client := NewHttpClientTest(int64(timeoutMillis))

	result, _, _ := client.Call(t.Context(), url)

	if !result {
		t.Errorf("Test failed. result: %v, expected: %v", result, true)
	}
}
