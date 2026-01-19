package http

import (
	"context"
	"log"
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
	return &HttpClient{}
}

func (c *HttpClient) Call(ctx context.Context, urlReq *UrlRequest) UrlResult {
	httpCtx, cancel := context.WithTimeout(ctx, time.Duration(urlReq.TimeoutMillis))
	defer cancel()

	req, _ := http.NewRequestWithContext(httpCtx, strings.ToUpper(urlReq.HttpMethod), urlReq.URL, nil)

	start := time.Now()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
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
