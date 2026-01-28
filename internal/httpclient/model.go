package httpclient

type UrlRequest struct {
	Name               string `json:"name"`
	URL                string `json:"url"`
	HttpMethod         string `json:"httpMethod"`
	ExpectedStatusCode int    `json:"expectedStatusCode"`
}

type UrlResult struct {
	Success      bool
	StatusCode   string
	ResponseTime int64
}

var (
	defaultUrls = []UrlRequest{
		{
			Name:               "google",
			URL:                "https://google.com",
			HttpMethod:         "GET",
			ExpectedStatusCode: 200,
		},
	}
)
