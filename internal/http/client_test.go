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
		TimeoutMillis:      5000,
	}

	client := NewHttpClient()

	result, err := client.Call(t.Context(), url)
	if err != nil {
		t.Error(err)
	}

	if !result.Success {
		t.Errorf("Test failed. result.Success: %v, expected: %v", result, true)
	}
}
