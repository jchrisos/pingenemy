package job

import (
	"context"
	"fmt"
	"log"
	"time"

	httpclient "github.com/jchrisos/pingenemy/internal/httpclient"
)

const (
	intervalSeconds = 60
	successColor    = "\033[1;30;102m %s \033[0m"
	errorColor      = "\033[1;97;101m %s \033[0m"
	urlTextMaxLen   = 50
)

func Execute(ctx context.Context, urlReq *httpclient.UrlRequest, intervalSecondsFromArgs int) {
	var interval = intervalSeconds
	if intervalSecondsFromArgs > 0 {
		interval = intervalSecondsFromArgs
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	call := func() {
		result, err := httpclient.Call(ctx, urlReq)
		if err != nil {
			log.Printf("Error calling %s error=%v", urlReq.Name, err)
		}

		fmt.Println(FormatMessage(*urlReq, *result))
	}

	call()

	for {
		select {
		case <-ticker.C:
			call()
		case <-ctx.Done():
			return
		}
	}
}

func FormatMessage(urlReq httpclient.UrlRequest, result httpclient.UrlResult) string {
	duration := time.Duration(result.ResponseTime) * time.Millisecond
	durationFmt := fmt.Sprintf("%.3fs", duration.Seconds())

	urlFmt := urlReq.URL
	if len(urlReq.URL) > urlTextMaxLen {
		urlFmt = urlReq.URL[:urlTextMaxLen] + "..."
	}

	message := fmt.Sprintf("%-19s | sc: %-3s | rt: %-6s | %-53s", urlReq.Name, result.StatusCode, durationFmt, urlFmt)

	if result.Success {
		return fmt.Sprintf(successColor, message)
	}

	return fmt.Sprintf(errorColor, message)
}
