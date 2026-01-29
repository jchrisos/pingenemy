package job

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	httpclient "github.com/jchrisos/pingenemy/internal/httpclient"
)

const (
	intervalSeconds = 60
	successColor    = "\033[1;30;102m %s \033[0m"
	errorColor      = "\033[1;97;101m %s \033[0m"
	urlTextMaxLen   = 50
)

func Execute(ctx context.Context, urls []httpclient.UrlRequest, intervalFromArgs int) {
	var interval = intervalSeconds
	if intervalFromArgs > 0 {
		interval = intervalFromArgs
	}
	var wg sync.WaitGroup

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	exec := func() {
		for _, url := range urls {
			u := url
			wg.Go(func() {
				err := Fetch(ctx, &u)
				if err != nil {
					return
				}
			})
		}
	}

	exec()

	for {
		select {
		case <-ticker.C:
			exec()
		case <-ctx.Done():
			wg.Wait()
			return
		}
	}
}

func Fetch(ctx context.Context, urlReq *httpclient.UrlRequest) error {
	result, err := httpclient.Call(ctx, urlReq)
	fmt.Println(FormatMessage(*urlReq, *result))
	if err != nil {
		if !errors.Is(err, context.DeadlineExceeded) {
			log.Printf("Error calling %s error=%v", urlReq.Name, err)
		}
		return err
	}
	return nil
}

func FormatMessage(urlReq httpclient.UrlRequest, result httpclient.UrlResult) string {
	duration := time.Duration(result.ResponseTime) * time.Millisecond
	durationFmt := fmt.Sprintf("%.3fs", duration.Seconds())

	urlFmt := urlReq.URL
	if len(urlReq.URL) > urlTextMaxLen {
		urlFmt = urlReq.URL[:urlTextMaxLen] + "..."
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now().In(loc).Format(time.DateTime)

	message := fmt.Sprintf("%s | %-19s | sc: %-3s | rt: %-6s | %-53s", now, urlReq.Name, result.StatusCode, durationFmt, urlFmt)

	if result.Success {
		return fmt.Sprintf(successColor, message)
	}

	return fmt.Sprintf(errorColor, message)
}
