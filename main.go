package main

import (
	"github.com/jchrisos/pingenemy/internal/http"
)

func main() {
	url := "https://google.com"
	method := "GET"

	http.Execute(method, url)
}
