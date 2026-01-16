package http

import (
	"encoding/json"
	"os"
)

type UrlRetriever struct {
}

func NewUrl() *UrlRetriever {
	return &UrlRetriever{}
}

func (u *UrlRetriever) RetriveUrls() ([]UrlRequest, error) {
	fileBytes, err := os.ReadFile("urls.json")
	if err != nil {
		return nil, err
	}

	var urlRequests []UrlRequest
	err = json.Unmarshal(fileBytes, &urlRequests)
	if err != nil {
		return nil, err
	}

	return urlRequests, nil
}
