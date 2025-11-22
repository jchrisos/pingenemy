package main

import (
	"fmt"

	"github.com/jchrisos/pingenemy/internal/http"
)

func main() {
	url := "https://google.com"
	method := "GET"

	exec := &http.HttpExecutor{}

	success, err := exec.Execute(method, url)
	if err != nil {
		panic("Failed to calling url")
	}

	if success {
		fmt.Println("OK")
	}
}
