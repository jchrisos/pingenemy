package http

import (
	"log"
	"net/http"
	"strings"
)

const (
	SUCCESS_MIN = 200
	SUCCESS_MAX = 299
)

type HttpExecutor struct {
}

func (s *HttpExecutor) Execute(method string, url string) (bool, error) {
	req, _ := http.NewRequest(strings.ToUpper(method), url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error calling url: %s", url)
		return false, err
	}

	isSuccess := resp.StatusCode >= SUCCESS_MIN && resp.StatusCode <= SUCCESS_MAX

	return isSuccess, nil
}
