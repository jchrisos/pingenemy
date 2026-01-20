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
		fmt.Println("\nJob canceled")
		stop()
	}()

	urls, err := httpclient.RetrieveUrlsFromLocalFile()
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
