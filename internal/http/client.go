package http

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"
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

func (c *HttpClient) Call(ctx context.Context, urlReq *UrlRequest) (bool, int64, error) {
	httpCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(httpCtx, strings.ToUpper(urlReq.HttpMethod), urlReq.URL, nil)

	start := time.Now()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error calling %s url: %s", urlReq.Name, urlReq.URL)
		return false, 0, err
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	return resp.StatusCode == urlReq.ExpectedStatusCode, elapsed.Milliseconds(), nil
}
