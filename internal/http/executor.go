package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/jchrisos/pingenemy/internal/url"
)

type HttpExecutor struct {
}

func (s *HttpExecutor) Execute(urlReq *url.UrlRequest) (bool, error) {
	req, _ := http.NewRequest(strings.ToUpper(urlReq.HttpMethod), urlReq.URL, nil)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error calling %s url: %s", urlReq.Name, urlReq.URL)
		return false, err
	}

	return resp.StatusCode == urlReq.ExpectedStatusCode, nil
}
