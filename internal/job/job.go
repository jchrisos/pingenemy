package job

import (
	"context"
	"fmt"
	"time"

	"github.com/jchrisos/pingenemy/internal/http"
)

const (
	printSuccess = "\033[1;30;102m %s \033[0m"
	printError   = "\033[1;97;101m %s \033[0m"
)

type Job struct {
}

func (j *Job) Execute(ctx context.Context, urlReq *http.UrlRequest) {
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
				result := client.Call(ctx, urlReq)

				duration := time.Duration(result.ResponseTime) * time.Millisecond
				durationFmt := fmt.Sprintf("%.3fs", duration.Seconds())

				urlFmt := urlReq.URL
				if len(urlReq.URL) > 50 {
					urlFmt = urlReq.URL[:50] + "..."
				}

				message := fmt.Sprintf("%-20s sc: %-3s rt: %-6s url: %-53s", urlReq.Name, result.StatusCode, durationFmt, urlFmt)

				if result.Success {
					fmt.Printf(printSuccess, message)
				} else {
					fmt.Printf(printError, message)
				}
				fmt.Println("")
			}()
		}
	}
}
