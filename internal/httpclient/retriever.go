package httpclient

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func RetriveUrlsFromLocalFile() ([]UrlRequest, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
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
