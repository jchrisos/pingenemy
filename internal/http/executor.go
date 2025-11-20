package http

import (
	"fmt"
	"log"
	"net/http"
)

func Execute(method string, url string) {
	req, _ := http.NewRequest(method, url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error calling url: %s", url)
	}

	if resp.StatusCode == 200 {
		fmt.Println("OK")
	}
}
