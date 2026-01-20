package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jchrisos/pingenemy/internal/httpclient"
	"github.com/jchrisos/pingenemy/internal/job"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer func() {
		fmt.Println("\nJob is done")
		stop()
	}()

	urls, err := httpclient.NewUrl().RetriveUrls()
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
