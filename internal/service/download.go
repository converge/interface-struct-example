package service

import (
	"fmt"
	"net/http"
)

type HTTPClientProvider interface {
	Do(req *http.Request) (*http.Response, error)
}

type Download struct {
	HttpClient  HTTPClientProvider
}

func NewDownloader(httpClient HTTPClientProvider) Download {
	return Download{
		HttpClient:  httpClient,
	}
}

func (download Download) GetIem() error {
	r, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)

	resp, _ := download.HttpClient.Do(r)

	fmt.Println(resp.StatusCode)
	return nil
}