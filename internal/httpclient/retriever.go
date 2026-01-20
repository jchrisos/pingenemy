package httpclient

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func RetrieveUrlsFromLocalFile() ([]UrlRequest, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, ".pingenemy", "urls.json")

	log.Printf("Loading urls file: %s\n", path)

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error loading urls file: %s", path)
		file, err := os.Create(path)
		if err != nil {
			log.Printf("Error Creating urls file: %s", path)
			return nil, err
		}
		defer file.Close()

		fileBytes, err = json.MarshalIndent(defaultUrls, "", "  ")
		if err != nil {
			return nil, err
		}

		_, err = file.Write(fileBytes)
		if err != nil {
			log.Printf("Error writing urls file: %s", path)
			return nil, err
		}
		log.Printf("File created: %s", path)
	}

	var urlRequests []UrlRequest
	err = json.Unmarshal(fileBytes, &urlRequests)
	if err != nil {
		return nil, err
	}

	return urlRequests, nil
}
