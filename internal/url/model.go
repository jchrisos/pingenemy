package url

type UrlRequest struct {
	Name               string
	URL                string
	HttpMethod         string
	ExpectedStatusCode int
}
