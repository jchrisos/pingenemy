package main

import (
	"context"
	"log"

	"github.com/jchrisos/pingenemy/internal/http"
	"github.com/jchrisos/pingenemy/internal/job"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urls, err := http.NewUrl().RetriveUrls()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	job := &job.Job{}

	for _, url := range urls {
		go job.Execute(ctx, &url)
	}

	select {}
}
