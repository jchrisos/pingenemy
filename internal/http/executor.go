package http

import (
	"fmt"
	"log"
	"net/http"
)

func Execute(url string, method string) {
	req, _ := http.NewRequest(url, method, nil)

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error calling url: %s", url)
	}

	if resp.StatusCode == 200 {
		fmt.Println("OK")
	}
}
