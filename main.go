package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jchrisos/pingenemy/internal/httpclient"
	"github.com/jchrisos/pingenemy/internal/job"
)

func main() {
	intervalSeconds := flag.Int("i", 0, "interval in seconds")
	flag.Parse()

	urls, err := httpclient.RetrieveUrlsFromLocalFile()
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer func() {
		stop()
	}()

	job.Execute(ctx, urls, *intervalSeconds)
}
