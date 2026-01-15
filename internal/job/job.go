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
				success, respTime, err := client.Call(ctx, urlReq)
				if err != nil {
					fmt.Printf("Failed to calling url %s %s\n", urlReq.Name, err.Error())
				}

				if success {
					fmt.Printf("OK - %s\t\t response time: %d\n", urlReq.Name, respTime)
				}
			}()
		}
	}
}
