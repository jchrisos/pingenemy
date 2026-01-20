package httpclient

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func RetrieveUrlsFromLocalFile() ([]UrlRequest, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, ".pingenemy", "urls.json")

	fmt.Printf("Loading url files: %s\n", path)

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		file, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		fileBytes, err = json.MarshalIndent(defaultUrls, "", "  ")
		if err != nil {
			return nil, err
		}

		_, err = file.Write(fileBytes)
		if err != nil {
			return nil, err
		}
	}

	var urlRequests []UrlRequest
	err = json.Unmarshal(fileBytes, &urlRequests)
	if err != nil {
		return nil, err
	}

	return urlRequests, nil
}
