package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jchrisos/pingenemy/internal/http"
	"github.com/jchrisos/pingenemy/internal/job"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	urls, err := http.NewUrl().RetriveUrls()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Go(func() {
			job.Execute(ctx, &url)
		})
	}

	wg.Wait()
}
