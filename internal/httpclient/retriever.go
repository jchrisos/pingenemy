package httpclient

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type UrlRetriever struct {
}

func NewUrl() *UrlRetriever {
	return &UrlRetriever{}
}

func (u *UrlRetriever) RetriveUrls() ([]UrlRequest, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(home, ".pingenemy", "urls.json")

	fileBytes, err := os.ReadFile(path)
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
