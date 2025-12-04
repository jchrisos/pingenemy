package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jchrisos/pingenemy/internal/http"
	"github.com/jchrisos/pingenemy/internal/url"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := &url.UrlRequest{
		Name:               "google",
		URL:                "https://google.com",
		HttpMethod:         "GET",
		ExpectedStatusCode: 200,
		IntervalSeconds:    5,
	}

	go executeUrl(ctx, url)

	select {}
}

func executeUrl(ctx context.Context, urlReq *url.UrlRequest) {
	interval := time.Duration(urlReq.IntervalSeconds) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	exec := &http.HttpExecutor{}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Job is done")
			return
		case <-ticker.C:
			func() {
				success, err := exec.Execute(urlReq)
				if err != nil {
					panic("Failed to calling url")
				}

				if success {
					fmt.Println("OK")
				}
			}()
		}
	}
}
