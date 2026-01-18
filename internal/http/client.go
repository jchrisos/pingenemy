package http

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

type HttpClient struct {
	timeout time.Duration
}

func NewHttpClient() *HttpClient {
	return &HttpClient{200 * time.Millisecond}
}

func NewHttpClientTest(millis int64) *HttpClient {
	return &HttpClient{time.Duration(millis) * time.Millisecond}
}

func (c *HttpClient) Call(ctx context.Context, urlReq *UrlRequest) UrlResult {
	httpCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(httpCtx, strings.ToUpper(urlReq.HttpMethod), urlReq.URL, nil)

	start := time.Now()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return undefinedResult
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	return UrlResult{
		Success:      resp.StatusCode == urlReq.ExpectedStatusCode,
		StatusCode:   strconv.Itoa(resp.StatusCode),
		ResponseTime: elapsed.Milliseconds(),
	}
}
