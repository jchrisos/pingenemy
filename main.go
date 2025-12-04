package main

import (
	"fmt"

	"github.com/jchrisos/pingenemy/internal/http"
	"github.com/jchrisos/pingenemy/internal/url"
)

func main() {
	url := &url.UrlRequest{
		Name:               "google",
		URL:                "https://google.com",
		HttpMethod:         "GET",
		ExpectedStatusCode: 200,
	}

	exec := &http.HttpExecutor{}

	success, err := exec.Execute(url)
	if err != nil {
		panic("Failed to calling url")
	}

	if success {
		fmt.Println("OK")
	}
}
