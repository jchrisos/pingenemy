package job

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jchrisos/pingenemy/internal/http"
)

const (
	successColor = "\033[1;30;102m %s \033[0m"
	errorColor   = "\033[1;97;101m %s \033[0m"
	maxLength    = 50
)

func Execute(ctx context.Context, urlReq *http.UrlRequest) {
	interval := time.Duration(urlReq.IntervalSeconds) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	client := http.NewHttpClient()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Job is done")
			return
		case <-ticker.C:
			func() {
				result, err := client.Call(ctx, urlReq)
				if err != nil {
					log.Printf("Error calling %s error=%v", urlReq.Name, err)
				}

				fmt.Println(FormatMessage(*urlReq, result))
			}()
		}
	}
}

func FormatMessage(urlReq http.UrlRequest, result http.UrlResult) string {
	duration := time.Duration(result.ResponseTime) * time.Millisecond
	durationFmt := fmt.Sprintf("%.3fs", duration.Seconds())

	urlFmt := urlReq.URL
	if len(urlReq.URL) > maxLength {
		urlFmt = urlReq.URL[:maxLength] + "..."
	}

	message := fmt.Sprintf("%-15s | sc: %-3s | rt: %-6s | %-53s", urlReq.Name, result.StatusCode, durationFmt, urlFmt)

	if result.Success {
		return fmt.Sprintf(successColor, message)
	}

	return fmt.Sprintf(errorColor, message)
}
