package http

import (
	"testing"

	"github.com/jchrisos/pingenemy/internal/url"
)

func TestExecute(t *testing.T) {

	url := &url.UrlRequest{
		Name:               "google",
		URL:                "https://google.com",
		HttpMethod:         "GET",
		ExpectedStatusCode: 200,
	}

	exec := &HttpExecutor{}

	result, _ := exec.Execute(url)

	if !result {
		t.Errorf("Test failed. result: %v, expected: %v", result, true)
	}

}
