package main

import (
	"context"

	"github.com/jchrisos/pingenemy/internal/http"
	"github.com/jchrisos/pingenemy/internal/job"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urls := []http.UrlRequest{
		{
			Name:               "google",
			URL:                "https://google.com",
			HttpMethod:         "GET",
			ExpectedStatusCode: 200,
			IntervalSeconds:    2,
		},
		{
			Name:               "uol",
			URL:                "https://uol.com.br",
			HttpMethod:         "GET",
			ExpectedStatusCode: 200,
			IntervalSeconds:    5,
		},
	}

	job := &job.Job{}

	for _, url := range urls {
		go job.Execute(ctx, &url)
	}

	select {}
}
