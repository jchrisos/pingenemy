package http

type UrlRequest struct {
	Name               string
	URL                string
	HttpMethod         string
	ExpectedStatusCode int
	IntervalSeconds    int
}
