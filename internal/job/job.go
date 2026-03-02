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
	intervalSeconds = 60 * 5 //5 minutes
	successColor    = "\033[1;30;102m sc: %-3s \033[0m"
	errorColor      = "\033[1;97;101m sc: %-3s \033[0m"
	location        = "America/Sao_Paulo"
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
		start := time.Now()

		for i := range urls {
			wg.Go(func() {
				err := Fetch(ctx, &urls[i])
				if err != nil {
					return
				}
			})
		}
		wg.Wait()

		total := time.Since(start)
		fmt.Printf("Total execution time: %v\n", formatDuration(float64(total.Seconds())))
	}

	exec()

	for {
		select {
		case <-ticker.C:
			exec()
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("Gracefully Stopping...")
			return
		}
	}
}

func Fetch(ctx context.Context, urlReq *httpclient.UrlRequest) error {
	result, err := httpclient.Call(ctx, urlReq)
	fmt.Println(formatMessage(*urlReq, *result))
	if err != nil {
		if !errors.Is(err, context.DeadlineExceeded) {
			log.Printf("Error calling %s error=%v", urlReq.Name, err)
		}
		return err
	}
	return nil
}

func formatMessage(urlReq httpclient.UrlRequest, result httpclient.UrlResult) string {
	duration := time.Duration(result.ResponseTime) * time.Millisecond

	loc, _ := time.LoadLocation(location)
	now := time.Now().In(loc).Format(time.DateTime)

	statusCode := fmt.Sprintf(successColor, result.StatusCode)
	if !result.Success {
		statusCode = fmt.Sprintf(errorColor, result.StatusCode)
	}

	return fmt.Sprintf("%s | %-19s | %s | rt: %-6s | %s %s",
		now,
		urlReq.Name,
		statusCode,
		formatDuration(duration.Seconds()),
		urlReq.HttpMethod,
		urlReq.URL)
}

func formatDuration(duration float64) string {
	return fmt.Sprintf("%.3fs", duration)
}
