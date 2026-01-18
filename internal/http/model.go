package http

type UrlRequest struct {
	Name               string `json:"name"`
	URL                string `json:"url"`
	HttpMethod         string `json:"httpMethod"`
	ExpectedStatusCode int    `json:"expectedStatusCode"`
	IntervalSeconds    int    `json:"intervalSeconds"`
}

type UrlResult struct {
	Success      bool
	StatusCode   string
	ResponseTime int64
}
