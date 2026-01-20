package httpclient

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	undefinedResult = UrlResult{
		Success:      false,
		StatusCode:   " - ",
		ResponseTime: 0,
	}
)

func Call(ctx context.Context, urlReq *UrlRequest) (UrlResult, error) {
	httpCtx, cancel := context.WithTimeout(ctx, time.Duration(urlReq.TimeoutMillis)*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(httpCtx, strings.ToUpper(urlReq.HttpMethod), urlReq.URL, nil)

	start := time.Now()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return undefinedResult, err
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	return UrlResult{
		Success:      resp.StatusCode == urlReq.ExpectedStatusCode,
		StatusCode:   strconv.Itoa(resp.StatusCode),
		ResponseTime: elapsed.Milliseconds(),
	}, nil
}
