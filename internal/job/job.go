package job

import (
	"context"
	"fmt"
	"time"

	"github.com/jchrisos/pingenemy/internal/http"
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

				var status = "NOK"
				if result.Success {
					status = "OK"
				}

				fmt.Printf("%s - %s response time: %d\n", status, urlReq.Name, result.ResponseTime)
			}()
		}
	}
}
